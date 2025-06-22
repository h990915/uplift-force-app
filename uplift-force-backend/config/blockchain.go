package config

import (
	"os"
	"uplift-force-backend/models"
)

func IsDevelopmentMode() bool {
	env := os.Getenv("APP_ENV")
	return env == "development" || env == "dev" || env == ""
}

func IsProductionMode() bool {
	return os.Getenv("APP_ENV") == "production"
}

func GetLastDeployedContractAddress() string {
	// 从数据库或缓存获取上次使用的合约地址
	var config models.SystemConfig
	if err := DB.Where("config_key = ?", "last_contract_address").First(&config).Error; err != nil {
		return ""
	}
	return config.ConfigValue
}

func UpdateLastContractAddress(address string) error {
	config := models.SystemConfig{
		ConfigKey:   "last_contract_address",
		ConfigValue: address,
	}

	return DB.Where("config_key = ?", "last_contract_address").
		Assign(config).
		FirstOrCreate(&config).Error
}
