package skland

import (
	"fmt"
	"github.com/tidwall/gjson"
	"net/url"
)

type GachaRecord struct {
	SeqId    string `json:"seqId"`
	PoolId   string `json:"poolId"`
	PoolName string `json:"poolName"`
	ItemId   string `json:"itemId"`   // 对应 API 中的 charId 或 itemId
	ItemName string `json:"itemName"` // 对应 API 中的 charName 或 itemName
	ItemType string `json:"itemType"` // CHARACTER, WEAPON
	Rarity   int    `json:"rarity"`   // 3, 4, 5
	IsNew    bool   `json:"isNew"`    // 是否首次获得
	Ts       int64  `json:"ts"`       // 对应 API 中的 gachaTs 或 ts
	PoolType string `json:"poolType"` // char, weapon
}

type GachaListResp struct {
	List    []GachaRecord `json:"list"`
	HasMore bool          `json:"hasMore"`
}

// GetGachaRecords 获取抽卡记录 (自动获取所有页)
func GetGachaRecords(u8Token, serverId, serverName, poolType string, poolId string) (*GachaListResp, error) {
	provider := "hypergryph"
	if serverName == "国际服" {
		provider = "gryphline"
	}

	endpoint := fmt.Sprintf("https://ef-webview.%s.com/api/record/%s", provider, poolType)
	var allRecords []GachaRecord
	lastSeqId := ""
	hasMore := true

	for hasMore {
		fullUrl := fmt.Sprintf("%s?token=%s&server_id=%s&lang=zh-cn", endpoint, url.QueryEscape(u8Token), serverId)

		if poolType == "char" && poolId != "" {
			fullUrl += fmt.Sprintf("&pool_type=%s", poolId)
		} else if poolType == "weapon" && poolId != "" {
			fullUrl += fmt.Sprintf("&pool_id=%s", poolId)
		}

		if lastSeqId != "" {
			fullUrl += fmt.Sprintf("&seq_id=%s", lastSeqId)
		}

		req := HR()
		res, err := req.Execute("GET", fullUrl)
		if err != nil {
			return nil, fmt.Errorf("fetch gacha records error: %w", err)
		}

		body := string(res.Body())
		code := gjson.Get(body, "code").Int()
		if code != 0 {
			msg := gjson.Get(body, "msg").String()
			return nil, fmt.Errorf("gacha api error: %s (code: %d)", msg, code)
		}

		data := gjson.Get(body, "data")
		list := data.Get("list").Array()
		if len(list) == 0 {
			break
		}

		for _, item := range list {
			record := GachaRecord{
				SeqId:    item.Get("seqId").String(),
				PoolId:   item.Get("poolId").String(),
				PoolName: item.Get("poolName").String(),
				Rarity:   int(item.Get("rarity").Int()),
				IsNew:    item.Get("isNew").Bool(),
				PoolType: poolType,
			}

			if poolType == "char" {
				// 适配角色池特有字段
				record.ItemId = item.Get("charId").String()
				record.ItemName = item.Get("charName").String()
				record.Ts = item.Get("gachaTs").Int()
				record.ItemType = "CHARACTER"
			} else {
				// 适配武器池字段
				// 根据用户提供的代码逻辑，武器池接口直接返回列表，字段应为 weaponId/weaponName
				record.ItemId = item.Get("weaponId").String()
				// 防御性 fallback，以防万一
				if record.ItemId == "" {
					record.ItemId = item.Get("itemId").String()
				}

				record.ItemName = item.Get("weaponName").String()
				if record.ItemName == "" {
					record.ItemName = item.Get("itemName").String()
				}

				record.Ts = item.Get("gachaTs").Int()
				if record.Ts == 0 {
					record.Ts = item.Get("ts").Int()
				}
				record.ItemType = "WEAPON"
			}
			allRecords = append(allRecords, record)
		}

		hasMore = data.Get("hasMore").Bool()
		if hasMore {
			lastSeqId = allRecords[len(allRecords)-1].SeqId
		}
	}

	return &GachaListResp{
		List:    allRecords,
		HasMore: false,
	}, nil
}

type WeaponPool struct {
	PoolId   string `json:"poolId"`
	PoolName string `json:"poolName"`
}

// GetWeaponPools 获取武器池列表
func GetWeaponPools(u8Token, serverId, serverName string) ([]WeaponPool, error) {
	provider := "hypergryph"
	if serverName == "国际服" {
		provider = "gryphline"
	}

	endpoint := fmt.Sprintf("https://ef-webview.%s.com/api/record/weapon/pool", provider)
	fullUrl := fmt.Sprintf("%s?token=%s&server_id=%s&lang=zh-cn", endpoint, url.QueryEscape(u8Token), serverId)

	req := HR()
	res, err := req.Execute("GET", fullUrl)
	if err != nil {
		return nil, fmt.Errorf("fetch weapon pools error: %w", err)
	}

	body := string(res.Body())
	code := gjson.Get(body, "code").Int()
	if code != 0 {
		msg := gjson.Get(body, "msg").String()
		return nil, fmt.Errorf("weapon pool api error: %s (code: %d)", msg, code)
	}

	var pools []WeaponPool
	data := gjson.Get(body, "data").Array()
	for _, item := range data {
		pools = append(pools, WeaponPool{
			PoolId:   item.Get("poolId").String(),
			PoolName: item.Get("poolName").String(),
		})
	}

	return pools, nil
}
