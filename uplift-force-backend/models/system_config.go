package models

import (
	"time"
)

type SystemConfig struct {
	ID          uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	ConfigKey   string    `json:"config_key" gorm:"uniqueIndex;not null"`
	ConfigValue string    `json:"config_value" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

func (SystemConfig) TableName() string {
	return "system_configs"
}
