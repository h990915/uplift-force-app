// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";

contract BoostChainMainContract is ReentrancyGuard, Ownable, Pausable {
    
    // ==================== 状态变量 ====================
    
    // 平台费率 (5% = 500 basis points)
    uint256 public platformFeeRate = 500; // 5%
    uint256 public constant BASIS_POINTS = 10000; // 100%
    
    // 保证金比例 (15% = 1500 basis points)
    uint256 public depositRate = 1500; // 15%
    
    // 惩罚分配比例 (2/3 给对方, 1/3 给平台)
    uint256 public penaltyToVictimRate = 6667; // 66.67%
    uint256 public penaltyToPlatformRate = 3333; // 33.33%
    
    // 平台收益地址
    address public platformTreasury;
    
    // 订单计数器
    uint256 public orderCounter;
    
    // ==================== 数据结构 ====================
    
    enum OrderStatus {
        Posted,      // 0: 已发布 (玩家已质押15%)
        Accepted,    // 1: 已接单 (代练已质押15%)
        Confirmed,   // 2: 已确认 (玩家已支付85%, 子合约已创建)
        InProgress,  // 3: 进行中
        Completed,   // 4: 已完成
        Failed,      // 5: 失败 (未达成目标)
        Cancelled    // 6: 已取消
    }
    
    struct Order {
        uint256 orderId;
        address player;           // 玩家地址
        address booster;         // 代练地址
        uint256 totalAmount;     // 订单总金额
        uint256 playerDeposit;   // 玩家保证金 (15%)
        uint256 boosterDeposit;  // 代练保证金 (15%)
        uint256 remainingAmount; // 剩余金额 (85%)
        OrderStatus status;
        uint256 deadline;        // 截止时间
        address childContract;   // 子合约地址
        uint256 createdAt;
        uint256 acceptedAt;
        uint256 confirmedAt;
        string gameType;         // 游戏类型
        string gameMode;         // 游戏模式
        string requirements;     // 需求描述
    }
    
    // ==================== 存储映射 ====================
    
    mapping(uint256 => Order) public orders;
    mapping(address => uint256[]) public playerOrders;   // 玩家的订单列表
    mapping(address => uint256[]) public boosterOrders;  // 代练的订单列表
    mapping(address => bool) public authorizedCallers;   // 授权的调用者映射
    
    // ==================== 事件定义 ====================
    
    event OrderCreated(
        uint256 indexed orderId,
        address indexed player,
        uint256 totalAmount,
        uint256 playerDeposit,
        string gameType,
        string gameMode
    );
    
    event OrderAccepted(
        uint256 indexed orderId,
        address indexed booster,
        uint256 boosterDeposit
    );
    
    event OrderConfirmed(
        uint256 indexed orderId,
        address indexed childContract,
        uint256 totalAmountLocked
    );
    
    event OrderCancelled(
        uint256 indexed orderId,
        address indexed cancelledBy,
        uint256 penaltyAmount,
        address penaltyReceiver
    );
    
    event OrderCompleted(
        uint256 indexed orderId,
        uint256 platformFee,
        uint256 boosterReward,
        bytes32 currentTxHash
    );
    
    event OrderFailed(
        uint256 indexed orderId,
        uint256 playerRefund,
        uint256 penaltyToPlayer,
        uint256 penaltyToPlatform,
        bytes32 currentTxHash
    );

    // 授权事件
    event CallerAuthorized(address indexed caller, bool authorized);
    
    // ==================== 修饰器 ====================
    
    modifier onlyValidOrder(uint256 _orderId) {
        require(_orderId > 0 && _orderId <= orderCounter, "Invalid order ID");
        _;
    }
    
    modifier onlyPlayer(uint256 _orderId) {
        require(orders[_orderId].player == msg.sender, "Only player can call this");
        _;
    }
    
    modifier onlyBooster(uint256 _orderId) {
        require(orders[_orderId].booster == msg.sender, "Only booster can call this");
        _;
    }
    
    modifier onlyPlayerOrBooster(uint256 _orderId) {
        require(
            orders[_orderId].player == msg.sender || 
            orders[_orderId].booster == msg.sender,
            "Only player or booster can call this"
        );
        _;
    }

    modifier onlyAuthorizedOrOwner() {
        require(
            authorizedCallers[msg.sender] || msg.sender == owner(),
            "Unauthorized: Not authorized caller or owner"
        );
        _;
    }
    
    // ==================== 构造函数 ====================
    
    constructor(address _platformTreasury) Ownable(msg.sender) {
        require(_platformTreasury != address(0), "Invalid treasury address");
        platformTreasury = _platformTreasury;
    }

    // ==================== 授权调用函数 ====================

    /**
     * @dev 授权调用者
     * @param _caller 要授权的地址
     * @param _authorized 是否授权
     */
    function setAuthorizedCaller(address _caller, bool _authorized) external onlyOwner {
        require(_caller != address(0), "Invalid caller address");
        authorizedCallers[_caller] = _authorized;
        emit CallerAuthorized(_caller, _authorized);
    }
    
    /**
     * @dev 批量授权
     * @param _callers 要授权的地址数组
     * @param _authorized 是否授权
     */
    function setAuthorizedCallers(address[] memory _callers, bool _authorized) external onlyOwner {
        for (uint256 i = 0; i < _callers.length; i++) {
            require(_callers[i] != address(0), "Invalid caller address");
            authorizedCallers[_callers[i]] = _authorized;
            emit CallerAuthorized(_callers[i], _authorized);
        }
    }
    
    // ==================== 核心功能函数 ====================
    
    /**
     * @dev 创建订单 - 玩家质押15%保证金
     * @param _totalAmount 订单总金额
     * @param _deadline 截止时间
     * @param _gameType 游戏类型
     * @param _gameMode 游戏模式  
     * @param _requirements 需求描述
     */
    function createOrder(
        uint256 _totalAmount,
        uint256 _deadline,
        string memory _gameType,
        string memory _gameMode,
        string memory _requirements
    ) external payable nonReentrant whenNotPaused {
        require(_totalAmount > 0, "Total amount must be greater than 0");
        require(_deadline > block.timestamp, "Deadline must be in the future");
        require(bytes(_gameType).length > 0, "Game type cannot be empty");
        require(bytes(_gameMode).length > 0, "Game mode cannot be empty");
        
        // 计算保证金 (15%)
        uint256 playerDeposit = (_totalAmount * depositRate) / BASIS_POINTS;
        require(msg.value == playerDeposit, "Incorrect deposit amount");
        
        // 创建订单
        orderCounter++;
        uint256 orderId = orderCounter;
        
        orders[orderId] = Order({
            orderId: orderId,
            player: msg.sender,
            booster: address(0),
            totalAmount: _totalAmount,
            playerDeposit: msg.value,
            boosterDeposit: 0,
            remainingAmount: _totalAmount - playerDeposit,
            status: OrderStatus.Posted,
            deadline: _deadline,
            childContract: address(0),
            createdAt: block.timestamp,
            acceptedAt: 0,
            confirmedAt: 0,
            gameType: _gameType,
            gameMode: _gameMode,
            requirements: _requirements
        });
        
        // 添加到玩家订单列表
        playerOrders[msg.sender].push(orderId);
        
        emit OrderCreated(
            orderId,
            msg.sender,
            _totalAmount,
            msg.value,
            _gameType,
            _gameMode
        );
    }
    
    /**
     * @dev 接受订单 - 代练质押15%保证金
     * @param _orderId 订单ID
     */
    function acceptOrder(uint256 _orderId) 
        external 
        payable 
        nonReentrant 
        whenNotPaused 
        onlyValidOrder(_orderId) 
    {
        Order storage order = orders[_orderId];
        require(order.status == OrderStatus.Posted, "Order not available for acceptance");
        require(order.player != msg.sender, "Cannot accept your own order");
        require(block.timestamp < order.deadline, "Order has expired");
        
        // 计算代练保证金 (15%)
        uint256 boosterDeposit = (order.totalAmount * depositRate) / BASIS_POINTS;
        require(msg.value == boosterDeposit, "Incorrect deposit amount");
        
        // 更新订单
        order.booster = msg.sender;
        order.boosterDeposit = boosterDeposit;
        order.status = OrderStatus.Accepted;
        order.acceptedAt = block.timestamp;
        
        // 添加到代练订单列表
        boosterOrders[msg.sender].push(_orderId);
        
        emit OrderAccepted(_orderId, msg.sender, boosterDeposit);
    }
    
    /**
     * @dev 确认订单开始 - 玩家支付剩余85%并创建子合约
     * @param _orderId 订单ID
     */
    function confirmOrder(uint256 _orderId) 
        external 
        payable 
        nonReentrant 
        whenNotPaused 
        onlyValidOrder(_orderId) 
        onlyPlayer(_orderId) 
    {
        Order storage order = orders[_orderId];
        require(order.status == OrderStatus.Accepted, "Order not ready for confirmation");
        require(msg.value == order.remainingAmount, "Incorrect remaining amount");
        
        // 更新订单状态
        order.status = OrderStatus.Confirmed;
        order.confirmedAt = block.timestamp;
        
        // 创建子合约 (这里简化，实际可能需要部署新合约)
        // address childContract = deployChildContract(_orderId);
        // order.childContract = childContract;
        
        // 计算总锁定金额
        uint256 totalLocked = order.playerDeposit + order.boosterDeposit + order.remainingAmount;
        
        emit OrderConfirmed(_orderId, order.childContract, totalLocked);
    }
    
    /**
     * @dev 取消订单 - 处理不同阶段的取消逻辑
     * @param _orderId 订单ID
     */
    function cancelOrder(uint256 _orderId) 
        external 
        nonReentrant 
        whenNotPaused 
        onlyValidOrder(_orderId) 
        onlyPlayerOrBooster(_orderId) 
    {
        Order storage order = orders[_orderId];
        require(
            order.status == OrderStatus.Posted || order.status == OrderStatus.Accepted,
            "Cannot cancel order at this stage"
        );
        
        address cancelledBy = msg.sender;
        uint256 penaltyAmount = 0;
        address penaltyReceiver = address(0);
        
        if (order.status == OrderStatus.Posted) {
            // 代练未接单前，玩家可以取消，全额退还保证金
            require(cancelledBy == order.player, "Only player can cancel posted order");
            
            // 退还玩家保证金
            payable(order.player).transfer(order.playerDeposit);
            
        } else if (order.status == OrderStatus.Accepted) {
            // 代练已接单后，任一方取消都有惩罚
            if (cancelledBy == order.player) {
                // 玩家取消：扣除玩家保证金，2/3给代练，1/3给平台
                penaltyAmount = order.playerDeposit;
                uint256 toBooster = (penaltyAmount * penaltyToVictimRate) / BASIS_POINTS;
                uint256 toPlatform = penaltyAmount - toBooster;
                
                payable(order.booster).transfer(order.boosterDeposit + toBooster);
                payable(platformTreasury).transfer(toPlatform);
                penaltyReceiver = order.booster;
                
            } else if (cancelledBy == order.booster) {
                // 代练取消：扣除代练保证金，2/3给玩家，1/3给平台
                penaltyAmount = order.boosterDeposit;
                uint256 toPlayer = (penaltyAmount * penaltyToVictimRate) / BASIS_POINTS;
                uint256 toPlatform = penaltyAmount - toPlayer;
                
                payable(order.player).transfer(order.playerDeposit + toPlayer);
                payable(platformTreasury).transfer(toPlatform);
                penaltyReceiver = order.player;
            }
        }
        
        // 更新订单状态
        order.status = OrderStatus.Cancelled;
        
        emit OrderCancelled(_orderId, cancelledBy, penaltyAmount, penaltyReceiver);
    }
    
/**
     * @dev 完成订单 - 代练成功达成目标
     * @param _orderId 订单ID
     */
    function completeOrder(uint256 _orderId) 
        external 
        nonReentrant 
        onlyValidOrder(_orderId)
        onlyAuthorizedOrOwner  // 添加权限检查
    {
        Order storage order = orders[_orderId];
        require(
            order.status == OrderStatus.Confirmed || order.status == OrderStatus.InProgress,
            "Order not ready for completion"
        );
        
        // 生成当前交易哈希标识
        bytes32 currentTxHash = keccak256(abi.encodePacked(
            block.number,
            block.timestamp,
            msg.sender,
            _orderId,
            "COMPLETE"
        ));
        
        // 计算费用分配
        uint256 totalAmount = order.playerDeposit + order.boosterDeposit + order.remainingAmount;
        uint256 platformFee = (order.totalAmount * platformFeeRate) / BASIS_POINTS;
        uint256 boosterReward = totalAmount - platformFee;
        
        // 分配资金
        payable(platformTreasury).transfer(platformFee);
        payable(order.booster).transfer(boosterReward);
        
        // 更新订单状态
        order.status = OrderStatus.Completed;
        
        emit OrderCompleted(_orderId, platformFee, boosterReward, currentTxHash);
    }
    
    /**
     * @dev 订单失败 - 代练未达成目标
     * @param _orderId 订单ID
     */
    function failOrder(uint256 _orderId) 
        external 
        nonReentrant 
        onlyValidOrder(_orderId)
        onlyAuthorizedOrOwner  // 添加权限检查
    {
        Order storage order = orders[_orderId];
        require(
            order.status == OrderStatus.Confirmed || order.status == OrderStatus.InProgress,
            "Order not ready for failure"
        );
        
        // 生成当前交易哈希标识
        bytes32 currentTxHash = keccak256(abi.encodePacked(
            block.number,
            block.timestamp,
            msg.sender,
            _orderId,
            "FAIL"
        ));
        
        // 计算退款和惩罚分配
        uint256 playerRefund = order.playerDeposit + order.remainingAmount; // 玩家获得全额退款
        uint256 penaltyToPlayer = (order.boosterDeposit * penaltyToVictimRate) / BASIS_POINTS;
        uint256 penaltyToPlatform = order.boosterDeposit - penaltyToPlayer;
        
        // 分配资金
        payable(order.player).transfer(playerRefund + penaltyToPlayer);
        payable(platformTreasury).transfer(penaltyToPlatform);
        
        // 更新订单状态
        order.status = OrderStatus.Failed;
        
        emit OrderFailed(_orderId, playerRefund, penaltyToPlayer, penaltyToPlatform, currentTxHash);
    }
    
    // ==================== 查询函数 ====================
    
    /**
     * @dev 获取订单详情
     */
    function getOrder(uint256 _orderId) 
        external 
        view 
        onlyValidOrder(_orderId) 
        returns (Order memory) 
    {
        return orders[_orderId];
    }
    
    /**
     * @dev 获取玩家的订单列表
     */
    function getPlayerOrders(address _player) 
        external 
        view 
        returns (uint256[] memory) 
    {
        return playerOrders[_player];
    }
    
    /**
     * @dev 获取代练的订单列表
     */
    function getBoosterOrders(address _booster) 
        external 
        view 
        returns (uint256[] memory) 
    {
        return boosterOrders[_booster];
    }
    
    /**
     * @dev 计算保证金金额
     */
    function calculateDeposit(uint256 _totalAmount) 
        external 
        view 
        returns (uint256) 
    {
        return (_totalAmount * depositRate) / BASIS_POINTS;
    }
    
    // ==================== 管理员函数 ====================
    
    /**
     * @dev 设置平台费率
     */
    function setPlatformFeeRate(uint256 _newRate) external onlyOwner {
        require(_newRate <= 1000, "Fee rate cannot exceed 10%"); // 最大10%
        platformFeeRate = _newRate;
    }
    
    /**
     * @dev 设置保证金比例
     */
    function setDepositRate(uint256 _newRate) external onlyOwner {
        require(_newRate <= 3000, "Deposit rate cannot exceed 30%"); // 最大30%
        depositRate = _newRate;
    }
    
    /**
     * @dev 设置平台收益地址
     */
    function setPlatformTreasury(address _newTreasury) external onlyOwner {
        require(_newTreasury != address(0), "Invalid treasury address");
        platformTreasury = _newTreasury;
    }
    
    /**
     * @dev 紧急暂停
     */
    function pause() external onlyOwner {
        _pause();
    }
    
    /**
     * @dev 恢复运行
     */
    function unpause() external onlyOwner {
        _unpause();
    }
    
    /**
     * @dev 紧急提取（仅在紧急情况下使用）
     */
    function emergencyWithdraw() external onlyOwner {
        payable(owner()).transfer(address(this).balance);
    }
    
    // ==================== 接收ETH ====================
    
    receive() external payable {}
    
    fallback() external payable {}
}