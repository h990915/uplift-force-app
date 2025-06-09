package models

// 检查钱包状态请求
type CheckWalletRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required,len=42"`
}

// 检查钱包状态响应
type CheckWalletResponse struct {
	WalletAddress string `json:"wallet_address"`
	IsRegistered  bool   `json:"is_registered"`
	User          *User  `json:"user,omitempty"`
}

// 注册请求结构
type RegisterRequest struct {
	WalletAddress string   `json:"wallet_address" binding:"required,len=42"`
	Username      string   `json:"username" binding:"required,min=1,max=50"`
	Email         string   `json:"Email" binding:"required,min=1,max=100"`
	Role          UserRole `json:"role" binding:"required,oneof=player booster"`
	Signature     string   `json:"signature" binding:"required"`
	Message       string   `json:"message" binding:"required"`
}

// 登录请求结构
type LoginRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required,len=42"`
	Signature     string `json:"signature" binding:"required"`
	Message       string `json:"message" binding:"required"`
}

// 登录/注册响应结构
type AuthResponse struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	IsNewUser    bool   `json:"is_new_user"` // 标识是否为新注册用户
}

// 登录响应结构
type LoginResponse struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// 刷新token请求
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
