package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"uplift-force-backend/config"
	"uplift-force-backend/middleware"
	"uplift-force-backend/models"
	"uplift-force-backend/utils"
)

// 验证以太坊签名
func verifyEthereumSignature(message, signature, walletAddress string) error {
	address := common.HexToAddress(walletAddress)
	messageHash := accounts.TextHash([]byte(message))

	sigBytes, err := hexutil.Decode(signature)
	if err != nil {
		return fmt.Errorf("invalid signature format: %v", err)
	}

	if len(sigBytes) == 65 {
		if sigBytes[64] == 27 || sigBytes[64] == 28 {
			sigBytes[64] -= 27
		}
	}

	pubKey, err := crypto.SigToPub(messageHash, sigBytes)
	if err != nil {
		return fmt.Errorf("failed to recover public key: %v", err)
	}

	recoveredAddress := crypto.PubkeyToAddress(*pubKey)

	if !strings.EqualFold(recoveredAddress.Hex(), address.Hex()) {
		return fmt.Errorf("signature verification failed")
	}

	return nil
}

// CheckWallet 检查钱包是否已注册
// @Summary      检查钱包注册状态
// @Description  检查指定钱包地址是否已在系统中注册
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body models.CheckWalletRequest true "钱包地址信息"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}  utils.Response
// @Router       /auth/checkWallet [post]
func CheckWallet(c *gin.Context) {
	var req models.CheckWalletRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	if !strings.HasPrefix(req.WalletAddress, "0x") {
		utils.ErrorResponse(c, http.StatusBadRequest, "钱包地址格式错误", "钱包地址必须以0x开头")
		return
	}

	// 查找用户
	var user models.User
	err := config.DB.Where("wallet_address = ?", req.WalletAddress).First(&user).Error

	response := models.CheckWalletResponse{
		WalletAddress: req.WalletAddress,
		IsRegistered:  err == nil,
	}

	if err == nil {
		response.User = &user
	}

	utils.SuccessResponse(c, "检查钱包状态成功", response)
}

// Login 用户登录
// @Summary      钱包签名登录
// @Description  使用钱包签名进行用户登录验证
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body models.LoginRequest true "登录信息"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}  utils.Response
// @Failure      401  {object}  utils.Response
// @Failure      403  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /auth/login [post]
func Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	if !strings.HasPrefix(req.WalletAddress, "0x") {
		utils.ErrorResponse(c, http.StatusBadRequest, "钱包地址格式错误", "钱包地址必须以0x开头")
		return
	}

	// 验证签名
	if err := verifyEthereumSignature(req.Message, req.Signature, req.WalletAddress); err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "签名验证失败", err.Error())
		return
	}

	// 查找用户
	var user models.User
	if err := config.DB.Where("wallet_address = ?", req.WalletAddress).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "用户不存在", "请先注册")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "查询用户失败", err.Error())
		return
	}

	// 检查用户状态
	if user.Status != models.StatusNormal {
		utils.ErrorResponse(c, http.StatusForbidden, "账户状态异常", "账户被禁用或待审核")
		return
	}

	// 生成token对
	accessToken, refreshToken, err := middleware.GenerateTokenPair(&user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "生成token失败", err.Error())
		return
	}

	// 更新最后登录时间
	now := time.Now()
	config.DB.Model(&user).Update("last_login_at", now)
	user.LastLoginAt = &now

	// 返回登录成功响应
	response := models.AuthResponse{
		User:         &user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(middleware.AccessTokenTTL.Seconds()),
		IsNewUser:    false,
	}

	utils.SuccessResponse(c, "登录成功", response)
}

