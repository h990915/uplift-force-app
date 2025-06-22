package services

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"uplift-force-backend/config"
	"uplift-force-backend/contracts"
	"uplift-force-backend/models"
)

type BlockchainService struct {
	client       *ethclient.Client
	contract     *contracts.BoostChainMainContract
	filterer     *contracts.BoostChainMainContractFilterer // 添加这个
	privateKey   *ecdsa.PrivateKey
	contractAddr common.Address
	chainID      *big.Int
}

func NewBlockchainService() (*BlockchainService, error) {
	// 连接到区块链
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		return nil, fmt.Errorf("RPC_URL 环境变量未设置")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("连接区块链失败: %v", err)
	}

	// 从环境变量获取私钥字符串
	privateKeyHex := os.Getenv("PRIVATE_KEY_1")
	if privateKeyHex == "" {
		return nil, fmt.Errorf("环境变量 PRIVATE_KEY_1 未设置")
	}

	// 将十六进制字符串转换为私钥对象
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("加载私钥失败: %v", err)
	}

	// 连接合约
	contractAddr := common.HexToAddress(os.Getenv("MAIN_CONTRACT_ADDRESS"))
	contract, err := contracts.NewBoostChainMainContract(contractAddr, client)
	if err != nil {
		return nil, fmt.Errorf("连接合约失败: %v", err)
	}

	// 创建filterer用于解析事件
	filterer, err := contracts.NewBoostChainMainContractFilterer(contractAddr, client)
	if err != nil {
		return nil, fmt.Errorf("创建事件过滤器失败: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("获取链ID失败: %v", err)
	}

	return &BlockchainService{
		client:       client,
		contract:     contract,
		filterer:     filterer,
		privateKey:   privateKey,
		contractAddr: contractAddr,
		chainID:      chainID,
	}, nil
}

// 创建链上订单
func (bs *BlockchainService) CreateOrderOnChain(orderData models.CreateOrderChainRequest) (*OrderCreationResult, error) {
	// 准备交易参数
	auth, err := bs.prepareTransactor()
	if err != nil {
		return nil, err
	}

	// 计算保证金
	totalAmount := big.NewInt(orderData.TotalAmount)
	deposit := bs.calculateDeposit(totalAmount)
	auth.Value = deposit

	// 准备合约参数
	deadline := big.NewInt(time.Unix(orderData.DeadlineUnix, 0).Unix())

	// 调用合约
	tx, err := bs.contract.CreateOrder(
		auth,
		totalAmount,
		deadline,
		orderData.GameType,
		orderData.GameMode,
		orderData.Requirements,
	)

	if err != nil {
		return nil, fmt.Errorf("创建订单交易失败: %v", err)
	}

	return &OrderCreationResult{
		TxHash:        tx.Hash().Hex(),
		TotalAmount:   totalAmount,
		DepositAmount: deposit,
		Status:        "pending",
	}, nil
}

// 验证交易是否成功
func (bs *BlockchainService) VerifyTransaction(txHash string) (*TransactionResult, error) {
	hash := common.HexToHash(txHash)

	// 获取交易回执
	receipt, err := bs.client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil, fmt.Errorf("获取交易回执失败: %v", err)
	}

	// 检查交易状态
	if receipt.Status == 0 {
		return &TransactionResult{
			Success: false,
			Error:   "交易执行失败",
		}, nil
	}

	logs := make([]types.Log, len(receipt.Logs))
	for i, l := range receipt.Logs {
		logs[i] = *l // 解引用指针
	}

	// 解析事件日志获取订单ID
	orderID, err := bs.parseOrderCreatedEvent(logs)
	if err != nil {
		return nil, fmt.Errorf("解析事件失败: %v", err)
	}

	return &TransactionResult{
		Success:  true,
		OrderID:  orderID,
		BlockNum: receipt.BlockNumber.Uint64(),
		GasUsed:  receipt.GasUsed,
	}, nil
}

// 解析OrderCreated事件获取订单ID
func (bs *BlockchainService) parseOrderCreatedEvent(logs []types.Log) (uint64, error) {
	// 预计算OrderCreated事件的签名哈希
	orderCreatedSig := crypto.Keccak256Hash([]byte("OrderCreated(uint256,address,uint256,uint256,string,string)"))

	// 遍历所有日志查找OrderCreated事件
	for _, vLog := range logs {
		// 检查是否是我们关心的合约地址
		if vLog.Address != bs.contractAddr {
			continue
		}

		// 检查是否是OrderCreated事件
		if len(vLog.Topics) > 0 && vLog.Topics[0] == orderCreatedSig {
			// 使用abigen生成的解析方法
			event, err := bs.filterer.ParseOrderCreated(vLog)
			if err != nil {
				log.Printf("解析OrderCreated事件失败: %v", err)
				continue
			}

			log.Printf("成功解析OrderCreated事件:")
			log.Printf("  订单ID: %s", event.OrderId.String())
			log.Printf("  玩家地址: %s", event.Player.Hex())
			log.Printf("  总金额: %s wei", event.TotalAmount.String())

			return event.OrderId.Uint64(), nil
		}
	}

	return 0, fmt.Errorf("在交易日志中未找到OrderCreated事件")
}

