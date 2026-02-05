package skland

import (
	bot "endfield_bot/config"
	"endfield_bot/utils"
	"fmt"
	"log"
	"time"
)

// UserGacha 用户抽卡记录表
type UserGacha struct {
	Id         string    `json:"id" gorm:"primaryKey"`             // 主键
	Uid        string    `json:"uid" gorm:"index"`                 // 用户 UID，用于查询
	PoolId     string    `json:"poolId" gorm:"index"`              // 卡池 ID
	PoolName   string    `json:"poolName"`                         // 卡池名称
	PoolType   string    `json:"poolType"`                         // 卡池类型：char, weapon
	ItemId     string    `json:"itemId"`                           // 物品 ID (charId 或 itemId)
	ItemName   string    `json:"itemName"`                         // 物品名称
	ItemType   string    `json:"itemType"`                         // 物品类型：CHARACTER, WEAPON
	Rarity     int       `json:"rarity"`                           // 稀有度：3, 4, 5
	IsNew      bool      `json:"isNew"`                            // 是否首次获得
	GachaTs    int64     `json:"gachaTs"`                          // 抽卡时间戳（毫秒）
	SeqId      string    `json:"seqId" gorm:"uniqueIndex"`         // 序列 ID，唯一标识
	CreateTime time.Time `json:"createTime" gorm:"autoCreateTime"` // 记录创建时间
	UpdateTime time.Time `json:"updateTime" gorm:"autoUpdateTime"` // 记录更新时间
}

func (UserGacha) TableName() string {
	return "user_gacha"
}

// SyncGachaRecordsToDB 将抽卡记录同步到数据库（增量更新）
func SyncGachaRecordsToDB(uid string, records []GachaRecord) (int, error) {
	if len(records) == 0 {
		return 0, nil
	}

	// 1. 将记录按 PoolId 分组
	recordsByPool := make(map[string][]GachaRecord)
	for _, r := range records {
		recordsByPool[r.PoolId] = append(recordsByPool[r.PoolId], r)
	}

	var newRecords []UserGacha

	// 2. 对每个卡池分别处理增量逻辑
	for poolId, poolRecords := range recordsByPool {
		// 获取该用户在该卡池已有的最新记录的 SeqId
		var latestRecord UserGacha
		// 注意：这里必须同时限定 uid 和 pool_id
		bot.DBEngine.Where("uid = ? AND pool_id = ?", uid, poolId).Order("seq_id DESC").First(&latestRecord)
		latestSeqId := latestRecord.SeqId

		// 筛选该池子的新记录
		for _, r := range poolRecords {
			// 如果是新记录（SeqId 更大或该池子尚无记录）
			// 字符串比较 seqId，假设 seqId 是字典序递增的字符串，或者可以转 int 比较。通常 seqId 是字符串ID。
			if latestSeqId == "" || r.SeqId > latestSeqId {
				newRecords = append(newRecords, UserGacha{
					Id:       utils.NewId(), // 使用 utils.NewId() 生成 ID
					Uid:      uid,
					PoolId:   r.PoolId,
					PoolName: r.PoolName,
					PoolType: r.PoolType,
					ItemId:   r.ItemId,
					ItemName: r.ItemName,
					ItemType: r.ItemType,
					Rarity:   r.Rarity,
					IsNew:    r.IsNew,
					GachaTs:  r.Ts,
					SeqId:    r.SeqId,
				})
			}
		}
	}

	if len(newRecords) == 0 {
		log.Printf("UID %s 没有新的抽卡记录需要同步", uid)
		return 0, nil
	}

	// 批量插入新记录
	result := bot.DBEngine.Create(&newRecords)
	if result.Error != nil {
		return 0, fmt.Errorf("保存抽卡记录失败: %w", result.Error)
	}

	log.Printf("UID %s 成功同步 %d 条新抽卡记录", uid, len(newRecords))
	return len(newRecords), nil
}

// GetGachaRecordsFromDB 从数据库获取用户的所有抽卡记录
func GetGachaRecordsFromDB(uid string) ([]GachaRecord, error) {
	var dbRecords []UserGacha
	result := bot.DBEngine.Where("uid = ?", uid).Order("seq_id DESC").Find(&dbRecords)
	if result.Error != nil {
		return nil, fmt.Errorf("查询抽卡记录失败: %w", result.Error)
	}

	// 转换为 GachaRecord 格式
	var records []GachaRecord
	for _, r := range dbRecords {
		records = append(records, GachaRecord{
			SeqId:    r.SeqId,
			PoolId:   r.PoolId,
			PoolName: r.PoolName,
			ItemId:   r.ItemId,
			ItemName: r.ItemName,
			ItemType: r.ItemType,
			Rarity:   r.Rarity,
			IsNew:    r.IsNew,
			Ts:       r.GachaTs,
			PoolType: r.PoolType,
		})
	}

	return records, nil
}
