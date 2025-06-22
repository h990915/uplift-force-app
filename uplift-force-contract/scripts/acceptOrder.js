// scripts/acceptOrder.js
const { ethers } = require("hardhat");

async function main() {
        const [booster] = await ethers.getSigners();
        console.log("代练地址:", booster.address);

        // 连接合约
        const contractAddress = "0x31122497CbaF01cD1C9b16E0D9F63b3223e4e2FA"; // 从 deploy 脚本输出中获取
        console.log("连接到合约地址:", contractAddress);
        const BoostChainMainContract = await ethers.getContractFactory("BoostChainMainContract");
        const contract = BoostChainMainContract.attach(contractAddress);

        // 检查当前订单数量
        const orderCount = await contract.orderCounter();
        console.log("📊 当前订单总数:", orderCount.toString());
                if (orderCount === 0n) {
            console.log("❌ 没有可接受的订单，请先创建订单");
            return;
        }
        // 使用最新的订单ID进行测试
        const orderId = orderCount;
        console.log("🔍 准备接受订单ID:", orderId.toString());
        
        // 获取订单信息
        const orderBefore = await contract.getOrder(orderId);
        console.log("\n📋 订单信息:");
        console.log("• 玩家:", orderBefore.player);
        console.log("• 状态:", getStatusName(orderBefore.status));
        console.log("• 总金额:", ethers.formatEther(orderBefore.totalAmount), "ETH");
                // 检查订单状态
        if (orderBefore.status !== 0n) { // 0 = Posted
            console.log("❌ 订单状态不是 Posted，无法接受");
            console.log("当前状态:", getStatusName(orderBefore.status));
            return;
        }
                // 计算代练保证金
        const totalAmount = orderBefore.totalAmount;
        const depositRate = 1500; // 15%
        const basisPoints = 10000;
        const boosterDeposit = (totalAmount * BigInt(depositRate)) / BigInt(basisPoints);
        
        console.log("\n💰 保证金计算:");
        console.log("• 订单总金额:", ethers.formatEther(totalAmount), "ETH");
        console.log("• 代练保证金:", ethers.formatEther(boosterDeposit), "ETH");
        console.log("• 保证金比例: 15%");
        
        // 检查代练余额
        const boosterBalance = await ethers.provider.getBalance(booster.address);
        console.log("\n💳 代练账户余额:", ethers.formatEther(boosterBalance), "ETH");
        
        if (boosterBalance < boosterDeposit) {
            console.log("❌ 代练账户余额不足，无法支付保证金");
            return;
        }
        
        // 执行接受订单
        console.log("\n🚀 发送接受订单交易...");
        const tx = await contract.connect(booster).acceptOrder(orderId, {
            value: boosterDeposit,
            gasLimit: 300000
        });
        
        console.log("⏳ 交易已发送，等待确认...");
        console.log("📄 交易哈希:", tx.hash);
        
        // 等待交易确认
        const receipt = await tx.wait();
        console.log("✅ 交易已确认!");
}

function getStatusName(status) {
    const statusNames = [
        "Posted",     // 0
        "Accepted",   // 1
        "Confirmed",  // 2
        "Completed",  // 3
        "Cancelled",  // 4
        "Disputed"    // 5
    ];
    return statusNames[status] || "Unknown";
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