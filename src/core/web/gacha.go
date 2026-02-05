package web

import (
	bot "endfield_bot/config"
	"endfield_bot/plugins/account"
	"endfield_bot/plugins/skland"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"time"
)

type GachaAnalysisData struct {
	UID               string
	PlayerName        string
	ServerName        string
	RoleId            string
	SyncTime          string
	CharStats         StatsSummary
	WeaponStats       StatsSummary
	RecentCharPulls   []SixStarDetail
	RecentWeaponPulls []SixStarDetail
	Pools             []PoolAnalysis
}

type StatsSummary struct {
	TotalCount   int
	SixStarCount int
	AveragePity  float64
}

type PoolAnalysis struct {
	Name           string
	Type           string // char, weapon
	Total          int
	CurrentPity    int
	SixStarRate    float64
	RecentSixStars []SixStarDetail
	LastTimeMs     int64 // 用于排序活跃度
}

type SixStarDetail struct {
	PoolName  string
	ItemName  string
	TimeStr   string
	Rarity    int
	Pity      int
	GlobalIdx int
}

func Gacha(r *gin.Engine) {
	r.GET("/gacha_analysis", func(c *gin.Context) {
		r.LoadHTMLFiles("./template/GachaAnalysis.tmpl")
		uid := c.Query("uid")

		records, err := skland.GetGachaRecordsFromDB(uid)
		if err != nil {
			c.String(http.StatusInternalServerError, "Database error: "+err.Error())
			return
		}

		if len(records) == 0 {
			c.String(http.StatusNotFound, "No data for UID: "+uid)
			return
		}

		// 获取玩家详细信息
		var player account.UserPlayer
		result := bot.DBEngine.Table("user_player").Where("uid = ?", uid).First(&player)
		if result.Error != nil {
			player.Uid = uid
			player.PlayerName = "Unknown"
			player.ServerName = "Unknown"
			player.RoleId = uid
		}

		analysis := analyzeGacha(uid, records, player)
		c.HTML(http.StatusOK, "GachaAnalysis.tmpl", analysis)
	})
}

