package services

import (
	"encoding/json"
	"fmt"
	"uplift-force-backend/utils"
)

type RiotService struct {
	client *utils.RiotClient
}

func NewRiotService(client *utils.RiotClient) *RiotService {
	return &RiotService{
		client: client,
	}
}

func (s *RiotService) GetSummonerByName(summonerName string) (*map[string]interface{}, error) {
	// 使用Call方法获取原始数据
	data, err := s.client.Call(
		utils.SummonerByName,
		map[string]string{"summonerName": summonerName},
	)
	if err != nil {
		return nil, fmt.Errorf("获取召唤师信息失败: %w", err)
	}

	// 手动解析JSON
	var summoner map[string]interface{}
	if err := json.Unmarshal(data, &summoner); err != nil {
		return nil, fmt.Errorf("解析召唤师信息失败: %w", err)
	}

	return &summoner, nil
}

func (s *RiotService) GetMatchHistory(puuid string, count int) ([]string, error) {
	params := map[string]string{
		"puuid": puuid,
	}
	if count > 0 {
		params["count"] = fmt.Sprintf("%d", count)
	}

	data, err := s.client.Call(utils.MatchHistory, params)
	if err != nil {
		return nil, fmt.Errorf("获取比赛历史失败: %w", err)
	}

	var matches []string
	if err := json.Unmarshal(data, &matches); err != nil {
		return nil, fmt.Errorf("解析比赛历史失败: %w", err)
	}

	return matches, nil
}
