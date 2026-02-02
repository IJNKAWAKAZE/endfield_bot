package web

import (
	"endfield_bot/plugins/account"
	"endfield_bot/plugins/skland"
	"endfield_bot/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func Box(r *gin.Engine) {
	r.GET("/box", func(c *gin.Context) {
		r.LoadHTMLFiles("./template/Box.tmpl")
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
			utils.WebC <- err
			return
		}

		// 稀有度 > 等级 > 职业 > 属性
		sort.SliceStable(playerData.Data.Detail.Chars, func(i, j int) bool {
			charI := playerData.Data.Detail.Chars[i]
			charJ := playerData.Data.Detail.Chars[j]

			// 1. 稀有度排序 (降序)
			ri, _ := strconv.Atoi(charI.CharData.Rarity.Value)
			rj, _ := strconv.Atoi(charJ.CharData.Rarity.Value)
			if ri != rj {
				return ri > rj
			}

			// 2. 等级排序 (降序)
			if charI.Level != charJ.Level {
				return charI.Level > charJ.Level
			}

			// 3. 职业排序 (升序)
			if charI.CharData.Profession.Value != charJ.CharData.Profession.Value {
				return charI.CharData.Profession.Value < charJ.CharData.Profession.Value
			}

			// 4. 属性排序 (升序)
			return charI.CharData.Property.Value < charJ.CharData.Property.Value
		})

		c.HTML(http.StatusOK, "Box.tmpl", playerData)
	})
}
