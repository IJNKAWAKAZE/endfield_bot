package web

import (
	"endfield_bot/plugins/account"
	"endfield_bot/plugins/skland"
	"endfield_bot/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type MappedCollection struct {
	Name          string `json:"name"`
	LevelID       string `json:"levelId"`
	PuzzleCount   int    `json:"puzzleCount"`
	TrchestCount  int    `json:"trchestCount"`
	PieceCount    int    `json:"pieceCount"`
	BlackboxCount int    `json:"blackboxCount"`
}

type MappedDomain struct {
	Name        string             `json:"name"`
	Level       int                `json:"level"`
	Settlements interface{}        `json:"settlements"`
	Collections []MappedCollection `json:"collections"`
}

func Explore(r *gin.Engine) {
	r.GET("/explore", func(c *gin.Context) {
		r.LoadHTMLFiles("./template/Explore.tmpl")
		var userAccount account.UserAccount
		var skAccount skland.Account
		var player account.UserPlayer
		userId, _ := strconv.ParseInt(c.Query("userId"), 10, 64)
		uid := c.Query("uid")
		sklandId := c.Query("sklandId")

		utils.GetAccountByUserIdAndSklandId(userId, sklandId).Scan(&userAccount)
		utils.GetPlayerByUserId(userId, uid).Scan(&player)
		skAccount.Hypergryph.Token = userAccount.HypergryphToken
		skAccount.Skland.Token = userAccount.SklandToken
		skAccount.Skland.Cred = userAccount.SklandCred

		playerData, err := skland.GetPlayerData(player.RoleId, userAccount.ServerName, player.ServerName, skAccount)
		if err != nil {
			log.Println(err)
			return
		}

		// 调用新的接口获取地图名称树
		mapNameMap := make(map[string]string)
		mapTree, err := skland.GetMapTree(skAccount)
		if err == nil && mapTree != nil {
			for _, category := range mapTree.Maps {
				// 如果 Level 自己的 Name 为空，可以使用 Category 的 Name
				for _, level := range category.Levels {
					name := level.Name
					if name == "" {
						name = category.Name
					}
					mapNameMap[level.ID] = name
				}
			}
		} else {
			log.Println("获取地图树失败:", err)
		}

		// 处理地图名称映射
		domainData := playerData.Data.Detail.Domain
		mappedDomain := make([]MappedDomain, len(domainData))
		for i, d := range domainData {
			collections := make([]MappedCollection, len(d.Collections))
			for j, col := range d.Collections {
				name, ok := mapNameMap[col.LevelID]
				if !ok {
					name = col.LevelID // 没匹配到就用原名
				}
				collections[j] = MappedCollection{
					Name:          name,
					LevelID:       col.LevelID,
					PuzzleCount:   col.PuzzleCount,
					TrchestCount:  col.TrchestCount,
					PieceCount:    col.PieceCount,
					BlackboxCount: col.BlackboxCount,
				}
			}
			mappedDomain[i] = MappedDomain{
				Name:        d.Name,
				Level:       d.Level,
				Settlements: d.Settlements,
				Collections: collections,
			}
		}

		c.HTML(http.StatusOK, "Explore.tmpl", gin.H{
			"Player": playerData,
			"Domain": mappedDomain,
		})
	})
}
