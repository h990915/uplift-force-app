// controllers/order_controller.go
package controllers

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"uplift-force-backend/config"
	"uplift-force-backend/models"
	"uplift-force-backend/services"
	"uplift-force-backend/utils"
)

// CreateOrder 创建游戏代练订单
// @Summary		创建代练订单
// @Description	用户提交游戏代练订单信息，包含区块链交易哈希进行验证
// @Tags			Orders
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		models.CreateOrderRequest	true	"创建订单信息"
// @Success		200		{object}	utils.Response{data=models.Order}
// @Failure		400		{object}	utils.Response
// @Failure		401		{object}	utils.Response
// @Failure		403		{object}	utils.Response
// @Failure		422		{object}	utils.Response
// @Failure		500		{object}	utils.Response
// @Router			/orders [post]
func CreateOrder(c *gin.Context) {
	var req models.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	// 验证游戏模式
	if !isValidGameMode(req.GameMode) {
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的游戏模式", "")
		return
	}

	// 验证服务类型
	if !isValidServiceType(req.ServiceType) {
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的服务类型", "")
		return
	}

	// 验证交易哈希格式
	if !strings.HasPrefix(req.TxHash, "0x") {
		utils.ErrorResponse(c, http.StatusBadRequest, "交易哈希格式错误", "交易哈希必须以0x开头")
		return
	}

	// 获取当前用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权", "")
		return
	}

	playerID := userID.(uint64)

	// 解析截止时间
	deadlineInt, err := strconv.ParseInt(req.Deadline, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "时间格式错误", "请传入有效的Unix时间戳字符串")
		return
	}
	deadline := time.Unix(deadlineInt, 0) // 转为 time.Time

	// 计算金额 (15% 保证金, 85% 剩余)
	totalAmount := req.TotalAmount
	playerDeposit := calculatePercentage(totalAmount, 15)
	remainingAmount := calculatePercentage(totalAmount, 85)

	// 验证交易确实存在且成功
	blockchainService, err := services.NewBlockchainService()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "区块链服务初始化失败", err.Error())
		return
	}

	// 验证交易
	isValid, chainOrderID, err := blockchainService.VerifyOrderTransaction(req.TxHash, playerID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "交易验证失败", err.Error())
		return
	}

	if !isValid {
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的交易", "")
		return
	}

	// 生成订单号
	orderNo := generateOrderNo()

	// 去除字符串末尾的数字
	re := regexp.MustCompile(`\d+$`)
	cleanServerRegion := re.ReplaceAllString(req.ServerRegion, "")
	// 创建订单
	order := models.Order{
		OrderNo:         orderNo,
		PlayerID:        playerID,
		GameType:        req.GameType,
		ServerRegion:    cleanServerRegion,
		GameAccount:     req.GameAccount,
		GameMode:        req.GameMode,
		ServiceType:     req.ServiceType,
		CurrentRank:     stringPtr(req.CurrentRank),
		TargetRank:      stringPtr(req.TargetRank),
		Requirements:    stringPtr(req.Requirements),
		TotalAmount:     totalAmount,
		PlayerDeposit:   playerDeposit,
		RemainingAmount: remainingAmount,
		Status:          "posted",
		Deadline:        deadline,
		DepositTxHash:   &req.TxHash,
		PostedAt:        time.Now(),
		ChainOrderID:    &chainOrderID,
		PUUID:           &req.PUUID,
	}

	// 开启事务
	tx := config.DB.Begin()

	// 保存订单
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "订单创建失败", err.Error())
		return
	}

	// 记录日志
	log := models.OrderLog{
		OrderID:   order.ID,
		UserID:    playerID,
		Action:    "create",
		NewStatus: stringPtr(order.Status),
		Amount:    &playerDeposit,
		TxHash:    &req.TxHash,
		Note:      stringPtr(fmt.Sprintf("玩家创建%s订单并支付保证金", getGameModeDisplayName(req.GameMode))),
	}

	if err := tx.Create(&log).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "日志记录失败", err.Error())
		return
	}

	tx.Commit()

	utils.SuccessResponse(c, "订单创建成功", order)
}

