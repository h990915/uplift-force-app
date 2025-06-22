// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {ILogAutomation, Log} from "@chainlink/contracts/src/v0.8/automation/interfaces/ILogAutomation.sol";
import {ConfirmedOwner} from "@chainlink/contracts/src/v0.8/shared/access/ConfirmedOwner.sol";

interface IBoostChainMainContract {
    function completeOrder(uint256 _orderId) external;
    function failOrder(uint256 _orderId) external;
}

contract RiotOrderProcessor is ILogAutomation, ConfirmedOwner {
    
    address public mainContractAddress;
    
    // 防止重复处理
    mapping(uint256 => bool) public processedOrders;

    event OrderProcessed(uint256 indexed orderId, bool success, string target, string tier);
    event ProcessingFailed(uint256 indexed orderId, string reason);

    constructor() ConfirmedOwner(msg.sender) {}

    function setMainContract(address _mainContractAddress) external onlyOwner {
        require(_mainContractAddress != address(0), "Invalid address");
        mainContractAddress = _mainContractAddress;
    }

    function checkLog(
        Log calldata log,
        bytes memory
    ) external pure override returns (bool upkeepNeeded, bytes memory performData) {
        // TierVerificationResult 事件的签名
        bytes32 tierVerificationSignature = keccak256("TierVerificationResult(uint256,string,string)");
        
        if (log.topics[0] == tierVerificationSignature) {
            // 解码 orderId (第一个 indexed 参数)
            uint256 orderId = uint256(log.topics[1]);
            
            // 解码 target 和 tier (从 data 中)
            (string memory target, string memory tier) = abi.decode(log.data, (string, string));
            
            upkeepNeeded = true;
            performData = abi.encode(orderId, target, tier);
        }
        
        return (upkeepNeeded, performData);
    }

    function performUpkeep(bytes calldata performData) external override {
        (uint256 orderId, string memory target, string memory tier) = abi.decode(
            performData, 
            (uint256, string, string)
        );
        
        // 防止重复处理
        // require(!processedOrders[orderId], "Order already processed");
        // require(mainContractAddress != address(0), "Main contract not set");
        
        processedOrders[orderId] = true;
        
        // 简单字符串比较：相等为成功，不等为失败
        bool success = _compareStrings(target, tier);
        
        IBoostChainMainContract mainContract = IBoostChainMainContract(mainContractAddress);
        
        if (success) {
            try mainContract.completeOrder(orderId) {
                emit OrderProcessed(orderId, true, target, tier);
            } catch Error(string memory reason) {
                emit ProcessingFailed(orderId, reason);
            } catch {
                emit ProcessingFailed(orderId, "Complete order failed");
            }
        } else {
            try mainContract.failOrder(orderId) {
                emit OrderProcessed(orderId, false, target, tier);
            } catch Error(string memory reason) {
                emit ProcessingFailed(orderId, reason);
            } catch {
                emit ProcessingFailed(orderId, "Fail order failed");
            }
        }
    }

    function _compareStrings(string memory a, string memory b) private pure returns (bool) {
        return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
    }

    // 手动处理函数（应急使用）
    function manualProcess(uint256 orderId, string calldata target, string calldata tier) external onlyOwner {
        require(!processedOrders[orderId], "Order already processed");
        require(mainContractAddress != address(0), "Main contract not set");
        
        processedOrders[orderId] = true;
        
        bool success = _compareStrings(target, tier);
        IBoostChainMainContract mainContract = IBoostChainMainContract(mainContractAddress);
        
        if (success) {
            mainContract.completeOrder(orderId);
        } else {
            mainContract.failOrder(orderId);
        }
        
        emit OrderProcessed(orderId, success, target, tier);
    }
}