// 解析单个OrderCreated事件（用于事件监听）
func (bs *BlockchainService) ParseSingleOrderCreatedEvent(vLog types.Log) (*contracts.BoostChainMainContractOrderCreated, error) {
	// 检查是否是我们关心的合约地址
	if vLog.Address != bs.contractAddr {
		return nil, fmt.Errorf("事件来源地址不匹配")
	}

	// 使用abigen生成的解析方法
	event, err := bs.filterer.ParseOrderCreated(vLog)
	if err != nil {
		return nil, fmt.Errorf("解析OrderCreated事件失败: %v", err)
	}

	return event, nil
}

// 解析OrderAccepted事件
func (bs *BlockchainService) ParseOrderAcceptedEvent(vLog types.Log) (*contracts.BoostChainMainContractOrderAccepted, error) {
	if vLog.Address != bs.contractAddr {
		return nil, fmt.Errorf("事件来源地址不匹配")
	}

	event, err := bs.filterer.ParseOrderAccepted(vLog)
	if err != nil {
		return nil, fmt.Errorf("解析OrderAccepted事件失败: %v", err)
	}

	return event, nil
}

// 解析OrderCompleted事件
func (bs *BlockchainService) ParseOrderCompletedEvent(vLog types.Log) (*contracts.BoostChainMainContractOrderCompleted, error) {
	if vLog.Address != bs.contractAddr {
		return nil, fmt.Errorf("事件来源地址不匹配")
	}

	event, err := bs.filterer.ParseOrderCompleted(vLog)
	if err != nil {
		return nil, fmt.Errorf("解析OrderCompleted事件失败: %v", err)
	}

	return event, nil
}

// 辅助方法
func (bs *BlockchainService) prepareTransactor() (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(bs.privateKey, bs.chainID)
	if err != nil {
		return nil, err
	}

	// 获取nonce
	publicKey := bs.privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	nonce, err := bs.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	// 获取gas价格
	gasPrice, err := bs.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(500000)

	return auth, nil
}

func (bs *BlockchainService) calculateDeposit(totalAmount *big.Int) *big.Int {
	depositRate := big.NewInt(1500) // 15%
	basisPoints := big.NewInt(10000)
	deposit := new(big.Int).Mul(totalAmount, depositRate)
	return deposit.Div(deposit, basisPoints)
}

// 同步创建订单（等待区块链确认）
func (bs *BlockchainService) CreateOrderOnChainSync(orderData models.CreateOrderChainRequest) (*OrderCreationSyncResult, error) {
	log.Println("开始发送创建订单交易...")

	// 1. 发送交易
	result, err := bs.CreateOrderOnChain(orderData)
	if err != nil {
		return nil, fmt.Errorf("发送交易失败: %v", err)
	}

	log.Printf("交易已发送，哈希: %s，等待确认...", result.TxHash)

	// 2. 等待交易确认
	confirmed, err := bs.WaitForTransactionConfirmation(result.TxHash, 60) // 60秒超时
	if err != nil {
		return nil, fmt.Errorf("等待交易确认失败: %v", err)
	}

	if !confirmed.Success {
		return nil, fmt.Errorf("交易执行失败: %s", confirmed.Error)
	}

	log.Printf("交易确认成功！区块号: %d，Gas使用: %d", confirmed.BlockNumber, confirmed.GasUsed)

	return &OrderCreationSyncResult{
		TxHash:        result.TxHash,
		ChainOrderID:  confirmed.ChainOrderID,
		BlockNumber:   confirmed.BlockNumber,
		GasUsed:       confirmed.GasUsed,
		TotalAmount:   result.TotalAmount,
		DepositAmount: result.DepositAmount,
		Status:        "confirmed",
	}, nil
}

// 等待交易确认
func (bs *BlockchainService) WaitForTransactionConfirmation(txHash string, timeoutSeconds int) (*TransactionConfirmResult, error) {
	hash := common.HexToHash(txHash)
	timeout := time.After(time.Duration(timeoutSeconds) * time.Second)
	ticker := time.NewTicker(2 * time.Second) // 每2秒检查一次
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return nil, fmt.Errorf("交易确认超时 (%d 秒)", timeoutSeconds)

		case <-ticker.C:
			// 检查交易状态
			receipt, err := bs.client.TransactionReceipt(context.Background(), hash)
			if err != nil {
				// 交易还没被打包，继续等待
				log.Printf("等待交易被打包... (%s)", txHash)
				continue
			}

			// 交易已被打包，检查执行状态
			if receipt.Status == 0 {
				return &TransactionConfirmResult{
					Success: false,
					Error:   "交易执行失败",
				}, nil
			}

			// 交易成功，解析事件获取订单ID
			logs := make([]types.Log, len(receipt.Logs))
			for i, l := range receipt.Logs {
				logs[i] = *l // 解引用指针
			}

			// 解析事件日志获取订单ID
			chainOrderID, err := bs.parseOrderCreatedEvent(logs)
			if err != nil {
				return nil, fmt.Errorf("解析交易事件失败: %v", err)
			}

			log.Printf("交易确认成功！区块: %d, Gas: %d, 链上订单ID: %d",
				receipt.BlockNumber.Uint64(), receipt.GasUsed, chainOrderID)

			return &TransactionConfirmResult{
				Success:      true,
				ChainOrderID: chainOrderID,
				BlockNumber:  receipt.BlockNumber.Uint64(),
				GasUsed:      receipt.GasUsed,
				TxHash:       txHash,
			}, nil
		}
	}
}

