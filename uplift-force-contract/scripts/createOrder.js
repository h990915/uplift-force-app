const { ethers } = require("hardhat");

async function main() {
    console.log("=== 开始创建订单 ===\n");
    
    try {
        // 1. 获取签名者账户
        const [owner, player] = await ethers.getSigners();
        console.log("部署者地址:", owner.address);
        console.log("玩家地址:", player.address);
        
        // 显示账户余额
        const playerBalance = await ethers.provider.getBalance(player.address);
        console.log("玩家余额:", ethers.formatEther(playerBalance), "ETH\n");
        
        // 2. 连接到已部署的合约
        // ⚠️ 替换为你实际部署的合约地址
        const contractAddress = "0x31122497CbaF01cD1C9b16E0D9F63b3223e4e2FA"; // 从 deploy 脚本输出中获取
        
        console.log("连接到合约地址:", contractAddress);
        
        const BoostChainMainContract = await ethers.getContractFactory("BoostChainMainContract");
        const contract = BoostChainMainContract.attach(contractAddress);
        
        // 3. 设置订单参数 (模仿测试用例)
        const TOTAL_AMOUNT = ethers.parseEther("0.5"); 
        const DEPOSIT_RATE = 1500; // 15%
        const BASIS_POINTS = 10000;
        
        const deadline = Math.floor(Date.now() / 1000) + 86400; // 24小时后
        const gameType = "League of Legends";
        const gameMode = "Ranked";
        const requirements = "From Silver to Gold";
        
        // 计算预期的保证金
        const expectedDeposit = (TOTAL_AMOUNT * BigInt(DEPOSIT_RATE)) / BigInt(BASIS_POINTS);
        
        console.log("\n=== 订单参数 ===");
        console.log("总金额:", ethers.formatEther(TOTAL_AMOUNT), "ETH");
        console.log("保证金:", ethers.formatEther(expectedDeposit), "ETH");
        console.log("游戏类型:", gameType);
        console.log("游戏模式:", gameMode);
        console.log("要求:", requirements);
        console.log("截止时间:", new Date(deadline * 1000).toLocaleString());
        
        // 4. 创建订单
        console.log("\n=== 创建订单 ===");
        console.log("发送交易...");
        
        const tx = await contract.connect(player).createOrder(
            TOTAL_AMOUNT,
            deadline,
            gameType,
            gameMode,
            requirements,
            { 
                value: expectedDeposit,
                gasLimit: 500000 // 设置Gas限制
            }
        );
        
        console.log("✅ 交易已发送!");
        console.log("交易哈希:", tx.hash);
        
        // 5. 等待交易确认
        console.log("\n等待交易确认...");
        const receipt = await tx.wait();
        
        console.log("✅ 交易已确认!");
        console.log("区块号:", receipt.blockNumber);
        console.log("Gas 使用:", receipt.gasUsed.toString());
        console.log("交易状态:", receipt.status === 1 ? "成功" : "失败");
        
    } catch (error) {
        console.error("\n❌ 创建订单失败:", error.message);
        
        if (error.reason) {
            console.error("失败原因:", error.reason);
        }
        
        if (error.code) {
            console.error("错误代码:", error.code);
        }
        
        // 常见错误提示
        if (error.message.includes("insufficient funds")) {
            console.error("💡 提示: 账户余额不足，请确保有足够的 ETH");
        } else if (error.message.includes("contract not deployed")) {
            console.error("💡 提示: 合约未部署，请先运行 deploy 脚本");
        } else if (error.message.includes("invalid address")) {
            console.error("💡 提示: 合约地址无效，请检查 contractAddress 变量");
        }
        
        process.exit(1);
    }
}


main()
    .then(() => {
        console.log("\n✅ 脚本执行完成");
        process.exit(0);
    })
    .catch((error) => {
        console.error("\n💥 未捕获的错误:", error);
        process.exit(1);
    });
