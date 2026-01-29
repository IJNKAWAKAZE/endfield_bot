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

func State(r *gin.Engine) {
	r.GET("/state", func(c *gin.Context) {
		r.LoadHTMLFiles("./template/State.tmpl")
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
		c.HTML(http.StatusOK, "State.tmpl", playerData)
	})
}
