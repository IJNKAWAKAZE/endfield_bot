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

func BoxDetail(r *gin.Engine) {
	r.GET("/box_detail", func(c *gin.Context) {
		r.LoadHTMLFiles("./template/BoxDetail.tmpl")
		var userAccount account.UserAccount
		var skAccount skland.Account
		var player account.UserPlayer
		userId, _ := strconv.ParseInt(c.Query("userId"), 10, 64)
		uid := c.Query("uid")
		sklandId := c.Query("sklandId")
		charId := c.Query("charId")

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

		// 找到选中的角色
		var selectedChar interface{}
		for _, char := range playerData.Data.Detail.Chars {
			if char.CharData.ID == charId {
				selectedChar = char
				break
			}
		}

		if selectedChar == nil {
			log.Println("Character not found:", charId)
			c.String(http.StatusNotFound, "Character not found")
			return
		}

		c.HTML(http.StatusOK, "BoxDetail.tmpl", gin.H{
			"Player": playerData,
			"Char":   selectedChar,
		})
	})
}
