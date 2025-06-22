package utils

import (
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"os"
	"strings"
	"time"
	"uplift-force-backend/config"
)

type RiotAPI int

const (
	SummonerByName RiotAPI = iota
	MatchHistory
	LeagueEntries
)

type APIConfig struct {
	Endpoint       string
	Method         string
	RequireAuth    bool
	UseRegionalURL bool
}

var apiConfigs = map[RiotAPI]APIConfig{
	SummonerByName: {
		Endpoint:       "/riot/account/v1/accounts/by-riot-id/{gameName}/{tagLine}",
		Method:         "GET",
		RequireAuth:    true,
		UseRegionalURL: true,
	},
	MatchHistory: {
		Endpoint:       "/lol/match/v5/matches/by-puuid/{puuid}/ids",
		Method:         "GET",
		RequireAuth:    true,
		UseRegionalURL: true,
	},
	LeagueEntries: {
		Endpoint:       "/lol/league/v4/entries/by-puuid/{encryptedPUUID}",
		Method:         "GET",
		RequireAuth:    true,
		UseRegionalURL: false,
	},
}

type RiotClient struct {
	httpClient  *http.Client
	apiKey      string
	region      string
	regionGroup config.RegionGroup
	platformURL string // 平台特定的URL
	regionalURL string // 地区分组的URL
}

func NewRiotClient(region string) (*RiotClient, error) {
	// 验证地区是否有效
	//if !config.IsValidRegion(region) {
	//	return nil, fmt.Errorf("不支持的地区: %s，支持的地区: %v", region, config.GetAllRegions())
	//}

	// 获取地区分组
	group, err := config.GetRegionGroup(region)
	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv("RIOT_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("RIOT_API_KEY 环境变量未设置")
	}

	return &RiotClient{
		httpClient:  &http.Client{Timeout: 10 * time.Second},
		apiKey:      apiKey,
		region:      region,
		regionGroup: group,
		platformURL: config.GetRegionBaseURL(region),
		regionalURL: config.GetGroupBaseURL(group),
	}, nil
}

// 根据API类型选择合适的基础URL
func (c *RiotClient) getBaseURL(api RiotAPI) string {
	config := apiConfigs[api]

	if config.UseRegionalURL {
		return c.regionalURL
	}
	return c.platformURL
}

func (c *RiotClient) buildURL(endpoint string, params map[string]string, api RiotAPI) string {
	// 选择基础URL
	baseURL := c.getBaseURL(api)

	// 替换路径参数
	url := endpoint
	for key, value := range params {
		placeholder := "{" + key + "}"
		if strings.Contains(url, placeholder) {
			url = strings.ReplaceAll(url, placeholder, value)
			delete(params, key)
		}
	}

	// 添加查询参数
	if len(params) > 0 {
		values := neturl.Values{}
		for key, value := range params {
			values.Add(key, value)
		}
		url += "?" + values.Encode()
	}

	return baseURL + url
}

func (c *RiotClient) Call(api RiotAPI, params map[string]string) ([]byte, error) {
	config := apiConfigs[api]

	// 构建URL
	url := c.buildURL(config.Endpoint, params, api)

	// 创建请求
	req, err := http.NewRequest(config.Method, url, nil)
	if err != nil {
		return nil, err
	}

	// 添加认证头
	if config.RequireAuth {
		req.Header.Set("X-Riot-Token", c.apiKey)
	}

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 处理响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func GetRegionByTagLine(tagLine string) string {
	// 转为大写并去除首尾空格
	normalizedTagLine := strings.ToUpper(strings.TrimSpace(tagLine))

	if region, exists := config.TagLineRegionMap[normalizedTagLine]; exists {
		return region
	}

	// 未找到匹配的TagLine，返回默认地区
	return config.DefaultRegion
}

// 获取客户端信息
func (c *RiotClient) GetRegion() string {
	return c.region
}

func (c *RiotClient) GetRegionGroup() config.RegionGroup {
	return c.regionGroup
}