// 2. 代练接单
func AcceptOrder(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权", "")
		return
	}

	boosterID := userID.(uint64)

	var req struct {
		OrderID uint64 `json:"order_id" binding:"required"`
		TxHash  string `json:"tx_hash" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	if !strings.HasPrefix(req.TxHash, "0x") {
		utils.ErrorResponse(c, http.StatusBadRequest, "交易哈希格式错误", "交易哈希必须以0x开头")
		return
	}

	// 开启事务
	tx := config.DB.Begin()

	var order models.Order
	if err := tx.Where("id = ? AND status = ?", req.OrderID, "posted").First(&order).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusBadRequest, "订单不存在或状态不正确", "")
		return
	}

	// 不能接自己的单
	if order.PlayerID == boosterID {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusBadRequest, "不能接取自己发布的订单", "")
		return
	}

	// 计算代练保证金 (15%)
	boosterDeposit := calculatePercentage(order.TotalAmount, 15)
	now := time.Now()

	// 更新订单
	oldStatus := order.Status
	order.BoosterID = &boosterID
	order.BoosterDeposit = &boosterDeposit
	order.Status = "accepted"
	order.BoosterDepositTxHash = &req.TxHash
	order.AcceptedAt = &now

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "订单更新失败", err.Error())
		return
	}

	// 记录日志
	log := models.OrderLog{
		OrderID:   order.ID,
		UserID:    boosterID,
		Action:    "accept",
		OldStatus: &oldStatus,
		NewStatus: stringPtr(order.Status),
		Amount:    &boosterDeposit,
		TxHash:    &req.TxHash,
		Note:      stringPtr(fmt.Sprintf("代练接取%s订单并支付保证金", getGameModeDisplayName(order.GameMode))),
	}

	if err := tx.Create(&log).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "日志记录失败", err.Error())
		return
	}

	tx.Commit()

	utils.SuccessResponse(c, "接单成功", order)
}

// 3. 玩家确认订单并支付剩余金额
func ConfirmOrder(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权", "")
		return
	}

	playerID := userID.(uint64)

	var req struct {
		OrderID uint64 `json:"order_id" binding:"required"`
		TxHash  string `json:"tx_hash" binding:"required"`
		//ContractAddress string `json:"contract_address" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	if !strings.HasPrefix(req.TxHash, "0x") {
		utils.ErrorResponse(c, http.StatusBadRequest, "地址格式错误", "交易哈希和合约地址必须以0x开头")
		return
	}

	// 开启事务
	tx := config.DB.Begin()

	var order models.Order
	if err := tx.Where("id = ? AND player_id = ? AND status = ?", req.OrderID, playerID, "accepted").First(&order).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusBadRequest, "订单不存在或无权限操作", "")
		return
	}

	// 更新订单状态
	oldStatus := order.Status
	now := time.Now()
	order.Status = "in_progress"
	order.PaymentTxHash = &req.TxHash
	order.ConfirmedAt = &now

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "订单确认失败", err.Error())
		return
	}

	// 记录日志
	log := models.OrderLog{
		OrderID:   order.ID,
		UserID:    playerID,
		Action:    "confirm",
		OldStatus: &oldStatus,
		NewStatus: stringPtr(order.Status),
		Amount:    &order.RemainingAmount,
		TxHash:    &req.TxHash,
		Note:      stringPtr(fmt.Sprintf("玩家确认%s订单并支付剩余金额,订单开始", getGameModeDisplayName(order.GameMode))),
	}

	if err := tx.Create(&log).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "日志记录失败", err.Error())
		return
	}

	tx.Commit()

	response := gin.H{
		"order": order,
	}

	utils.SuccessResponse(c, "订单确认成功，智能合约已创建", response)
}

