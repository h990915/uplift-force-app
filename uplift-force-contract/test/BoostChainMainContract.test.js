const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("BoostChainMainContract - Create Order", function () {
  let contract;
  let owner;
  let player;
  let booster;
  let platformTreasury;
  let addrs;

  // 测试用的常量
  const TOTAL_AMOUNT = ethers.parseEther("1.0"); // 这已经是 BigInt
  const DEPOSIT_RATE = 1500n; // 15%
  const BASIS_POINTS = 10000n;
  
  beforeEach(async function () {
    // 获取测试账户
    [owner, player] = await ethers.getSigners();

    // 部署合约
    const BoostChainMainContract = await ethers.getContractFactory("BoostChainMainContract");
    contract = await BoostChainMainContract.deploy(owner.address);
    await contract.waitForDeployment();
    const contractAddress = await contract.getAddress();
    console.log('测试脚本合约部署地址：', contractAddress);
  });

  describe("创建订单测试", function () {
    
    it("应该成功创建订单并正确设置参数", async function () {
      const deadline = Math.floor(Date.now() / 1000) + 86400; // 24小时后
      const gameType = "League of Legends";
      const gameMode = "Ranked";
      const requirements = "From Silver to Gold";
      
      // 计算预期的保证金 (ethers v6 语法)
      const expectedDeposit = (TOTAL_AMOUNT * DEPOSIT_RATE) / BASIS_POINTS;
      
      // 创建订单
      const tx = await contract.connect(player).createOrder(
        TOTAL_AMOUNT,
        deadline,
        gameType,
        gameMode,
        requirements,
        { value: expectedDeposit }
      );

      console.log("交易哈希:", tx.hash);
      // const receipt = await tx.wait();
      // console.log("\n🎯 用于后端验证的交易哈希:", receipt.hash);

      // 验证交易成功
      await expect(tx).to.emit(contract, "OrderCreated");
      
      // 获取创建的订单
      const order = await contract.getOrder(1);
      
      // 验证订单参数
      expect(order.orderId).to.equal(1);
      expect(order.player).to.equal(player.address);
      expect(order.totalAmount).to.equal(TOTAL_AMOUNT);
      expect(order.playerDeposit).to.equal(expectedDeposit);
      expect(order.gameType).to.equal(gameType);
      expect(order.gameMode).to.equal(gameMode);
      expect(order.requirements).to.equal(requirements);
      expect(order.status).to.equal(0); // OrderStatus.Posted
    });

    it("应该正确计算和收取保证金", async function () {
      const deadline = Math.floor(Date.now() / 1000) + 86400;
      const expectedDeposit = TOTAL_AMOUNT.mul(DEPOSIT_RATE).div(BASIS_POINTS);
      
      // 记录玩家初始余额
      const initialBalance = await player.getBalance();
      
      const tx = await contract.connect(player).createOrder(
        TOTAL_AMOUNT,
        deadline,
        "DOTA2",
        "Ranked",
        "Calibration matches",
        { value: expectedDeposit }
      );
      
      const receipt = await tx.wait();
      const gasUsed = receipt.gasUsed.mul(receipt.effectiveGasPrice);
      
      // 验证玩家余额减少了保证金 + gas费
      const finalBalance = await player.getBalance();
      const expectedBalance = initialBalance.sub(expectedDeposit).sub(gasUsed);
      expect(finalBalance).to.equal(expectedBalance);
      
      // 验证合约余额增加了保证金
      expect(await ethers.provider.getBalance(contract.address)).to.equal(expectedDeposit);
    });

    it("应该正确触发 OrderCreated 事件", async function () {
      const deadline = Math.floor(Date.now() / 1000) + 86400;
      const gameType = "Valorant";
      const gameMode = "Competitive";
      const expectedDeposit = TOTAL_AMOUNT.mul(DEPOSIT_RATE).div(BASIS_POINTS);
      
      // 验证事件触发
      await expect(contract.connect(player).createOrder(
        TOTAL_AMOUNT,
        deadline,
        gameType,
        gameMode,
        "Rank up to Immortal",
        { value: expectedDeposit }
      )).to.emit(contract, "OrderCreated")
        .withArgs(
          1, // orderId
          player.address,
          TOTAL_AMOUNT,
          expectedDeposit,
          gameType,
          gameMode
        );
    });

    it("应该将订单添加到玩家订单列表", async function () {
      const deadline = Math.floor(Date.now() / 1000) + 86400;
      const expectedDeposit = TOTAL_AMOUNT.mul(DEPOSIT_RATE).div(BASIS_POINTS);
      
      // 创建第一个订单
      await contract.connect(player).createOrder(
        TOTAL_AMOUNT,
        deadline,
        "CS2",
        "Premier",
        "Reach Global Elite",
        { value: expectedDeposit }
      );
      
      // 创建第二个订单
      await contract.connect(player).createOrder(
        TOTAL_AMOUNT.mul(2),
        deadline,
        "Overwatch",
        "Competitive",
        "Reach Grandmaster",
        { value: TOTAL_AMOUNT.mul(2).mul(DEPOSIT_RATE).div(BASIS_POINTS) }
      );
      
      // 验证玩家订单列表
      const playerOrderList = await contract.getPlayerOrders(player.address);
      expect(playerOrderList.length).to.equal(2);
      expect(playerOrderList[0]).to.equal(1);
      expect(playerOrderList[1]).to.equal(2);
    });

    it("应该正确使用 calculateDeposit 函数", async function () {
      const testAmount = ethers.parseEther("2.5");
      const calculatedDeposit = await contract.calculateDeposit(testAmount);
      const expectedDeposit = testAmount.mul(DEPOSIT_RATE).div(BASIS_POINTS);
      
      expect(calculatedDeposit).to.equal(expectedDeposit);
    });
  });

  describe("创建订单失败场景", function () {
    
    it("应该在金额为0时失败", async function () {
      const deadline = Math.floor(Date.now() / 1000) + 86400;
      
      await expect(contract.connect(player).createOrder(
        0,
        deadline,
        "Game",
        "Mode",
        "Requirements",
        { value: 0 }
      )).to.be.revertedWith("Total amount must be greater than 0");
    });

    it("应该在截止时间已过时失败", async function () {
      const pastDeadline = Math.floor(Date.now() / 1000) - 3600; // 1小时前
      const expectedDeposit = TOTAL_AMOUNT.mul(DEPOSIT_RATE).div(BASIS_POINTS);
      
      await expect(contract.connect(player).createOrder(
        TOTAL_AMOUNT,
        pastDeadline,
        "Game",
        "Mode",
        "Requirements",
        { value: expectedDeposit }
      )).to.be.revertedWith("Deadline must be in the future");
    });

    it("应该在游戏类型为空时失败", async function () {
      const deadline = Math.floor(Date.now() / 1000) + 86400;
      const expectedDeposit = TOTAL_AMOUNT.mul(DEPOSIT_RATE).div(BASIS_POINTS);
      
      await expect(contract.connect(player).createOrder(
        TOTAL_AMOUNT,
        deadline,
        "", // 空字符串
        "Mode",
        "Requirements",
        { value: expectedDeposit }
      )).to.be.revertedWith("Game type cannot be empty");
    });

    it("应该在游戏模式为空时失败", async function () {
      const deadline = Math.floor(Date.now() / 1000) + 86400;
      const expectedDeposit = TOTAL_AMOUNT.mul(DEPOSIT_RATE).div(BASIS_POINTS);
      
      await expect(contract.connect(player).createOrder(
        TOTAL_AMOUNT,
        deadline,
        "Game",
        "", // 空字符串
        "Requirements",
        { value: expectedDeposit }
      )).to.be.revertedWith("Game mode cannot be empty");
    });

    it("应该在保证金金额不正确时失败", async function () {
      const deadline = Math.floor(Date.now() / 1000) + 86400;
      const expectedDeposit = TOTAL_AMOUNT.mul(DEPOSIT_RATE).div(BASIS_POINTS);
      const incorrectDeposit = expectedDeposit.sub(1); // 少1 wei
      
      await expect(contract.connect(player).createOrder(
        TOTAL_AMOUNT,
        deadline,
        "Game",
        "Mode",
        "Requirements",
        { value: incorrectDeposit }
      )).to.be.revertedWith("Incorrect deposit amount");
    });

    it("应该在合约暂停时失败", async function () {
      // 暂停合约
      await contract.connect(owner).pause();
      
      const deadline = Math.floor(Date.now() / 1000) + 86400;
      const expectedDeposit = TOTAL_AMOUNT.mul(DEPOSIT_RATE).div(BASIS_POINTS);
      
      await expect(contract.connect(player).createOrder(
        TOTAL_AMOUNT,
        deadline,
        "Game",
        "Mode",
        "Requirements",
        { value: expectedDeposit }
      )).to.be.revertedWith("Pausable: paused");
    });
  });

  describe("边界情况测试", function () {
    
    it("应该处理非常大的金额", async function () {
      const largeAmount = ethers.parseEther("1000");
      const deadline = Math.floor(Date.now() / 1000) + 86400;
      const expectedDeposit = largeAmount.mul(DEPOSIT_RATE).div(BASIS_POINTS);
      
      await expect(contract.connect(player).createOrder(
        largeAmount,
        deadline,
        "Game",
        "Mode",
        "Requirements",
        { value: expectedDeposit }
      )).to.not.be.reverted;
    });

    it("应该处理最小金额（1 wei）", async function () {
      const minAmount = 1;
      const deadline = Math.floor(Date.now() / 1000) + 86400;
      const expectedDeposit = Math.floor(minAmount * DEPOSIT_RATE / BASIS_POINTS); // 0
      
      await expect(contract.connect(player).createOrder(
        minAmount,
        deadline,
        "Game",
        "Mode",
        "Requirements",
        { value: expectedDeposit }
      )).to.not.be.reverted;
    });

    it("应该处理很长的字符串", async function () {
      const deadline = Math.floor(Date.now() / 1000) + 86400;
      const expectedDeposit = TOTAL_AMOUNT.mul(DEPOSIT_RATE).div(BASIS_POINTS);
      const longString = "A".repeat(1000);
      
      await expect(contract.connect(player).createOrder(
        TOTAL_AMOUNT,
        deadline,
        longString,
        longString,
        longString,
        { value: expectedDeposit }
      )).to.not.be.reverted;
    });
  });
});