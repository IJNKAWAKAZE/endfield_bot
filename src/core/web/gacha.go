package web

import (
	bot "endfield_bot/config"
	"endfield_bot/plugins/account"
	"endfield_bot/plugins/skland"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	poolCategoryBase    = "base"
	poolCategoryLimited = "limited"
)

type GachaAnalysisData struct {
	UID        string
	PlayerName string
	ServerName string
	RoleId     string
	SyncTime   string
	Sections   []GachaTypeSection
}

type GachaTypeSection struct {
	TypeLabel          string
	TypeKey            string
	Stats              StatsSummary
	LimitedCurrentPity int
	RecentPulls        []SixStarDetail
	BasePools          []PoolAnalysis
	LimitedPools       []PoolAnalysis
}

type StatsSummary struct {
	TotalCount   int
	SixStarCount int
	AveragePity  float64
}

type PoolAnalysis struct {
	Name           string
	Type           string
	Category       string
	CategoryLabel  string
	Total          int
	CurrentPity    int
	SixStarRate    float64
	RecentSixStars []SixStarDetail
	FirstTimeMs    int64
	LastTimeMs     int64
	FirstSeqId     string
	LastSeqId      string
}

type SixStarDetail struct {
	PoolName  string
	ItemName  string
	TimeStr   string
	Rarity    int
	Pity      int
	GlobalIdx int
	Ts        int64
	SeqId     string
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
	}

	data.Sections = []GachaTypeSection{
		buildGachaTypeSection("char", "角色池", filterRecordsByType(allRecords, "char")),
		buildGachaTypeSection("weapon", "武器池", filterRecordsByType(allRecords, "weapon")),
	}

	return data
}

func buildGachaTypeSection(typeKey, typeLabel string, records []skland.GachaRecord) GachaTypeSection {
	basePools, limitedPools, recentPulls := buildPoolAnalyses(records)

	return GachaTypeSection{
		TypeKey:            typeKey,
		TypeLabel:          typeLabel,
		Stats:              calculateStats(records),
		LimitedCurrentPity: calculateLimitedCurrentPity(records),
		RecentPulls:        recentPulls,
		BasePools:          basePools,
		LimitedPools:       limitedPools,
	}
}

func filterRecordsByType(allRecords []skland.GachaRecord, poolType string) []skland.GachaRecord {
	var result []skland.GachaRecord
	for _, record := range allRecords {
		if record.PoolType == poolType {
			result = append(result, record)
		}
	}
	return result
}

func buildPoolAnalyses(records []skland.GachaRecord) ([]PoolAnalysis, []PoolAnalysis, []SixStarDetail) {
	grouped := make(map[string][]skland.GachaRecord)
	infoMap := make(map[string]skland.GachaRecord)

	for _, record := range records {
		grouped[record.PoolId] = append(grouped[record.PoolId], record)
		if _, ok := infoMap[record.PoolId]; !ok {
			infoMap[record.PoolId] = record
		}
	}

	var basePools []PoolAnalysis
	var limitedPools []PoolAnalysis
	var allSixStars []SixStarDetail

	for poolID, poolRecords := range grouped {
		sortRecordsAsc(poolRecords)

		info := infoMap[poolID]
		category := classifyPoolCategory(info)
		analysis := PoolAnalysis{
			Name:          info.PoolName,
			Type:          info.PoolType,
			Category:      category,
			CategoryLabel: categoryLabel(category),
			Total:         len(poolRecords),
		}

		if len(poolRecords) > 0 {
			analysis.FirstTimeMs = poolRecords[0].Ts
			analysis.LastTimeMs = poolRecords[len(poolRecords)-1].Ts
			analysis.FirstSeqId = poolRecords[0].SeqId
			analysis.LastSeqId = poolRecords[len(poolRecords)-1].SeqId
		}

		pity := 0
		var sixStars []SixStarDetail
		for index, record := range poolRecords {
			pity++
			if record.Rarity != 6 {
				continue
			}

			detail := SixStarDetail{
				PoolName:  info.PoolName,
				ItemName:  record.ItemName,
				TimeStr:   formatGachaTime(record.Ts),
				Rarity:    6,
				Pity:      pity,
				GlobalIdx: index + 1,
				Ts:        record.Ts,
				SeqId:     record.SeqId,
			}
			sixStars = append(sixStars, detail)
			allSixStars = append(allSixStars, detail)
			pity = 0
		}

		analysis.CurrentPity = pity
		if analysis.Total > 0 {
			analysis.SixStarRate = float64(len(sixStars)) / float64(analysis.Total) * 100
		}

		if len(sixStars) > 5 {
			analysis.RecentSixStars = reverseSixStars(sixStars[len(sixStars)-5:])
		} else {
			analysis.RecentSixStars = reverseSixStars(sixStars)
		}

		if category == poolCategoryLimited {
			limitedPools = append(limitedPools, analysis)
		} else {
			basePools = append(basePools, analysis)
		}
	}

	sortPoolAnalyses(basePools)
	sortPoolAnalyses(limitedPools)
	sortSixStarsDesc(allSixStars)

	if len(allSixStars) > 20 {
		allSixStars = allSixStars[:20]
	}

	return basePools, limitedPools, allSixStars
}

