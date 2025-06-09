package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uplift-force-backend/config"
	"uplift-force-backend/models"
	"uplift-force-backend/utils"
)

// 创建用户
func CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	// 验证钱包地址格式
	if !strings.HasPrefix(req.WalletAddress, "0x") {
		utils.ErrorResponse(c, http.StatusBadRequest, "钱包地址格式错误", "钱包地址必须以0x开头")
		return
	}

	// 创建用户对象
	user := models.User{
		WalletAddress: req.WalletAddress,
		Username:      req.Username,
		Email:         req.Email,
		Phone:         *req.Phone,
		Avatar:        req.Avatar,
		Nickname:      req.Nickname,
		Role:          models.RolePlayer, // 默认角色
		Status:        models.StatusNormal,
		IsVerified:    models.NotVerified,
		Language:      "en",
		Timezone:      "UTC",
	}

	// 设置可选字段
	if req.Role != nil {
		user.Role = *req.Role
	}
	if req.Language != nil {
		user.Language = *req.Language
	}
	if req.Timezone != nil {
		user.Timezone = *req.Timezone
	}

	// 保存到数据库
	if err := config.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "wallet_address") {
				utils.ErrorResponse(c, http.StatusConflict, "钱包地址已存在", "")
				return
			}
			if strings.Contains(err.Error(), "username") {
				utils.ErrorResponse(c, http.StatusConflict, "用户名已存在", "")
				return
			}
			if strings.Contains(err.Error(), "email") {
				utils.ErrorResponse(c, http.StatusConflict, "邮箱已存在", "")
				return
			}
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "创建用户失败", err.Error())
		return
	}

	utils.SuccessResponse(c, "创建用户成功", user)
}

// 获取用户列表
func GetUsers(c *gin.Context) {
	var params models.UserQueryParams

	if err := c.ShouldBindQuery(&params); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	var users []models.User
	var total int64

	// 构建查询
	query := config.DB.Model(&models.User{})

	// 添加筛选条件
	if params.Role != nil {
		query = query.Where("role = ?", *params.Role)
	}
	if params.Status != nil {
		query = query.Where("status = ?", *params.Status)
	}
	if params.IsVerified != nil {
		query = query.Where("is_verified = ?", *params.IsVerified)
	}
	if params.Search != "" {
		searchPattern := "%" + params.Search + "%"
		query = query.Where("username LIKE ? OR email LIKE ? OR wallet_address LIKE ?",
			searchPattern, searchPattern, searchPattern)
	}

	// 获取总数
	query.Count(&total)

	// 分页查询
	offset := (params.Page - 1) * params.PageSize
	if err := query.Offset(offset).Limit(params.PageSize).Order("created_at DESC").Find(&users).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "查询用户列表失败", err.Error())
		return
	}

	// 返回分页结果
	result := gin.H{
		"users": users,
		"pagination": gin.H{
			"page":        params.Page,
			"page_size":   params.PageSize,
			"total":       total,
			"total_pages": (total + int64(params.PageSize) - 1) / int64(params.PageSize),
		},
	}

	utils.SuccessResponse(c, "获取用户列表成功", result)
}

// 根据ID获取用户
func GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "用户ID格式错误", "")
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "用户不存在", "")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "查询用户失败", err.Error())
		return
	}

	utils.SuccessResponse(c, "获取用户成功", user)
}

// 根据钱包地址获取用户
func GetUserByWallet(c *gin.Context) {
	walletAddress := c.Param("address")

	var user models.User
	if err := config.DB.Where("wallet_address = ?", walletAddress).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "用户不存在", "")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "查询用户失败", err.Error())
		return
	}

	utils.SuccessResponse(c, "获取用户成功", user)
}

// 更新用户
func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "用户ID格式错误", "")
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	// 查找用户
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "用户不存在", "")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "查询用户失败", err.Error())
		return
	}

	// 更新字段
	updates := make(map[string]interface{})
	if req.Username != nil {
		updates["username"] = *req.Username
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.Phone != nil {
		updates["phone"] = *req.Phone
	}
	if req.Avatar != nil {
		updates["avatar"] = *req.Avatar
	}
	if req.Nickname != nil {
		updates["nickname"] = *req.Nickname
	}
	if req.Role != nil {
		updates["role"] = *req.Role
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.Language != nil {
		updates["language"] = *req.Language
	}
	if req.Timezone != nil {
		updates["timezone"] = *req.Timezone
	}

	// 执行更新
	if err := config.DB.Model(&user).Updates(updates).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "更新用户失败", err.Error())
		return
	}

	// 重新查询更新后的用户信息
	config.DB.First(&user, id)
	utils.SuccessResponse(c, "更新用户成功", user)
}

// 更新最后登录时间
func UpdateLastLogin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "用户ID格式错误", "")
		return
	}

	now := time.Now()
	if err := config.DB.Model(&models.User{}).Where("id = ?", id).Update("last_login_at", now).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "更新登录时间失败", err.Error())
		return
	}

	utils.SuccessResponse(c, "更新登录时间成功", gin.H{"last_login_at": now})
}

// 软删除用户
func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "用户ID格式错误", "")
		return
	}

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "删除用户失败", err.Error())
		return
	}

	utils.SuccessResponse(c, "删除用户成功", nil)
}