// Register 用户注册
// @Summary      钱包地址注册
// @Description  使用钱包地址注册新用户账户
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body models.RegisterRequest true "注册信息"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}  utils.Response
// @Failure      409  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /auth/register [post]
func Register(c *gin.Context) {
	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	if !strings.HasPrefix(req.WalletAddress, "0x") {
		utils.ErrorResponse(c, http.StatusBadRequest, "钱包地址格式错误", "钱包地址必须以0x开头")
		return
	}

	// 验证签名
	//if err := verifyEthereumSignature(req.Message, req.Signature, req.WalletAddress); err != nil {
	//	utils.ErrorResponse(c, http.StatusUnauthorized, "签名验证失败", err.Error())
	//	return
	//}

	// 检查钱包地址是否已注册
	var existingUser models.User
	if err := config.DB.Where("wallet_address = ?", req.WalletAddress).First(&existingUser).Error; err == nil {
		utils.ErrorResponse(c, http.StatusConflict, "钱包地址已注册", "")
		return
	}

	// 检查用户名是否已存在
	//if err := config.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
	//	utils.ErrorResponse(c, http.StatusConflict, "用户名已存在", "")
	//	return
	//}

	// 创建新用户
	user := models.User{
		WalletAddress: req.WalletAddress,
		Username:      req.Username,
		Email:         &req.Email,
		Role:          req.Role,
		Status:        models.StatusNormal,
		IsVerified:    models.NotVerified,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "创建用户失败", err.Error())
		return
	}

	// 生成token对
	accessToken, refreshToken, err := middleware.GenerateTokenPair(&user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "生成token失败", err.Error())
		return
	}

	// 返回注册成功响应
	response := models.AuthResponse{
		User:         &user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(middleware.AccessTokenTTL.Seconds()),
		IsNewUser:    true,
	}

	utils.SuccessResponse(c, "注册成功", response)
}

// RefreshToken 刷新访问令牌
// @Summary      刷新访问令牌
// @Description  使用刷新令牌获取新的访问令牌
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body models.RefreshTokenRequest true "刷新令牌信息"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}  utils.Response
// @Failure      401  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Router       /auth/refresh [post]
func RefreshToken(c *gin.Context) {
	var req models.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	// 解析刷新token
	claims, err := middleware.ParseToken(req.RefreshToken)
	if err != nil || claims.TokenType != "refresh" {
		utils.ErrorResponse(c, http.StatusUnauthorized, "刷新token无效", err.Error())
		return
	}

	// 验证用户
	user, err := middleware.ValidateUser(claims.UserID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户无效", err.Error())
		return
	}

	// 生成新的token对
	newAccessToken, newRefreshToken, err := middleware.GenerateTokenPair(user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "生成新token失败", err.Error())
		return
	}

	// 返回新的token
	response := models.LoginResponse{
		User:         user,
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    int64(middleware.AccessTokenTTL.Seconds()),
	}

	utils.SuccessResponse(c, "刷新token成功", response)
}

// GetProfile 获取用户资料
// @Summary      获取当前用户资料
// @Description  获取已登录用户的个人资料信息
// @Tags         User
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  utils.Response
// @Failure      401  {object}  utils.Response
// @Router       /auth/profile [get]
func GetProfile(c *gin.Context) {
	user, exists := middleware.GetCurrentUser(c)
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "获取用户信息失败", "")
		return
	}

	utils.SuccessResponse(c, "获取用户信息成功", user)
}

// Logout 用户登出
// @Summary      用户登出
// @Description  用户登出系统
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  utils.Response
// @Failure      401  {object}  utils.Response
// @Router       /auth/logout [post]
func Logout(c *gin.Context) {
	// 在生产环境中，你可能想要将token加入黑名单
	utils.SuccessResponse(c, "登出成功", nil)
}

// VerifyWallet 验证钱包签名
// @Summary      验证钱包签名
// @Description  验证钱包签名的有效性，不进行登录操作
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body models.LoginRequest true "钱包验证信息"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}  utils.Response
// @Failure      401  {object}  utils.Response
// @Router       /auth/verify [post]
func VerifyWallet(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "参数验证失败", err.Error())
		return
	}

	// 验证钱包地址格式
	if !strings.HasPrefix(req.WalletAddress, "0x") {
		utils.ErrorResponse(c, http.StatusBadRequest, "钱包地址格式错误", "钱包地址必须以0x开头")
		return
	}

	// 验证签名
	if err := verifyEthereumSignature(req.Message, req.Signature, req.WalletAddress); err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "签名验证失败", err.Error())
		return
	}

	utils.SuccessResponse(c, "钱包验证成功", gin.H{
		"wallet_address": req.WalletAddress,
		"verified":       true,
		"message":        req.Message,
	})
}
