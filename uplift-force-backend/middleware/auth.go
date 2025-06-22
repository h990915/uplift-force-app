package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"uplift-force-backend/config"
	"uplift-force-backend/models"
	"uplift-force-backend/utils"
)

// JWT配置
var (
	JWTSecret       = []byte(os.Getenv("JWT_SECRET")) // JWT TOKEN
	AccessTokenTTL  = time.Hour * 4                   // 访问token有效期：4小时
	RefreshTokenTTL = time.Hour * 24 * 7              // 刷新token有效期：7天
)

// JWT Claims结构
type Claims struct {
	UserID        uint64          `json:"user_id"`
	WalletAddress string          `json:"wallet_address"`
	Username      string          `json:"username"`
	Role          models.UserRole `json:"role"`
	TokenType     string          `json:"token_type"` // "access" 或 "refresh"
	jwt.RegisteredClaims
}

// 生成token对
func GenerateTokenPair(user *models.User) (accessToken, refreshToken string, err error) {
	now := time.Now()

	// 生成访问token
	accessClaims := &Claims{
		UserID:        user.ID,
		WalletAddress: user.WalletAddress,
		Username:      user.Username,
		Role:          user.Role,
		TokenType:     "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(AccessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "uplift-force-backend",
			Subject:   fmt.Sprintf("%d", user.ID),
		},
	}

	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = accessTokenObj.SignedString(JWTSecret)
	if err != nil {
		return "", "", err
	}

	// 生成刷新token
	refreshClaims := &Claims{
		UserID:        user.ID,
		WalletAddress: user.WalletAddress,
		Username:      user.Username,
		Role:          user.Role,
		TokenType:     "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(RefreshTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "uplift-force-backend",
			Subject:   fmt.Sprintf("%d", user.ID),
		},
	}

	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refreshTokenObj.SignedString(JWTSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// 解析token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// 验证用户是否仍然有效
func ValidateUser(userID uint64) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	// 检查用户状态
	if user.Status != models.StatusNormal {
		return nil, fmt.Errorf("user account is disabled or pending")
	}

	return &user, nil
}

// JWT中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从header获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "缺少Authorization头", "")
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Authorization头格式错误", "Bearer token required")
			c.Abort()
			return
		}

		accessToken := parts[1]

		// 解析访问token
		claims, err := ParseToken(accessToken)
		if err == nil && claims.TokenType == "access" {
			// 访问token有效，验证用户
			user, err := ValidateUser(claims.UserID)
			if err != nil {
				utils.ErrorResponse(c, http.StatusUnauthorized, "用户无效", err.Error())
				c.Abort()
				return
			}

			// 设置用户信息到上下文
			c.Set("user_id", user.ID)
			c.Set("user", user)
			c.Set("wallet_address", user.WalletAddress)
			c.Set("role", user.Role)
			c.Next()
			return
		}

		// 访问token无效或过期，尝试使用刷新token
		refreshToken := c.GetHeader("X-Refresh-Token")
		if refreshToken == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "访问token已过期且未提供刷新token", "")
			c.Abort()
			return
		}

		// 解析刷新token
		refreshClaims, err := ParseToken(refreshToken)
		if err != nil || refreshClaims.TokenType != "refresh" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "刷新token无效", err.Error())
			c.Abort()
			return
		}

		// 验证用户
		user, err := ValidateUser(refreshClaims.UserID)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "用户无效", err.Error())
			c.Abort()
			return
		}

		// 生成新的token对
		newAccessToken, newRefreshToken, err := GenerateTokenPair(user)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "生成新token失败", err.Error())
			c.Abort()
			return
		}

		// 在响应头中返回新的token
		c.Header("X-New-Access-Token", newAccessToken)
		c.Header("X-New-Refresh-Token", newRefreshToken)

		// 设置用户信息到上下文
		c.Set("user_id", user.ID)
		c.Set("user", user)
		c.Set("wallet_address", user.WalletAddress)
		c.Set("role", user.Role)
		c.Set("token_refreshed", true) // 标记token已刷新

		c.Next()
	}
}

// 角色权限中间件
func RequireRole(roles ...models.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			utils.ErrorResponse(c, http.StatusUnauthorized, "用户角色信息缺失", "")
			c.Abort()
			return
		}

		role := userRole.(models.UserRole)

		// 检查是否有权限
		hasPermission := false
		for _, requiredRole := range roles {
			if role == requiredRole {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			utils.ErrorResponse(c, http.StatusForbidden, "权限不足", fmt.Sprintf("需要角色: %v", roles))
			c.Abort()
			return
		}

		c.Next()
	}
}

// 管理员权限中间件
func RequireAdmin() gin.HandlerFunc {
	return RequireRole(models.RoleAdmin)
}

// 获取当前用户信息的辅助函数
func GetCurrentUser(c *gin.Context) (*models.User, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}
	return user.(*models.User), true
}

// 获取当前用户ID的辅助函数
func GetCurrentUserID(c *gin.Context) (uint64, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}
	return userID.(uint64), true
}
