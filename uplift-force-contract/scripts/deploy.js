const { ethers } = require("hardhat");

async function main() {
  // 获取所有 signer
  const signers = await ethers.getSigners();

  const deployer = signers[0];

  console.log("Deploying contracts with the account:", deployer.address);
  console.log("Account balance:", (await ethers.provider.getBalance(deployer.address)).toString());

  // 设置平台收益地址
  const platformTreasury = deployer.address;

  const BoostChainMainContract = await ethers.getContractFactory("BoostChainMainContract", deployer);
  const contract = await BoostChainMainContract.deploy(platformTreasury);

  await contract.waitForDeployment();

  console.log("BoostChainMainContract deployed to:", contract.target);
  console.log("Platform Treasury:", platformTreasury);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});