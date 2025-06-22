// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {FunctionsClient} from "@chainlink/contracts/src/v0.8/functions/v1_0_0/FunctionsClient.sol";
import {ConfirmedOwner} from "@chainlink/contracts/src/v0.8/shared/access/ConfirmedOwner.sol";
import {FunctionsRequest} from "@chainlink/contracts/src/v0.8/functions/v1_0_0/libraries/FunctionsRequest.sol";

contract RiotDataConsumer is FunctionsClient, ConfirmedOwner {
    using FunctionsRequest for FunctionsRequest.Request;

    // ==================== 状态变量 ====================
    
    bytes32 public s_lastRequestId;
    string public s_lastResponse;
    string public s_lastError;
    
    // Chainlink Functions配置
    uint8 public secretsSlotId;
    uint64 public secretsVersion;
    uint64 public subId;
    uint32 private s_gasLimit = 300000;
    bytes32 private DON_ID = 0x66756e2d657468657265756d2d7365706f6c69612d3100000000000000000000;
    address public constant ROUTER = 0xb83E47C2bC239B3bf370bc41e1459A34b41238D0;

    // ==================== 数据结构 ====================
    
    struct RequestData {
        uint256 orderId;
        string target;
    }
    
    mapping(bytes32 => RequestData) public requestData;

    // ==================== JavaScript源代码 (修改返回格式) ====================
    
    string public constant SOURCE = 
        "const encryptedPUUID = args[0]; "
        "const region = args[1] || 'kr'; "
        "const target = args[2]; "
        "const orderId = args[3]; "
        "const apiKey = secrets.apiKey; "
        "if(!apiKey) {throw Error('apiKey not found in secrets');} "
        "if(!encryptedPUUID) {throw Error('encryptedPUUID is required');} "
        "const apiUrl = `https://${region}.api.riotgames.com/lol/league/v4/entries/by-puuid/${encryptedPUUID}`; "
        "const riotRequest = Functions.makeHttpRequest({url: apiUrl, method: 'GET', headers: {'X-Riot-Token': apiKey, 'Accept': 'application/json'}, timeout: 20000}); "
        "const riotResponse = await riotRequest; "
        "if(riotResponse.error) {throw Error(`Request failed: ${riotResponse.error}`);} "
        "if(riotResponse.status !== 200) {throw Error(`API request failed with status ${riotResponse.status}`);} "
        "const leagueData = riotResponse.data; "
        "let tier = 'UNRANKED'; "
        "if(leagueData && leagueData.length > 0) { "
        "   let soloRankData = leagueData.find(entry => entry.queueType === 'RANKED_SOLO_5x5'); "
        "   const rankData = soloRankData || leagueData[0]; "
        "   tier = rankData.tier || 'UNRANKED'; "
        "} "
        "const result = `${orderId},${target},${tier}`; "
        "return Functions.encodeString(result);";

    // ==================== 事件定义 ====================
    
    event RequestSent(bytes32 indexed requestId, uint256 indexed orderId);
    
    // 核心事件：为 Automation 提供简化数据
    event TierVerificationResult(
        uint256 indexed orderId,
        string target,
        string tier
    );

    event FunctionError(uint256 indexed orderId, string errorMessage);

    // ==================== 构造函数 ====================
    
    constructor() FunctionsClient(ROUTER) ConfirmedOwner(msg.sender) {}

    // ==================== 配置函数 ====================
    
    function setConfig(uint8 _secretsSlotId, uint64 _secretsVersion, uint64 _subId) public onlyOwner {
        secretsSlotId = _secretsSlotId;
        secretsVersion = _secretsVersion;
        subId = _subId;
    }

    // ==================== 核心功能函数 ====================
    
    function requestPlayerData(
        string memory _puuid,
        string memory _region,
        string memory _target,
        uint256 _orderId  
    ) external returns (bytes32 requestId) {
        require(bytes(_puuid).length > 0, "PUUID required");
        require(subId > 0, "Subscription not configured");
        
        FunctionsRequest.Request memory req;
        req.initializeRequestForInlineJavaScript(SOURCE);

        if (secretsVersion > 0) {
            req.addDONHostedSecrets(secretsSlotId, secretsVersion);
        }
        
        // 修改：传递 orderId 给 JS
        string[] memory args = new string[](4);  
        args[0] = _puuid;
        args[1] = _region;
        args[2] = _target;
        args[3] = _uint256ToString(_orderId);  // 转换为字符串
        req.setArgs(args);
        
        s_lastRequestId = _sendRequest(
            req.encodeCBOR(),
            subId,
            s_gasLimit,
            DON_ID
        );

        // 存储请求数据
        requestData[s_lastRequestId] = RequestData({
            orderId: _orderId,
            target: _target
        });
        
        emit RequestSent(s_lastRequestId, _orderId);
        return s_lastRequestId;
    }

    function fulfillRequest(
        bytes32 requestId,
        bytes memory response,
        bytes memory err
    ) internal override {
        s_lastResponse = string(response);
        
        RequestData memory reqData = requestData[requestId];
        
        if (err.length > 0) {
            s_lastError = string(err);
            emit FunctionError(reqData.orderId, string(err));
        } else {
            // 解析简化的响应: "orderId,target,tier"
            string memory responseStr = string(response);
            (uint256 orderId, string memory target, string memory tier) = _parseSimpleResponse(responseStr);
            
            // 验证 orderId 匹配
            if (orderId == reqData.orderId) {
                emit TierVerificationResult(orderId, target, tier);
            } else {
                emit FunctionError(reqData.orderId, "Order ID mismatch");
            }
        }
        
        // 清理
        delete requestData[requestId];
    }

    // ==================== 辅助函数 ====================
    
    function _parseSimpleResponse(string memory response) 
        private 
        pure 
        returns (uint256 orderId, string memory target, string memory tier) 
    {
        bytes memory responseBytes = bytes(response);
        
        // 查找两个逗号的位置
        uint256 firstComma = 0;
        uint256 secondComma = 0;
        uint256 commaCount = 0;
        
        for (uint256 i = 0; i < responseBytes.length; i++) {
            if (responseBytes[i] == 0x2C) { // 逗号的ASCII码
                commaCount++;
                if (commaCount == 1) {
                    firstComma = i;
                } else if (commaCount == 2) {
                    secondComma = i;
                    break;
                }
            }
        }
        
        if (commaCount < 2) {
            return (0, "", "");
        }
        
        // 提取 orderId
        orderId = _stringToUint256(_substring(response, 0, firstComma));
        
        // 提取 target
        target = _substring(response, firstComma + 1, secondComma);
        
        // 提取 tier
        tier = _substring(response, secondComma + 1, responseBytes.length);
    }
    
    function _substring(string memory str, uint256 start, uint256 end) 
        private 
        pure 
        returns (string memory) 
    {
        bytes memory strBytes = bytes(str);
        bytes memory result = new bytes(end - start);
        
        for (uint256 i = 0; i < end - start; i++) {
            result[i] = strBytes[start + i];
        }
        
        return string(result);
    }
    
    function _uint256ToString(uint256 value) private pure returns (string memory) {
        if (value == 0) {
            return "0";
        }
        uint256 temp = value;
        uint256 digits;
        while (temp != 0) {
            digits++;
            temp /= 10;
        }
        bytes memory buffer = new bytes(digits);
        while (value != 0) {
            digits -= 1;
            buffer[digits] = bytes1(uint8(48 + uint256(value % 10)));
            value /= 10;
        }
        return string(buffer);
    }
    
    function _stringToUint256(string memory str) private pure returns (uint256) {
        bytes memory b = bytes(str);
        uint256 result = 0;
        for (uint256 i = 0; i < b.length; i++) {
            if (b[i] >= 0x30 && b[i] <= 0x39) {
                result = result * 10 + (uint8(b[i]) - 48);
            }
        }
        return result;
    }
}