require("@nomicfoundation/hardhat-toolbox");
require("@nomicfoundation/hardhat-verify");
require("dotenv").config();
// require("./task");
// require('hardhat-deploy');
// require("@nomicfoundation/hardhat-ethers");
// require("hardhat-deploy");
// require("hardhat-deploy-ethers");

const { ProxyAgent, setGlobalDispatcher } = require("undici");
const proxyAgent = new ProxyAgent("http://127.0.0.1:1083");
setGlobalDispatcher(proxyAgent);

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    solidity: {
    version: "0.8.28",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      },
      viaIR: true  // 启用 IR 编译器
    }
  },
  networks: {
    hardhat: {
      accounts: [
        {
          privateKey: process.env.PRIVATE_KEY_1,
          balance: "10000000000000000000000"
        },
        {
          privateKey: process.env.PRIVATE_KEY_2,
          balance: "10000000000000000000000"
        }
      ]
    },
    sepolia: {
      url: process.env.SEPOLIA_URL,
      accounts: [process.env.PRIVATE_KEY_1, process.env.PRIVATE_KEY_2],
      chainId: 11155111
    },
      avalancheFuji: {
      url: process.env.AVALANCHE_FUJI_URL || "https://api.avax-test.network/ext/bc/C/rpc",
      accounts: [process.env.PRIVATE_KEY_1, process.env.PRIVATE_KEY_2],
      chainId: 43113
    }
  },
  namedAccounts: {
    firstAccount: {
      default: 0
    },
    secondAccount: {
      default: 1
    },
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY
  },
  mocha: {
    timeout: 400000
  },
};