func sortPoolAnalyses(pools []PoolAnalysis) {
	sort.Slice(pools, func(i, j int) bool {
		if pools[i].LastTimeMs != pools[j].LastTimeMs {
			return pools[i].LastTimeMs > pools[j].LastTimeMs
		}
		seqCmp := compareSeqID(pools[i].LastSeqId, pools[j].LastSeqId)
		if seqCmp != 0 {
			return seqCmp > 0
		}
		if pools[i].FirstTimeMs != pools[j].FirstTimeMs {
			return pools[i].FirstTimeMs > pools[j].FirstTimeMs
		}
		if pools[i].FirstSeqId != pools[j].FirstSeqId {
			return compareSeqID(pools[i].FirstSeqId, pools[j].FirstSeqId) > 0
		}
		if pools[i].Name == pools[j].Name {
			return pools[i].Type < pools[j].Type
		}
		return pools[i].Name < pools[j].Name
	})
}

func reverseSixStars(items []SixStarDetail) []SixStarDetail {
	result := make([]SixStarDetail, len(items))
	copy(result, items)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func sortSixStarsDesc(list []SixStarDetail) {
	sort.Slice(list, func(i, j int) bool {
		if list[i].Ts != list[j].Ts {
			return list[i].Ts > list[j].Ts
		}
		seqCmp := compareSeqID(list[i].SeqId, list[j].SeqId)
		if seqCmp != 0 {
			return seqCmp > 0
		}
		return list[i].ItemName < list[j].ItemName
	})
}

func calculateStats(records []skland.GachaRecord) StatsSummary {
	stats := StatsSummary{
		TotalCount: len(records),
	}

	grouped := make(map[string][]skland.GachaRecord)
	for _, record := range records {
		key := effectivePityKey(record)
		grouped[key] = append(grouped[key], record)
	}

	totalPitySum := 0
	for _, group := range grouped {
		sortRecordsAsc(group)
		pity := 0
		for _, record := range group {
			pity++
			if record.Rarity != 6 {
				continue
			}
			stats.SixStarCount++
			totalPitySum += pity
			pity = 0
		}
	}

	if stats.SixStarCount > 0 {
		stats.AveragePity = float64(totalPitySum) / float64(stats.SixStarCount)
	}

	return stats
}

func calculateLimitedCurrentPity(records []skland.GachaRecord) int {
	var limitedRecords []skland.GachaRecord
	for _, record := range records {
		if classifyPoolCategory(record) == poolCategoryLimited {
			limitedRecords = append(limitedRecords, record)
		}
	}

	sortRecordsAsc(limitedRecords)

	pity := 0
	for _, record := range limitedRecords {
		pity++
		if record.Rarity == 6 {
			pity = 0
		}
	}

	return pity
}

func effectivePityKey(record skland.GachaRecord) string {
	if classifyPoolCategory(record) == poolCategoryLimited {
		return record.PoolType + ":" + poolCategoryLimited
	}
	return record.PoolType + ":" + record.PoolId
}

func classifyPoolCategory(record skland.GachaRecord) string {
	poolID := strings.ToLower(record.PoolId)

	switch record.PoolType {
	case "char":
		if strings.HasPrefix(poolID, "special") {
			return poolCategoryLimited
		}
		return poolCategoryBase
	case "weapon":
		if strings.Contains(poolID, "weponbox") {
			return poolCategoryLimited
		}
		return poolCategoryBase
	default:
		return poolCategoryBase
	}
}

func categoryLabel(category string) string {
	if category == poolCategoryLimited {
		return "限定池"
	}
	return "基础池"
}

func sortRecordsAsc(records []skland.GachaRecord) {
	sort.Slice(records, func(i, j int) bool {
		if records[i].Ts != records[j].Ts {
			return records[i].Ts < records[j].Ts
		}
		seqCmp := compareSeqID(records[i].SeqId, records[j].SeqId)
		if seqCmp != 0 {
			return seqCmp < 0
		}
		return records[i].ItemName < records[j].ItemName
	})
}

func formatGachaTime(ts int64) string {
	return time.Unix(ts/1000, 0).Format("2006-01-02 15:04")
}

func compareSeqID(left, right string) int {
	leftSeq, leftErr := strconv.ParseInt(left, 10, 64)
	rightSeq, rightErr := strconv.ParseInt(right, 10, 64)
	if leftErr == nil && rightErr == nil {
		switch {
		case leftSeq < rightSeq:
			return -1
		case leftSeq > rightSeq:
			return 1
		default:
			return 0
		}
	}

	switch {
	case left < right:
		return -1
	case left > right:
		return 1
	default:
		return 0
	}
}
