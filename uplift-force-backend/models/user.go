package models

import (
	"gorm.io/gorm"
	"time"
)

type UserRole string
type UserStatus int8
type UserVerified int8

const (
	RoleAdmin   UserRole = "admin"
	RolePlayer  UserRole = "player"
	RoleBooster UserRole = "booster"
)

const (
	StatusDisabled UserStatus = 0 // 禁用
	StatusNormal   UserStatus = 1 // 正常
	StatusPending  UserStatus = 2 // 待审核
)

const (
	NotVerified UserVerified = 0 // 未认证
	Verified    UserVerified = 1 // 已认证
)

type User struct {
	ID            uint64  `json:"id" gorm:"primaryKey;autoIncrement;comment:用户ID"`
	WalletAddress string  `json:"wallet_address" gorm:"type:varchar(42);uniqueIndex:uk_wallet_address;not null;comment:钱包地址(以0x开头)"`
	Username      string  `json:"username" gorm:"type:varchar(50);uniqueIndex:uk_username;not null;comment:用户名"`
	Email         *string `json:"email" gorm:"type:varchar(100);uniqueIndex:uk_email;comment:邮箱"`
	Phone         *string `json:"phone" gorm:"type:varchar(20);comment:手机号"`
	Avatar        *string `json:"avatar" gorm:"type:varchar(500);comment:头像URL"`
	Nickname      *string `json:"nickname" gorm:"type:varchar(100);comment:昵称"`

	// 角色身份相关
	Role       UserRole     `json:"role" gorm:"type:enum('admin','player','booster');not null;default:'player';index:idx_role;comment:用户角色"`
	Status     UserStatus   `json:"status" gorm:"type:tinyint;not null;default:1;index:idx_status;comment:账户状态: 0-禁用 1-正常 2-待审核"`
	IsVerified UserVerified `json:"is_verified" gorm:"type:tinyint;default:0;comment:是否实名认证: 0-未认证 1-已认证"`

	// 基础设置
	Language string `json:"language" gorm:"type:varchar(10);default:'en';comment:首选语言: en, zh, ko, ja等"`
	Timezone string `json:"timezone" gorm:"type:varchar(50);default:'UTC';comment:时区"`

	// 系统字段
	LastLoginAt *time.Time     `json:"last_login_at" gorm:"comment:最后登录时间"`
	CreatedAt   time.Time      `json:"created_at" gorm:"index:idx_created_at;comment:创建时间"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"comment:更新时间"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"comment:软删除时间"`
}

// 设置表名
func (User) TableName() string {
	return "uf_users"
}

// 创建用户请求结构体
type CreateUserRequest struct {
	WalletAddress string    `json:"wallet_address" binding:"required,len=42"`
	Username      string    `json:"username" binding:"required,min=3,max=50"`
	Email         *string   `json:"email" binding:"omitempty,email,max=100"`
	Phone         *string   `json:"phone" binding:"omitempty,max=20"`
	Avatar        *string   `json:"avatar" binding:"omitempty,url,max=500"`
	Nickname      *string   `json:"nickname" binding:"omitempty,max=100"`
	Role          *UserRole `json:"role" binding:"omitempty,oneof=admin player booster"`
	Language      *string   `json:"language" binding:"omitempty,max=10"`
	Timezone      *string   `json:"timezone" binding:"omitempty,max=50"`
}

// 更新用户请求结构体
type UpdateUserRequest struct {
	Username *string     `json:"username" binding:"omitempty,min=3,max=50"`
	Email    *string     `json:"email" binding:"omitempty,email,max=100"`
	Phone    *string     `json:"phone" binding:"omitempty,max=20"`
	Avatar   *string     `json:"avatar" binding:"omitempty,url,max=500"`
	Nickname *string     `json:"nickname" binding:"omitempty,max=100"`
	Role     *UserRole   `json:"role" binding:"omitempty,oneof=admin player booster"`
	Status   *UserStatus `json:"status" binding:"omitempty,oneof=0 1 2"`
	Language *string     `json:"language" binding:"omitempty,max=10"`
	Timezone *string     `json:"timezone" binding:"omitempty,max=50"`
}

// 用户查询参数
type UserQueryParams struct {
	Page       int           `form:"page,default=1" binding:"min=1"`
	PageSize   int           `form:"page_size,default=10" binding:"min=1,max=100"`
	Role       *UserRole     `form:"role" binding:"omitempty,oneof=admin player booster"`
	Status     *UserStatus   `form:"status" binding:"omitempty,oneof=0 1 2"`
	IsVerified *UserVerified `form:"is_verified" binding:"omitempty,oneof=0 1"`
	Search     string        `form:"search"` // 搜索用户名、邮箱、钱包地址
}
