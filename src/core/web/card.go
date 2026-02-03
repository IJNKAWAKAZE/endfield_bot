package web

import (
	"endfield_bot/plugins/account"
	"endfield_bot/plugins/skland"
	"endfield_bot/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Card(r *gin.Engine) {
	r.GET("/card", func(c *gin.Context) {
		r.LoadHTMLFiles("./template/Card.tmpl")
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

		// 处理服务器显示
		playerData.Data.Detail.Base.ServerName = player.ServerName

		// 处理苏醒日时间戳转换
		ts, _ := strconv.ParseInt(playerData.Data.Detail.Base.CreateTime, 10, 64)
		if ts > 0 {
			playerData.Data.Detail.Base.CreateTime = time.Unix(ts, 0).Format("2006-01-02")
		}

		c.HTML(http.StatusOK, "Card.tmpl", playerData)
	})
}
