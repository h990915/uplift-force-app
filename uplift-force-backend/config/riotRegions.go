package config

import "fmt"

type RegionGroup string

const (
	Asia    RegionGroup = "asia"
	Europe  RegionGroup = "europe"
	America RegionGroup = "americas"
)

// 地区到分组的映射
var regionGroupMap = map[string]RegionGroup{
	// 亚洲地区
	"kr":  Asia,
	"kr1": Asia,
	"jp1": Asia,
	"ph2": Asia,
	"sg2": Asia,
	"th2": Asia,
	"tw2": Asia,
	"vn2": Asia,

	// 欧洲地区
	"euw1": Europe,
	"eun1": Europe,
	"tr1":  Europe,
	"ru":   Europe,
	"ru1":  Europe,

	// 美洲地区
	"na1": America,
	"br1": America,
	"la1": America,
	"la2": America,
	"oc1": America,
}

// 分组到基础URL的映射
var groupBaseURLMap = map[RegionGroup]string{
	Asia:    "https://asia.api.riotgames.com",
	Europe:  "https://europe.api.riotgames.com",
	America: "https://americas.api.riotgames.com",
}

// TagLine到地区的映射表
var TagLineRegionMap = map[string]string{
	// 韩国
	"KR":  "kr",
	"KR1": "kr",

	// 北美
	"NA":  "na1",
	"NA1": "na1",

	// 欧洲西部
	"EUW":  "euw1",
	"EUW1": "euw1",
	"WEST": "euw1",

	// 欧洲东部
	"EUNE": "eun1",
	"EUN":  "eun1",
	"NE":   "eun1",

	// 日本
	"JP":  "jp1",
	"JP1": "jp1",

	// 巴西
	"BR":  "br1",
	"BR1": "br1",

	// 拉丁美洲北部
	"LAN": "la1",
	"LA1": "la1",

	// 拉丁美洲南部
	"LAS": "la2",
	"LA2": "la2",

	// 土耳其
	"TR":  "tr1",
	"TR1": "tr1",

	// 俄罗斯
	"RU":  "ru",
	"RU1": "ru",
	"RUS": "ru",

	// 大洋洲
	"OCE": "oc1",
	"OC1": "oc1",
	"AUS": "oc1",

	// 菲律宾
	"PH":  "ph2",
	"PH2": "ph2",

	// 新加坡/马来西亚
	"SG":  "sg2",
	"SG2": "sg2",
	"MY":  "sg2",

	// 泰国
	"TH":  "th2",
	"TH2": "th2",

	// 台湾
	"TW":  "tw2",
	"TW2": "tw2",

	// 越南
	"VN":  "vn2",
	"VN2": "vn2",
}

// 默认地区
const DefaultRegion = "na1"

// 获取地区对应的分组
func GetRegionGroup(region string) (RegionGroup, error) {
	group, exists := regionGroupMap[region]
	if !exists {
		return Europe, nil
	}
	return group, nil
}

// 获取分组的基础URL
func GetGroupBaseURL(group RegionGroup) string {
	return groupBaseURLMap[group]
}

// 获取地区的平台URL（用于平台特定的API）
func GetRegionBaseURL(region string) string {
	return fmt.Sprintf("https://%s.api.riotgames.com", region)
}

// 检查地区是否有效
func IsValidRegion(region string) bool {
	_, exists := regionGroupMap[region]
	return exists
}

// 获取所有支持的地区
func GetAllRegions() []string {
	regions := make([]string, 0, len(regionGroupMap))
	for region := range regionGroupMap {
		regions = append(regions, region)
	}
	return regions
}

// 按分组获取地区列表
func GetRegionsByGroup(group RegionGroup) []string {
	var regions []string
	for region, g := range regionGroupMap {
		if g == group {
			regions = append(regions, region)
		}
	}
	return regions
}