func analyzeGacha(uid string, allRecords []skland.GachaRecord, player account.UserPlayer) GachaAnalysisData {
	data := GachaAnalysisData{
		UID:        uid,
		PlayerName: player.PlayerName,
		ServerName: player.ServerName,
		RoleId:     player.RoleId,
		SyncTime:   time.Now().Format("2006-01-02 15:04:05"),
		Pools:      []PoolAnalysis{},
	}

	// 1. 分离记录
	var charRecords, weaponRecords []skland.GachaRecord
	for _, r := range allRecords {
		if r.PoolType == "char" {
			charRecords = append(charRecords, r)
		} else {
			weaponRecords = append(weaponRecords, r)
		}
	}

	// 2. 计算统计数据
	data.CharStats = calculateStats(charRecords)
	data.WeaponStats = calculateStats(weaponRecords)

	// 收集所有六星记录用于展示（带PoolName）
	var allCharSixStars []SixStarDetail
	var allWeaponSixStars []SixStarDetail

	// 3. 按池子分析详情
	poolMap := make(map[string][]skland.GachaRecord)
	poolInfo := make(map[string]skland.GachaRecord)

	for _, r := range allRecords {
		poolMap[r.PoolId] = append(poolMap[r.PoolId], r)
		if _, ok := poolInfo[r.PoolId]; !ok {
			poolInfo[r.PoolId] = r
		}
	}

	for pId, pRecords := range poolMap {
		// 按 SeqId 排序（时间正序）
		sort.Slice(pRecords, func(i, j int) bool {
			return pRecords[i].SeqId < pRecords[j].SeqId
		})

		// 获取该池子最新的记录时间
		lastTs := int64(0)
		if len(pRecords) > 0 {
			lastTs = pRecords[len(pRecords)-1].Ts
		}

		info := poolInfo[pId]
		pAnalysis := PoolAnalysis{
			Name:       info.PoolName,
			Type:       info.PoolType,
			Total:      len(pRecords),
			LastTimeMs: lastTs,
		}

		pity := 0
		var sixStarDetails []SixStarDetail
		for i, r := range pRecords {
			pity++
			if r.Rarity == 6 {
				detail := SixStarDetail{
					PoolName:  info.PoolName,
					ItemName:  r.ItemName,
					TimeStr:   time.Unix(r.Ts/1000, 0).Format("2006-01-02 15:04"),
					Rarity:    6,
					Pity:      pity,
					GlobalIdx: i + 1,
				}
				sixStarDetails = append(sixStarDetails, detail)
				pity = 0 // 重置保底
			}
		}
		pAnalysis.CurrentPity = pity
		if pAnalysis.Total > 0 {
			pAnalysis.SixStarRate = float64(len(sixStarDetails)) / float64(pAnalysis.Total) * 100
		}

		// 这里只需要最近5条给池子详情（如果还展示的话）
		if len(sixStarDetails) > 5 {
			pAnalysis.RecentSixStars = sixStarDetails[len(sixStarDetails)-5:]
		} else {
			pAnalysis.RecentSixStars = sixStarDetails
		}
		// 反转
		for i, j := 0, len(pAnalysis.RecentSixStars)-1; i < j; i, j = i+1, j-1 {
			pAnalysis.RecentSixStars[i], pAnalysis.RecentSixStars[j] = pAnalysis.RecentSixStars[j], pAnalysis.RecentSixStars[i]
		}

		data.Pools = append(data.Pools, pAnalysis)

		// 收集到全局列表
		if info.PoolType == "char" {
			allCharSixStars = append(allCharSixStars, sixStarDetails...)
		} else {
			allWeaponSixStars = append(allWeaponSixStars, sixStarDetails...)
		}
	}

	// 4. 处理全局近期记录
	// 倒序
	sortSixStarsDesc(allCharSixStars)
	sortSixStarsDesc(allWeaponSixStars)

	limit := 20
	if len(allCharSixStars) > limit {
		data.RecentCharPulls = allCharSixStars[:limit]
	} else {
		data.RecentCharPulls = allCharSixStars
	}

	if len(allWeaponSixStars) > limit {
		data.RecentWeaponPulls = allWeaponSixStars[:limit]
	} else {
		data.RecentWeaponPulls = allWeaponSixStars
	}

	// 对 Pools 进行筛选：只保留最近活跃的 6 个
	sort.Slice(data.Pools, func(i, j int) bool {
		// 先按最后活跃时间降序排，找出最近的 6 个
		return data.Pools[i].LastTimeMs > data.Pools[j].LastTimeMs
	})

	maxPools := 6
	if len(data.Pools) > maxPools {
		data.Pools = data.Pools[:maxPools]
	}

	// 对这 6 个再按类型排序，保证展示美观（角色在前）
	sort.Slice(data.Pools, func(i, j int) bool {
		if data.Pools[i].Type != data.Pools[j].Type {
			return data.Pools[i].Type == "char"
		}
		// 同类型按时间倒序（最新的在前）
		return data.Pools[i].LastTimeMs > data.Pools[j].LastTimeMs
	})

	return data
}

func sortSixStarsDesc(list []SixStarDetail) {
	sort.Slice(list, func(i, j int) bool {
		// 时间字符串降序比较
		return list[i].TimeStr > list[j].TimeStr
	})
}

// 辅助函数：计算一组记录的统计摘要
func calculateStats(records []skland.GachaRecord) StatsSummary {
	// 确保按时间/SeqId 排序
	sort.Slice(records, func(i, j int) bool {
		return records[i].SeqId < records[j].SeqId
	})

	stats := StatsSummary{
		TotalCount: len(records),
	}

	groupedByPool := make(map[string][]skland.GachaRecord)
	for _, r := range records {
		groupedByPool[r.PoolId] = append(groupedByPool[r.PoolId], r)
	}

	totalPitySum := 0

	for _, poolRecords := range groupedByPool {
		// 确保池内有序
		sort.Slice(poolRecords, func(i, j int) bool {
			return poolRecords[i].SeqId < poolRecords[j].SeqId
		})

		pity := 0
		for _, r := range poolRecords {
			pity++
			if r.Rarity == 6 {
				stats.SixStarCount++
				totalPitySum += pity
				pity = 0
			}
		}
	}

	if stats.SixStarCount > 0 {
		stats.AveragePity = float64(totalPitySum) / float64(stats.SixStarCount)
	}

	return stats
}
