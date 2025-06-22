package services

import (
	"context"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"uplift-force-backend/config"
	"uplift-force-backend/models"
)

type EventListener struct {
	blockchainService  *BlockchainService
	lastProcessedBlock uint64
	// 预计算的事件签名哈希
	eventSignatures map[string]common.Hash
}

func NewEventListener(bs *BlockchainService) *EventListener {
	// 根据你的合约事件定义调整这些签名
	eventSignatures := map[string]common.Hash{
		"OrderCreated":   crypto.Keccak256Hash([]byte("OrderCreated(uint256,address,uint256,uint256,string,string,string)")),
		"OrderAccepted":  crypto.Keccak256Hash([]byte("OrderAccepted(uint256,address,uint256)")),
		"OrderCompleted": crypto.Keccak256Hash([]byte("OrderCompleted(uint256,address,address,uint256)")),
		"OrderCancelled": crypto.Keccak256Hash([]byte("OrderCancelled(uint256,address,uint256,address)")),
	}

	// 打印事件签名哈希（用于调试）
	log.Println("事件签名哈希:")
	for name, hash := range eventSignatures {
		log.Printf("  %s: %s", name, hash.Hex())
	}

	return &EventListener{
		blockchainService:  bs,
		lastProcessedBlock: getLastProcessedBlock(),
		eventSignatures:    eventSignatures,
	}
}

// 启动事件监听
func (el *EventListener) Start() {
	log.Println("启动区块链事件监听...")

	for {
		if err := el.processNewBlocks(); err != nil {
			log.Printf("处理区块事件失败: %v", err)
		}
		time.Sleep(5 * time.Second) // 每5秒检查一次
	}
}

func (el *EventListener) processNewBlocks() error {
	// 获取最新区块号
	latestBlock, err := el.blockchainService.client.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	if latestBlock <= el.lastProcessedBlock {
		return nil // 没有新区块
	}

	// 处理新区块中的事件
	fromBlock := el.lastProcessedBlock + 1
	toBlock := latestBlock

	// 创建过滤器查询
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{el.blockchainService.contractAddr},
	}

	// 获取日志
	logs, err := el.blockchainService.client.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}

	// 处理每个日志
	for _, vLog := range logs {
		if err := el.processLog(vLog); err != nil {
			log.Printf("处理日志失败: %v", err)
		}
	}

	el.lastProcessedBlock = latestBlock
	return nil
}

func (el *EventListener) processLog(vLog types.Log) error {
	if len(vLog.Topics) == 0 {
		return nil
	}

	topicHash := vLog.Topics[0]

	// 匹配事件类型
	switch topicHash {
	case el.eventSignatures["OrderCreated"]:
		return el.handleOrderCreated(vLog)
	//case el.eventSignatures["OrderAccepted"]:
	//	return el.handleOrderAccepted(vLog)
	//case el.eventSignatures["OrderCompleted"]:
	//	return el.handleOrderCompleted(vLog)
	//case el.eventSignatures["OrderCancelled"]:
	//	return el.handleOrderCancelled(vLog)
	default:
		log.Printf("未知事件: %s", topicHash.Hex())
	}

	return nil
}

func (el *EventListener) handleOrderCreated(vLog types.Log) error {
	log.Printf("处理OrderCreated事件，交易哈希: %s", vLog.TxHash.Hex())

	// 解析事件数据
	event, err := el.blockchainService.contract.ParseOrderCreated(vLog)
	if err != nil {
		return err
	}

	// 更新数据库中的订单状态
	return el.updateOrderStatusByTxHash(vLog.TxHash.Hex(), "posted", event.OrderId.Uint64())
}

func (el *EventListener) updateOrderStatusByTxHash(txHash string, status string, chainOrderID uint64) error {
	tx := config.DB.Begin()

	// 查找对应的订单
	var order models.Order
	if err := tx.Where("deposit_tx_hash = ?", txHash).First(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新订单状态和链上ID
	order.Status = status
	order.ChainOrderID = &chainOrderID
	now := time.Now()
	order.ConfirmedAt = &now

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	note := "区块链确认订单创建成功"
	// 记录日志
	log := models.OrderLog{
		OrderID:   order.ID,
		UserID:    order.PlayerID,
		Action:    "blockchain_confirmed",
		NewStatus: &status,
		TxHash:    &txHash,
		Note:      &note,
	}

	if err := tx.Create(&log).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	println("订单 %d 状态更新为: %s", order.ID, status)
	return nil
}

func getLastProcessedBlock() uint64 {
	currentContract := os.Getenv("MAIN_CONTRACT_ADDRESS")
	lastContract := config.GetLastDeployedContractAddress()

	// 开发环境：如果合约地址改变了，自动重置
	if config.IsDevelopmentMode() {
		if currentContract != lastContract && lastContract != "" {
			log.Printf("检测到合约地址变化:")
			log.Printf("  旧地址: %s", lastContract)
			log.Printf("  新地址: %s", currentContract)
			log.Println("开发环境自动重置区块监听位置")

			// 更新合约地址记录
			config.UpdateLastContractAddress(currentContract)

			// 重置区块号
			resetLastProcessedBlock()
			return getCurrentBlockNumber() // 从当前区块开始
		}
	}

	// 正常获取逻辑
	return getLastProcessedBlockFromDB()
}

func resetLastProcessedBlock() {
	log.Println("重置上次处理的区块号")
	config.DB.Where("config_key = ?", "last_processed_block").Delete(&models.SystemConfig{})
}

func getCurrentBlockNumber() uint64 {
	// 连接到区块链
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		log.Printf("RPC_URL 环境变量未设置")
		return 0
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Printf("连接区块链失败")
		return 0
	}
	defer client.Close()

	blockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Printf("获取当前区块号失败: %v", err)
		return 0
	}

	log.Printf("从当前区块开始监听: %d", blockNum)
	return blockNum
}

func (el *EventListener) getCurrentBlockNumber() (uint64, error) {
	blockNum, err := el.blockchainService.client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return blockNum, nil
}

// 从数据库获取上次处理的区块号
func getLastProcessedBlockFromDB() uint64 {
	var systemConfig models.SystemConfig

	// 查询数据库中的配置记录
	err := config.DB.Where("config_key = ?", "last_processed_block").First(&systemConfig).Error
	if err != nil {
		if err.Error() == "record not found" {
			log.Println("数据库中未找到上次处理的区块号记录，返回默认值 0")
		} else {
			log.Printf("查询上次处理区块号失败: %v，返回默认值 0", err)
		}
		return 0
	}

	// 将字符串转换为uint64
	blockNum, err := strconv.ParseUint(systemConfig.ConfigValue, 10, 64)
	if err != nil {
		log.Printf("解析区块号失败: %v，返回默认值 0", err)
		return 0
	}

	log.Printf("从数据库读取上次处理的区块号: %d", blockNum)
	return blockNum
}