// 4. 取消订单
func CancelOrder(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权", "")
		return
	}

	currentUserID := userID.(uint64)

	var req struct {
		OrderID uint64 `json:"order_id" binding:"required"`
		Reason  string `json:"reason"`
		TxHash  string `json:"tx_hash"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	// 开启事务
	tx := config.DB.Begin()

	var order models.Order
	if err := tx.Where("id = ?", req.OrderID).First(&order).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusBadRequest, "订单不存在", "")
		return
	}

	// 验证权限和状态
	if order.Status == "confirmed" {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusBadRequest, "订单已确认，无法取消", "")
		return
	}

	var canCancel bool
	var penaltyNote string

	if order.Status == "posted" && order.PlayerID == currentUserID {
		canCancel = true
		penaltyNote = "代练未接单前取消，全额退还玩家保证金"
	} else if order.Status == "accepted" {
		if order.PlayerID == currentUserID {
			canCancel = true
			penaltyNote = "玩家在代练接单后取消，扣除玩家保证金，2/3给代练，1/3给平台"
		} else if order.BoosterID != nil && *order.BoosterID == currentUserID {
			canCancel = true
			penaltyNote = "代练接单后取消，扣除代练保证金，2/3给玩家，1/3给平台"
		}
	}

	if !canCancel {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusForbidden, "无权限取消此订单", "")
		return
	}

	// 更新订单状态
	oldStatus := order.Status
	now := time.Now()
	order.Status = "cancelled"
	order.CancelledAt = &now

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "订单取消失败", err.Error())
		return
	}

	// 记录日志
	log := models.OrderLog{
		OrderID:   order.ID,
		UserID:    currentUserID,
		Action:    "cancel",
		OldStatus: &oldStatus,
		NewStatus: stringPtr(order.Status),
		TxHash:    &req.TxHash,
		Note:      &penaltyNote,
	}

	if err := tx.Create(&log).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "日志记录失败", err.Error())
		return
	}

	tx.Commit()

	response := gin.H{
		"order":        order,
		"penalty_info": penaltyNote,
	}

	utils.SuccessResponse(c, "订单取消成功", response)
}

// GetOrders 获取订单列表
// @Summary 获取订单列表
// @Description 分页查询订单列表，支持多种筛选条件，包含关联的玩家和代练师用户名信息
// @Tags Orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "页码，默认为1" minimum(1) default(1)
// @Param page_size query int false "每页数量，默认为20，最大100" minimum(1) maximum(100) default(20)
// @Param status query string false "订单状态" Enums(posted, accepted, confirmed, in_progress, completed, cancelled, failed)
// @Param game_type query string false "游戏类型" example(LOL)
// @Param game_mode query string false "游戏模式" Enums(ranked_solo, ranked_flex)
// @Param service_type query string false "服务类型" Enums(boost)
// @Param user_filter query string false "用户筛选：my=我的订单，available=可接单订单" Enums(my, available)
// @Success 200 {object} object{code=int,message=string,data=object{orders=[]object{id=int,order_no=string,player_id=int,player_username=string,booster_id=int,booster_username=string,game_type=string,server_region=string,game_account=string,game_mode=string,service_type=string,current_rank=string,target_rank=string,total_amount=string,player_deposit=string,remaining_amount=string,status=string,deadline=string,posted_at=string,created_at=string,updated_at=string},total=int,page=int,page_size=int}} "查询成功"
// @Failure 400 {object} object{code=int,message=string,error=string} "参数错误"
// @Failure 401 {object} object{code=int,message=string} "未授权访问"
// @Failure 500 {object} object{code=int,message=string,error=string} "服务器内部错误"
// @Router /api/v1/orders [get]
func GetOrders(c *gin.Context) {
	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// 筛选参数
	status := c.Query("status")
	gameType := c.Query("game_type")
	gameMode := c.Query("game_mode")
	serviceType := c.Query("service_type")

	userID, exists := c.Get("user_id")
	var currentUserID uint64
	if exists {
		currentUserID = userID.(uint64)
	}

	// 应用筛选条件的函数
	applyFilters := func(q *gorm.DB) *gorm.DB {
		if status != "" {
			q = q.Where("o.status = ?", status)
		}
		if gameType != "" {
			q = q.Where("o.game_type = ?", gameType)
		}
		if gameMode != "" {
			q = q.Where("o.game_mode = ?", gameMode)
		}
		if serviceType != "" {
			q = q.Where("o.service_type = ?", serviceType)
		}

		// 用户相关筛选
		userFilter := c.Query("user_filter")
		if userFilter == "my" && exists {
			q = q.Where("o.player_id = ? OR o.booster_id = ?", currentUserID, currentUserID)
		} else if userFilter == "available" {
			if exists {
				q = q.Where("o.status = 'posted' AND o.player_id != ?", currentUserID)
			} else {
				q = q.Where("o.status = 'posted'")
			}
		}
		return q
	}

	// 计算总数 - 单独的查询
	var total int64
	countQuery := applyFilters(config.DB.Table("orders o"))
	countQuery.Count(&total)

	// 构建数据查询
	baseQuery := config.DB.Table("orders o").
		Select(`o.*, 
			p.username as player_username, 
			b.username as booster_username`).
		Joins("LEFT JOIN uf_users p ON o.player_id = p.id").
		Joins("LEFT JOIN uf_users b ON o.booster_id = b.id")

	// 应用筛选条件到数据查询
	dataQuery := applyFilters(baseQuery)

	// 分页查询 - 使用匿名结构体接收查询结果
	var orders []struct {
		models.Order
		PlayerUsername  *string `json:"player_username"`
		BoosterUsername *string `json:"booster_username"`
	}

	offset := (page - 1) * pageSize
	if err := dataQuery.Offset(offset).Limit(pageSize).Order("o.created_at DESC").Scan(&orders).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "查询失败", err.Error())
		return
	}

	response := gin.H{
		"orders":    orders,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}

	utils.SuccessResponse(c, "查询成功", response)
}

// 6. 获取订单详情
func GetOrderDetail(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "订单ID格式错误", "")
		return
	}

	var order models.Order
	if err := config.DB.Where("id = ?", orderID).First(&order).Error; err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "订单不存在", "")
		return
	}

	// 获取操作日志
	var logs []models.OrderLog
	config.DB.Where("order_id = ?", orderID).Order("created_at DESC").Find(&logs)

	response := gin.H{
		"order": order,
		"logs":  logs,
	}

	utils.SuccessResponse(c, "查询成功", response)
}

// CompleteOrder 手动完成订单（代练）
// @Summary		手动完成订单
// @Description	代练可以手动将订单标记为完成（成功或失败），支持备注和交易哈希
// @Tags			Orders
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			request	body		object{order_id=uint64,completion_status=string,note=string,tx_hash=string}	true	"完成订单信息"
// @Success		200		{object}	utils.Response{data=models.Order}
// @Failure		400		{object}	utils.Response
// @Failure		401		{object}	utils.Response
// @Failure		403		{object}	utils.Response
// @Failure		500		{object}	utils.Response
// @Router			/orders/complete [post]
func CompleteOrder(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权", "")
		return
	}

	boosterID := userID.(uint64)

	var req struct {
		OrderID          uint64 `json:"order_id" binding:"required"`          // 订单ID
		CompletionStatus string `json:"completion_status" binding:"required"` // 完成状态：completed 或 failed
		Note             string `json:"note"`                                 // 完成备注
		TxHash           string `json:"tx_hash"`                              // 结算交易哈希
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	// 验证完成状态参数
	if req.CompletionStatus != "completed" && req.CompletionStatus != "failed" {
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的完成状态", "completion_status 必须是 'completed' 或 'failed'")
		return
	}

	if req.TxHash != "" && !strings.HasPrefix(req.TxHash, "0x") {
		utils.ErrorResponse(c, http.StatusBadRequest, "交易哈希格式错误", "交易哈希必须以0x开头")
		return
	}

	// 开启事务
	tx := config.DB.Begin()

	var order models.Order
	if err := tx.Debug().Where("chain_order_id = ? AND booster_id = ? AND status IN ?", req.OrderID, boosterID, []string{"confirmed", "in_progress"}).First(&order).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusBadRequest, "订单不存在或无权限操作", "只有确认状态的订单可以手动完成")
		return
	}

	// 更新订单状态
	oldStatus := order.Status
	now := time.Now()
	order.Status = req.CompletionStatus

	if req.CompletionStatus == "completed" {
		order.CompletedAt = &now
		if req.TxHash != "" {
			order.SettlementTxHash = &req.TxHash
		}
	} else if req.CompletionStatus == "failed" {
		order.FailedAt = &now
		if req.TxHash != "" {
			order.SettlementTxHash = &req.TxHash
		}
	}

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "订单状态更新失败", err.Error())
		return
	}

	// 记录日志
	var noteText string
	var actionType string

	if req.CompletionStatus == "completed" {
		noteText = "代练手动完成订单（成功完成）"
		actionType = "complete"
	} else {
		noteText = "代练手动完成订单（失败完成）"
		actionType = "manual_fail"
	}

	if req.Note != "" {
		noteText += ": " + req.Note
	}

	log := models.OrderLog{
		OrderID:   order.ID,
		UserID:    boosterID,
		Action:    actionType,
		OldStatus: &oldStatus,
		NewStatus: stringPtr(order.Status),
		TxHash:    &req.TxHash,
		Note:      stringPtr(noteText),
	}

	if err := tx.Create(&log).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "日志记录失败", err.Error())
		return
	}

	tx.Commit()

	// 根据完成状态返回不同的消息
	var message string
	if req.CompletionStatus == "completed" {
		message = "订单成功完成"
	} else {
		message = "订单已标记为失败完成"
	}

	response := gin.H{
		"order":             order,
		"completion_status": req.CompletionStatus,
		"completion_type":   getCompletionTypeDisplay(req.CompletionStatus),
	}

	utils.SuccessResponse(c, message, response)
}

// 获取完成类型的显示名称
func getCompletionTypeDisplay(status string) string {
	switch status {
	case "completed":
		return "成功完成"
	case "failed":
		return "失败完成"
	default:
		return "未知状态"
	}
}

// 8. 创建争议（玩家/代练）
func CreateDispute(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "订单ID格式错误", "")
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权", "")
		return
	}

	currentUserID := userID.(uint64)

	var req struct {
		Reason string `json:"reason" binding:"required"` // 争议原因
		Detail string `json:"detail"`                    // 详细描述
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	// 开启事务
	tx := config.DB.Begin()

	var order models.Order
	if err := tx.Where("id = ?", orderID).First(&order).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusBadRequest, "订单不存在", "")
		return
	}

	// 验证权限（只有订单相关的玩家或代练可以创建争议）
	if order.PlayerID != currentUserID && (order.BoosterID == nil || *order.BoosterID != currentUserID) {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusForbidden, "无权限创建争议", "只有订单参与者可以创建争议")
		return
	}

	// 验证订单状态（只有进行中、已完成的订单可以争议）
	if order.Status != "confirmed" && order.Status != "in_progress" && order.Status != "completed" {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusBadRequest, "当前订单状态不允许创建争议", "只有确认后的订单可以创建争议")
		return
	}

	// 检查是否已经存在争议
	if order.Status == "disputed" {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusBadRequest, "订单已存在争议", "请等待管理员处理")
		return
	}

	// 更新订单状态
	oldStatus := order.Status
	order.Status = "disputed"

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "争议创建失败", err.Error())
		return
	}

	// 确定争议发起人身份
	userRole := "玩家"
	if order.BoosterID != nil && *order.BoosterID == currentUserID {
		userRole = "代练"
	}

	// 记录日志
	noteText := fmt.Sprintf("%s创建争议 - 原因: %s", userRole, req.Reason)
	if req.Detail != "" {
		noteText += ", 详情: " + req.Detail
	}

	log := models.OrderLog{
		OrderID:   order.ID,
		UserID:    currentUserID,
		Action:    "dispute",
		OldStatus: &oldStatus,
		NewStatus: stringPtr(order.Status),
		Note:      stringPtr(noteText),
	}

	if err := tx.Create(&log).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(c, http.StatusInternalServerError, "日志记录失败", err.Error())
		return
	}

	tx.Commit()

	response := gin.H{
		"order":      order,
		"dispute_id": order.ID, // 可以后续扩展单独的争议表
	}

	utils.SuccessResponse(c, "争议创建成功，等待管理员处理", response)
}

// 9. 获取我的订单（我发布的+我接取的）
func GetMyOrders(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权", "")
		return
	}

	currentUserID := userID.(uint64)

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// 筛选参数
	status := c.Query("status")      // 'posted','accepted','confirmed','in_progress','completed','cancelled','failed'
	role := c.Query("role")          // "player" 或 "booster" 或 空（全部）
	gameType := c.Query("game_type") // 目前仅支持 'League of Legends'
	gameMode := c.Query("game_mode") // RANKED_SOLO_5x5 or RANKED_FLEX_SR

	// 应用筛选条件的函数
	applyFilters := func(q *gorm.DB) *gorm.DB {
		// 根据角色筛选
		if role == "player" {
			q = q.Where("o.player_id = ?", currentUserID)
		} else if role == "booster" {
			q = q.Where("o.booster_id = ?", currentUserID)
		} else {
			// 默认显示所有相关订单（我发布的 + 我接取的）
			q = q.Where("o.player_id = ? OR o.booster_id = ?", currentUserID, currentUserID)
		}

		// 其他筛选条件
		if status != "" {
			q = q.Where("o.status = ?", status)
		}
		if gameType != "" {
			q = q.Where("o.game_type = ?", gameType)
		}
		if gameMode != "" {
			q = q.Where("o.game_mode = ?", gameMode)
		}

		return q
	}

	// 计算总数 - 单独的查询
	var total int64
	countQuery := applyFilters(config.DB.Table("orders o"))
	countQuery.Count(&total)

	// 构建数据查询
	baseQuery := config.DB.Table("orders o").
		Select(`o.*, 
			p.username as player_username, 
			p.email as player_email,
			b.username as booster_username, 
			b.email as booster_email`).
		Joins("LEFT JOIN uf_users p ON o.player_id = p.id").
		Joins("LEFT JOIN uf_users b ON o.booster_id = b.id")

	// 应用筛选条件到数据查询
	dataQuery := applyFilters(baseQuery)

	// 分页查询 - 使用匿名结构体接收查询结果
	var orders []struct {
		models.Order
		PlayerUsername  *string `json:"player_username"`
		PlayerEmail     *string `json:"player_email"`
		BoosterUsername *string `json:"booster_username"`
		BoosterEmail    *string `json:"booster_email"`
	}

	offset := (page - 1) * pageSize
	if err := dataQuery.Offset(offset).Limit(pageSize).Order("o.created_at DESC").Scan(&orders).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "查询失败", err.Error())
		return
	}

	// 为每个订单添加用户在其中的角色信息
	type OrderWithRoleAndUsernames struct {
		models.Order
		PlayerUsername  *string `json:"player_username"`
		BoosterUsername *string `json:"booster_username"`
		PlayerEmail     *string `json:"player_email"`
		BoosterEmail    *string `json:"booster_email"`
		MyRole          string  `json:"my_role"` // "player" 或 "booster"
	}

	ordersWithRoleAndUsernames := make([]OrderWithRoleAndUsernames, len(orders))
	for i, order := range orders {
		role := "player"
		if order.BoosterID != nil && *order.BoosterID == currentUserID {
			role = "booster"
		}
		ordersWithRoleAndUsernames[i] = OrderWithRoleAndUsernames{
			Order:           order.Order,
			PlayerUsername:  order.PlayerUsername,
			BoosterUsername: order.BoosterUsername,
			PlayerEmail:     order.PlayerEmail,
			BoosterEmail:    order.BoosterEmail,
			MyRole:          role,
		}
	}

	response := gin.H{
		"orders":    ordersWithRoleAndUsernames,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}

	utils.SuccessResponse(c, "查询成功", response)
}

// 10. 获取可接单的订单
func GetAvailableOrders(c *gin.Context) {
	userID, exists := c.Get("user_id")
	var currentUserID uint64
	if exists {
		currentUserID = userID.(uint64)
	}

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	// 筛选参数
	gameType := c.Query("game_type")
	gameMode := c.Query("game_mode")
	serviceType := c.Query("service_type")
	serverRegion := c.Query("server_region")
	minAmount := c.Query("min_amount") // 最小金额
	maxAmount := c.Query("max_amount") // 最大金额
	sortBy := c.Query("sort_by")       // "amount_desc", "amount_asc", "time_desc", "time_asc"

	// 只查询可接单的订单
	query := config.DB.Model(&models.Order{}).Where("status = 'posted'")

	// 排除自己发布的订单（如果已登录）
	if exists {
		query = query.Where("player_id != ?", currentUserID)
	}

	// 筛选条件
	if gameType != "" {
		query = query.Where("game_type = ?", gameType)
	}
	if gameMode != "" {
		query = query.Where("game_mode = ?", gameMode)
	}
	if serviceType != "" {
		query = query.Where("service_type = ?", serviceType)
	}
	if serverRegion != "" {
		query = query.Where("server_region = ?", serverRegion)
	}
	if minAmount != "" {
		query = query.Where("total_amount >= ?", minAmount)
	}
	if maxAmount != "" {
		query = query.Where("total_amount <= ?", maxAmount)
	}

	// 排序
	orderBy := "created_at DESC" // 默认按创建时间倒序
	switch sortBy {
	case "amount_desc":
		orderBy = "total_amount DESC"
	case "amount_asc":
		orderBy = "total_amount ASC"
	case "time_desc":
		orderBy = "created_at DESC"
	case "time_asc":
		orderBy = "created_at ASC"
	}

	// 计算总数
	var total int64
	query.Count(&total)

	// 分页查询
	var orders []models.Order
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order(orderBy).Find(&orders).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "查询失败", err.Error())
		return
	}

	// 添加一些额外信息
	type AvailableOrder struct {
		models.Order
		TimeRemaining string `json:"time_remaining"` // 剩余时间
		UrgencyLevel  string `json:"urgency_level"`  // 紧急程度
	}

	availableOrders := make([]AvailableOrder, len(orders))
	now := time.Now()

	for i, order := range orders {
		// 计算剩余时间
		timeRemaining := order.Deadline.Sub(now)
		var timeRemainingStr string
		var urgencyLevel string

		if timeRemaining < 0 {
			timeRemainingStr = "已过期"
			urgencyLevel = "expired"
		} else if timeRemaining < 24*time.Hour {
			timeRemainingStr = fmt.Sprintf("%.0f小时", timeRemaining.Hours())
			urgencyLevel = "urgent"
		} else if timeRemaining < 7*24*time.Hour {
			timeRemainingStr = fmt.Sprintf("%.0f天", timeRemaining.Hours()/24)
			urgencyLevel = "normal"
		} else {
			timeRemainingStr = fmt.Sprintf("%.0f天", timeRemaining.Hours()/24)
			urgencyLevel = "relaxed"
		}

		availableOrders[i] = AvailableOrder{
			Order:         order,
			TimeRemaining: timeRemainingStr,
			UrgencyLevel:  urgencyLevel,
		}
	}

	response := gin.H{
		"orders":    availableOrders,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
		"filters": gin.H{
			"game_type":     gameType,
			"game_mode":     gameMode,
			"service_type":  serviceType,
			"server_region": serverRegion,
			"min_amount":    minAmount,
			"max_amount":    maxAmount,
			"sort_by":       sortBy,
		},
	}

	utils.SuccessResponse(c, "查询成功", response)
}

// 11. 获取订单操作日志
func GetOrderLogs(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "订单ID格式错误", "")
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未授权", "")
		return
	}

	currentUserID := userID.(uint64)

	// 验证用户是否有权限查看此订单的日志
	var order models.Order
	if err := config.DB.Where("id = ?", orderID).First(&order).Error; err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "订单不存在", "")
		return
	}

	// 只有订单相关的用户可以查看日志
	if order.PlayerID != currentUserID && (order.BoosterID == nil || *order.BoosterID != currentUserID) {
		utils.ErrorResponse(c, http.StatusForbidden, "无权限查看此订单日志", "只有订单参与者可以查看日志")
		return
	}

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "50"))

	// 获取日志总数
	var total int64
	config.DB.Model(&models.OrderLog{}).Where("order_id = ?", orderID).Count(&total)

	// 获取日志
	var logs []models.OrderLog
	offset := (page - 1) * pageSize
	if err := config.DB.Where("order_id = ?", orderID).
		Offset(offset).Limit(pageSize).
		Order("created_at DESC").Find(&logs).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "查询日志失败", err.Error())
		return
	}

	// 增强日志信息
	type EnhancedLog struct {
		models.OrderLog
		ActionDisplay string `json:"action_display"` // 操作的中文描述
		UserRole      string `json:"user_role"`      // 操作用户的角色
	}

	enhancedLogs := make([]EnhancedLog, len(logs))
	for i, log := range logs {
		// 确定用户角色
		userRole := "unknown"
		if log.UserID == order.PlayerID {
			userRole = "player"
		} else if order.BoosterID != nil && log.UserID == *order.BoosterID {
			userRole = "booster"
		} else {
			userRole = "admin"
		}

		// 操作描述映射
		actionDisplay := getActionDisplayName(log.Action)

		enhancedLogs[i] = EnhancedLog{
			OrderLog:      log,
			ActionDisplay: actionDisplay,
			UserRole:      userRole,
		}
	}

	response := gin.H{
		"logs":      enhancedLogs,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
		"order_basic_info": gin.H{
			"id":           order.ID,
			"order_no":     order.OrderNo,
			"status":       order.Status,
			"game_type":    order.GameType,
			"game_mode":    order.GameMode,
			"service_type": order.ServiceType,
		},
	}

	utils.SuccessResponse(c, "查询成功", response)
}

// 获取操作的中文描述
func getActionDisplayName(action string) string {
	actionMap := map[string]string{
		"create":                "创建订单",
		"accept":                "接受订单",
		"confirm":               "确认订单",
		"cancel":                "取消订单",
		"complete":              "完成订单",
		"dispute":               "创建争议",
		"admin_update_status":   "管理员更新状态",
		"admin_resolve_dispute": "管理员解决争议",
	}

	if display, exists := actionMap[action]; exists {
		return display
	}
	return action
}

// 工具函数
func generateOrderNo() string {
	timestamp := time.Now().Unix()
	randomNum, _ := rand.Int(rand.Reader, big.NewInt(9999))
	return fmt.Sprintf("BO%d%04d", timestamp, randomNum.Int64())
}

func calculatePercentage(amount string, percentage int) string {
	// 这里简化处理，实际项目中需要使用精确的大数计算
	// 推荐使用 github.com/shopspring/decimal 库
	return fmt.Sprintf("%.8f", 0.15) // 示例返回，需要根据实际金额计算
}

func stringPtr(s string) *string {
	return &s
}

func isValidGameMode(mode string) bool {
	validModes := []string{"RANKED_SOLO_5x5", "RANKED_FLEX_SR"}
	for _, validMode := range validModes {
		if mode == validMode {
			return true
		}
	}
	return false
}

func isValidServiceType(serviceType string) bool {
	return serviceType == "Boosting" || serviceType == "PLAY WITH"
}

func getGameModeDisplayName(mode string) string {
	switch mode {
	case "ranked_solo":
		return "排位单双"
	case "ranked_flex":
		return "灵活排位"
	case "normal_draft":
		return "普通征召"
	case "aram":
		return "极地大乱斗"
	default:
		return "未知模式"
	}
}