// 简化版验证订单交易
func (bs *BlockchainService) VerifyOrderTransaction(txHash string, expectedPlayerID uint64) (bool, uint64, error) {
	// 1. 获取交易收据
	hash := common.HexToHash(txHash)
	receipt, err := bs.client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return false, 0, fmt.Errorf("获取交易收据失败: %v", err)
	}

	// 2. 检查交易是否成功
	if receipt.Status != 1 {
		return false, 0, fmt.Errorf("交易执行失败")
	}

	// 3. 验证交易目标地址
	tx, _, err := bs.client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return false, 0, fmt.Errorf("获取交易详情失败: %v", err)
	}

	if tx.To() == nil || tx.To().Hex() != bs.contractAddr.Hex() {
		return false, 0, fmt.Errorf("交易目标地址不正确")
	}

	// 4. 简单解析事件获取订单ID和玩家地址
	chainOrderID, playerAddress, err := bs.parseBasicOrderEvent(receipt)
	if err != nil {
		return false, 0, fmt.Errorf("解析订单事件失败: %v", err)
	}

	// 5. 验证玩家地址（可选，如果有绑定钱包的话）
	if expectedPlayerID > 0 {
		expectedAddr, err := bs.getPlayerAddressByID(expectedPlayerID)
		if err == nil && expectedAddr != "" {
			if !strings.EqualFold(playerAddress, expectedAddr) {
				return false, 0, fmt.Errorf("交易发送者不匹配")
			}
		}
	}

	log.Printf("交易验证成功: 哈希=%s, 链上订单ID=%d, 玩家地址=%s",
		txHash, chainOrderID, playerAddress)

	return true, chainOrderID, nil
}

// 简化版事件解析 - 只提取必要信息
func (bs *BlockchainService) parseBasicOrderEvent(receipt *types.Receipt) (uint64, string, error) {
	// OrderCreated事件的签名: OrderCreated(uint256,address,uint256,uint256,uint256,string,string,string)
	eventSignature := crypto.Keccak256Hash([]byte("OrderCreated(uint256,address,uint256,uint256,string,string)"))

	for _, log := range receipt.Logs {
		if len(log.Topics) >= 3 && log.Topics[0] == eventSignature {
			// Topic[1] = orderId (indexed)
			// Topic[2] = player address (indexed)

			orderID := new(big.Int).SetBytes(log.Topics[1][:]).Uint64()
			playerAddr := common.BytesToAddress(log.Topics[2][:]).Hex()

			return orderID, playerAddr, nil
		}
	}

	return 0, "", fmt.Errorf("未找到OrderCreated事件")
}

// 通过用户ID获取钱包地址（简化版）
func (bs *BlockchainService) getPlayerAddressByID(playerID uint64) (string, error) {
	var walletAddr string
	err := config.DB.Model(&models.User{}).
		Select("wallet_address").
		Where("id = ?", playerID).
		Scan(&walletAddr).Error

	return walletAddr, err
}

// 数据结构
type OrderCreationSyncResult struct {
	TxHash        string   `json:"tx_hash"`
	ChainOrderID  uint64   `json:"chain_order_id"`
	BlockNumber   uint64   `json:"block_number"`
	GasUsed       uint64   `json:"gas_used"`
	TotalAmount   *big.Int `json:"total_amount"`
	DepositAmount *big.Int `json:"deposit_amount"`
	Status        string   `json:"status"`
}

type TransactionConfirmResult struct {
	Success      bool   `json:"success"`
	ChainOrderID uint64 `json:"chain_order_id,omitempty"`
	BlockNumber  uint64 `json:"block_number,omitempty"`
	GasUsed      uint64 `json:"gas_used,omitempty"`
	TxHash       string `json:"tx_hash,omitempty"`
	Error        string `json:"error,omitempty"`
}

// 数据结构
type OrderCreationResult struct {
	TxHash        string   `json:"tx_hash"`
	TotalAmount   *big.Int `json:"total_amount"`
	DepositAmount *big.Int `json:"deposit_amount"`
	Status        string   `json:"status"`
}

type TransactionResult struct {
	Success  bool   `json:"success"`
	OrderID  uint64 `json:"order_id,omitempty"`
	BlockNum uint64 `json:"block_num,omitempty"`
	GasUsed  uint64 `json:"gas_used,omitempty"`
	Error    string `json:"error,omitempty"`
}
