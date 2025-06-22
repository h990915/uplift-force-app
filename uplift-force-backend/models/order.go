package models

import "time"

type Order struct {
	ID                   uint64     `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderNo              string     `json:"order_no" gorm:"uniqueIndex;size:32"`
	PlayerID             uint64     `json:"player_id" gorm:"not null"`
	BoosterID            *uint64    `json:"booster_id,omitempty"`
	ChainOrderID         *uint64    `json:"chain_order_id,omitempty"`
	GameType             string     `json:"game_type" gorm:"size:20;not null"`
	ServerRegion         string     `json:"server_region" gorm:"size:20;not null"`
	GameAccount          string     `json:"game_account" gorm:"size:100;not null"`
	GameMode             string     `json:"game_mode" gorm:"type:enum('RANKED_SOLO_5x5','RANKED_FLEX_SR');not null"`
	ServiceType          string     `json:"service_type" gorm:"type:enum('Boosting','PLAY WITH');not null"`
	CurrentRank          *string    `json:"current_rank,omitempty" gorm:"size:50"`
	TargetRank           *string    `json:"target_rank,omitempty" gorm:"size:50"`
	PUUID                *string    `json:"PUUID,omitempty" gorm:"size:100"`
	Requirements         *string    `json:"requirements,omitempty" gorm:"type:text"`
	TotalAmount          string     `json:"total_amount" gorm:"type:decimal(18,8);not null"`
	PlayerDeposit        string     `json:"player_deposit" gorm:"type:decimal(18,8);not null"`
	RemainingAmount      string     `json:"remaining_amount" gorm:"type:decimal(18,8);not null"`
	BoosterDeposit       *string    `json:"booster_deposit,omitempty" gorm:"type:decimal(18,8)"`
	Status               string     `json:"status" gorm:"type:enum('posted','accepted','confirmed','in_progress','completed','cancelled','failed');default:'posted'"`
	ContractAddress      *string    `json:"contract_address,omitempty" gorm:"size:42"`
	Deadline             time.Time  `json:"deadline" gorm:"not null"`
	DepositTxHash        *string    `json:"deposit_tx_hash,omitempty" gorm:"size:66"`
	BoosterDepositTxHash *string    `json:"booster_deposit_tx_hash,omitempty" gorm:"size:66"`
	PaymentTxHash        *string    `json:"payment_tx_hash,omitempty" gorm:"size:66"`
	SettlementTxHash     *string    `json:"settlement_tx_hash,omitempty" gorm:"size:66"`
	PostedAt             time.Time  `json:"posted_at" gorm:"default:CURRENT_TIMESTAMP"`
	AcceptedAt           *time.Time `json:"accepted_at,omitempty"`
	ConfirmedAt          *time.Time `json:"confirmed_at,omitempty"`
	CompletedAt          *time.Time `json:"completed_at,omitempty"`
	FailedAt             *time.Time `json:"failed_at,omitempty"`
	CancelledAt          *time.Time `json:"cancelled_at,omitempty"`
	CreatedAt            time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt            time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

type OrderLog struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID   uint64    `json:"order_id" gorm:"not null"`
	UserID    uint64    `json:"user_id" gorm:"not null"`
	Action    string    `json:"action" gorm:"size:50;not null"`
	OldStatus *string   `json:"old_status,omitempty" gorm:"size:20"`
	NewStatus *string   `json:"new_status,omitempty" gorm:"size:20"`
	Amount    *string   `json:"amount,omitempty" gorm:"type:decimal(18,8)"`
	TxHash    *string   `json:"tx_hash,omitempty" gorm:"size:66"`
	Note      *string   `json:"note,omitempty" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type CreateOrderRequest struct {
	GameType     string `json:"game_type" binding:"required"`
	ServerRegion string `json:"server_region" binding:"required"`
	GameAccount  string `json:"game_account" binding:"required"`
	GameMode     string `json:"game_mode" binding:"required"`
	ServiceType  string `json:"service_type" binding:"required"`
	CurrentRank  string `json:"current_rank"`
	TargetRank   string `json:"target_rank"`
	PUUID        string `json:"PUUID" binding:"required"`
	Requirements string `json:"requirements"`
	TotalAmount  string `json:"total_amount" binding:"required"`
	Deadline     string `json:"deadline" binding:"required"`
	TxHash       string `json:"tx_hash" binding:"required"`
}

type CreateOrderChainRequest struct {
	TotalAmount   int64  `json:"total_amount" binding:"required,min=1"`
	DeadlineUnix  int64  `json:"deadline_unix" binding:"required"`
	GameType      string `json:"game_type" binding:"required"`
	GameMode      string `json:"game_mode" binding:"required"`
	Requirements  string `json:"requirements"`
	ServerRegion  string `json:"server_region"`
	PlayerAddress string `json:"player_address"`
}
