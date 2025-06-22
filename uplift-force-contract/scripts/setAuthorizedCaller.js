const { ethers } = require("hardhat");

async function main() {
    // 获取部署者账户
    const [deployer] = await ethers.getSigners();
    
    console.log("Setting authorized caller with account:", deployer.address);
    console.log("Account balance:", ethers.formatEther(await ethers.provider.getBalance(deployer.address)), "ETH");
    
    // ==================== 配置区域 ====================
    
    // 主合约地址 - 请替换为你的实际合约地址
    const MAIN_CONTRACT_ADDRESS = "0xBa096e78F4D9b17959669b381CF8C1849EE525b5";
    
    // 要授权的地址 - 请替换为你想要授权的地址
    const AUTHORIZED_ADDRESS = "0xD86db2762D2abdf2e359C509B041bcEc5C993481";
    
    // 是否授权 (true: 授权, false: 取消授权)
    const IS_AUTHORIZED = true;
    
    // ==================== 验证配置 ====================
    
    if (MAIN_CONTRACT_ADDRESS === "YOUR_CONTRACT_ADDRESS_HERE") {
        console.error("❌ 错误: 请先设置 MAIN_CONTRACT_ADDRESS");
        process.exit(1);
    }
    
    if (!ethers.isAddress(MAIN_CONTRACT_ADDRESS)) {
        console.error("❌ 错误: 主合约地址格式无效");
        process.exit(1);
    }
    
    if (!ethers.isAddress(AUTHORIZED_ADDRESS)) {
        console.error("❌ 错误: 授权地址格式无效");
        process.exit(1);
    }
    
    console.log("\n==================== 配置信息 ====================");
    console.log("主合约地址:", MAIN_CONTRACT_ADDRESS);
    console.log("授权地址:", AUTHORIZED_ADDRESS);
    console.log("操作类型:", IS_AUTHORIZED ? "授权" : "取消授权");
    
    // ==================== 连接合约 ====================
    
    try {
        // 获取合约实例
        const BoostChainMainContract = await ethers.getContractFactory("BoostChainMainContract");
        const contract = BoostChainMainContract.attach(MAIN_CONTRACT_ADDRESS);
        
        console.log("\n✅ 合约连接成功");
        
        // 验证当前调用者是否为合约 owner
        const owner = await contract.owner();
        if (owner.toLowerCase() !== deployer.address.toLowerCase()) {
            console.error("❌ 错误: 当前账户不是合约的 owner");
            console.log("合约 owner:", owner);
            console.log("当前账户:", deployer.address);
            process.exit(1);
        }
        
        console.log("✅ Owner 验证通过");
        
        // 检查当前授权状态
        const currentStatus = await contract.authorizedCallers(AUTHORIZED_ADDRESS);
        console.log("\n当前授权状态:", currentStatus);
        
        if (currentStatus === IS_AUTHORIZED) {
            console.log(`⚠️  地址 ${AUTHORIZED_ADDRESS} 已经是${IS_AUTHORIZED ? "授权" : "未授权"}状态`);
            console.log("无需重复操作");
            return;
        }
        
        // ==================== 执行授权操作 ====================
        
        console.log(`\n开始${IS_AUTHORIZED ? "授权" : "取消授权"}操作...`);
        
        // 估算 gas
        const gasEstimate = await contract.setAuthorizedCaller.estimateGas(
            AUTHORIZED_ADDRESS,
            IS_AUTHORIZED
        );
        
        console.log("预估 Gas:", gasEstimate.toString());
        
        // 执行交易
        const tx = await contract.setAuthorizedCaller(
            AUTHORIZED_ADDRESS,
            IS_AUTHORIZED,
            {
                gasLimit: gasEstimate * 120n / 100n // 增加 20% gas 缓冲
            }
        );
        
        console.log("交易已发送，交易哈希:", tx.hash);
        console.log("等待交易确认...");
        
        // 等待交易确认
        const receipt = await tx.wait();
        
        console.log("\n🎉 交易成功确认!");
        console.log("区块号:", receipt.blockNumber);
        console.log("Gas 使用量:", receipt.gasUsed.toString());
        
        // 验证结果
        const newStatus = await contract.authorizedCallers(AUTHORIZED_ADDRESS);
        console.log("新的授权状态:", newStatus);
        
        // 查找事件日志
        const callerAuthorizedEvent = receipt.logs.find(log => {
            try {
                const parsed = contract.interface.parseLog(log);
                return parsed && parsed.name === "CallerAuthorized";
            } catch {
                return false;
            }
        });
        
        if (callerAuthorizedEvent) {
            const parsed = contract.interface.parseLog(callerAuthorizedEvent);
            console.log("\n📋 事件信息:");
            console.log("事件名称:", parsed.name);
            console.log("授权地址:", parsed.args.caller);
            console.log("授权状态:", parsed.args.authorized);
        }
        
        console.log("\n✅ 授权操作完成!");
        
    } catch (error) {
        console.error("\n❌ 操作失败:", error.message);
        
        // 特定错误处理
        if (error.message.includes("Unauthorized")) {
            console.error("提示: 请确保当前账户是合约的 owner");
        } else if (error.message.includes("Invalid caller address")) {
            console.error("提示: 请检查授权地址是否正确");
        } else if (error.message.includes("insufficient funds")) {
            console.error("提示: 账户余额不足，请充值后重试");
        }
        
        process.exitCode = 1;
    }
}

// ==================== 辅助函数 ====================

/**
 * 批量授权示例函数
 */
async function batchAuthorize() {
    const [deployer] = await ethers.getSigners();
    
    // 批量授权地址列表
    const addressesToAuthorize = [
        "0x1234567890123456789012345678901234567890",
        "0x0987654321098765432109876543210987654321"
    ];
    
    const MAIN_CONTRACT_ADDRESS = "YOUR_CONTRACT_ADDRESS_HERE";
    
    try {
        const BoostChainMainContract = await ethers.getContractFactory("BoostChainMainContract");
        const contract = BoostChainMainContract.attach(MAIN_CONTRACT_ADDRESS);
        
        console.log("开始批量授权...");
        
        const tx = await contract.setAuthorizedCallers(addressesToAuthorize, true);
        console.log("批量授权交易哈希:", tx.hash);
        
        const receipt = await tx.wait();
        console.log("批量授权完成，区块号:", receipt.blockNumber);
        
    } catch (error) {
        console.error("批量授权失败:", error.message);
    }
}

// 运行主函数
if (require.main === module) {
    main().catch((error) => {
        console.error(error);
        process.exitCode = 1;
    });
}

// 导出函数供其他脚本使用
module.exports = {
    main,
    batchAuthorize
};