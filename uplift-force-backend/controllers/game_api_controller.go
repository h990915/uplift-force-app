package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"uplift-force-backend/utils"
)

type GameApiHandler struct {
	RiotClient *utils.RiotClient
}

// NewGameApiHandler 创建新的处理器实例
func NewGameApiHandler(riotClient *utils.RiotClient) *GameApiHandler {
	return &GameApiHandler{
		RiotClient: riotClient,
	}
}

// GetSummonerPUUID 获取召唤师PUUID
func (h *GameApiHandler) GetSummonerPUUID(c *gin.Context) {
	summonerName := c.Query("characterName")
	tagLine := c.Query("tagLine")
	region := c.Query("region")

	riotClient, err := utils.NewRiotClient(region)
	if err != nil {
		log.Fatal("创建客户端失败:", err)
	}
	data, err := riotClient.Call(utils.SummonerByName, map[string]string{
		"gameName": summonerName,
		"tagLine":  tagLine,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		c.JSON(500, gin.H{"error": "解析响应失败"})
		return
	}

	c.JSON(200, result)
}

// 获取玩家游戏数据
func (h *GameApiHandler) GetLeagueEntries(c *gin.Context) {
	encryptedPUUID := c.Query("encryptedPUUID")
	region := c.Query("region")

	riotClient, err := utils.NewRiotClient(region)
	if err != nil {
		log.Fatal("创建客户端失败:", err)
	}
	data, err := riotClient.Call(utils.LeagueEntries, map[string]string{
		"encryptedPUUID": encryptedPUUID,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		c.JSON(500, gin.H{"error": "解析响应失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data":   result,
		"count":  len(result),
	})
}

// GetSummonerProfileWithRank 获取召唤师资料和排位信息
// @Summary 获取召唤师资料和排位信息
// @Description 根据游戏名称和标签获取召唤师的基本信息和排位数据，系统会根据tagLine自动识别对应的游戏地区
// @Tags 召唤师管理
// @Accept json
// @Produce json
// @Param characterName query string true "召唤师游戏名称，支持中文、英文等多种字符" minlength(1) maxlength(50) example("Hide on bush")
// @Param tagLine query string true "召唤师标签，用于识别游戏地区，大小写不敏感" minlength(1) maxlength(10) example("KR1") Enums(KR,KR1,NA,NA1,EUW,EUNE,JP,JP1,BR,BR1,LAN,LA1,LAS,LA2,TR,TR1,RU,RUS,OCE,OC1,AUS,PH,PH2,SG,SG2,MY,TH,TH2,TW,TW2,VN,VN2)
// @Success 200 {object} object{status=string,data=object{summoner=object{gameName=string,puuid=string,tagLine=string},leagueEntries=[]object{freshBlood=bool,hotStreak=bool,inactive=bool,leagueId=string,leaguePoints=int,losses=int,puuid=string,queueType=string,rank=string,summonerId=string,tier=string,veteran=bool,wins=int},leagueCount=int}} "获取成功"
// @Failure 400 {object} object{error=string} "参数错误：characterName或tagLine为空"
// @Failure 404 {object} object{error=string} "召唤师不存在"
// @Failure 500 {object} object{error=string} "服务器内部错误：API调用失败或数据解析失败"
// @Router /api/v1/summoner/getWithRank [get]
func (h *GameApiHandler) GetSummonerProfileWithRank(c *gin.Context) {
	summonerName := c.Query("characterName")
	tagLine := c.Query("tagLine")
	region := utils.GetRegionByTagLine(tagLine)

	riotClient, err := utils.NewRiotClient(region)
	if err != nil {
		log.Fatal("创建客户端失败:", err)
	}

	if region == "kr" {
		tagLine = "kr1"
	}
	// 第一步：获取召唤师PUUID
	summonerData, err := riotClient.Call(utils.SummonerByName, map[string]string{
		"gameName": summonerName,
		"tagLine":  tagLine,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "获取召唤师信息失败: " + err.Error()})
		return
	}

	var summonerResult map[string]interface{}
	if err := json.Unmarshal(summonerData, &summonerResult); err != nil {
		c.JSON(500, gin.H{"error": "解析召唤师信息失败"})
		return
	}

	// 从召唤师信息中提取PUUID
	encryptedPUUID, ok := summonerResult["puuid"].(string)
	if !ok {
		c.JSON(500, gin.H{"error": "未能获取到PUUID"})
		return
	}

	// 第二步：使用PUUID获取排位信息
	leagueData, err := riotClient.Call(utils.LeagueEntries, map[string]string{
		"encryptedPUUID": encryptedPUUID,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "获取排位信息失败: " + err.Error()})
		return
	}

	var leagueResult []map[string]interface{}
	if err := json.Unmarshal(leagueData, &leagueResult); err != nil {
		c.JSON(500, gin.H{"error": "解析排位信息失败: " + err.Error()})
		return
	}

	// 合并返回结果
	c.JSON(200, gin.H{
		"status": "success",
		"data": gin.H{
			"summoner":      summonerResult,
			"leagueEntries": leagueResult,
			"leagueCount":   len(leagueResult),
		},
	})
}
