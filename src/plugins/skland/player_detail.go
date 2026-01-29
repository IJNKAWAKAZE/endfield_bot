package skland

import (
	"encoding/json"
	bot "endfield_bot/config"
	"fmt"
	"log"
)

func GetPlayerData(roleId, serverName, playerServerName string, account Account) (*PlayerDetail, error) {
	var playerDetail *PlayerDetail
	account, err := RefreshToken(account, serverName)
	if err != nil {
		log.Println(err.Error())
		return playerDetail, err
	}
	var playerDetailstr string
	if serverName == "国服" {
		playerDetailstr, err = playerData(roleId, account.Skland)
	} else if serverName == "国际服" {
		playerDetailstr, err = iPlayerData(roleId, account.Skland, playerServerName)
	}
	if err != nil {
		return playerDetail, err
	}
	json.Unmarshal([]byte(playerDetailstr), &playerDetail)
	bot.DBEngine.Exec("update user_player set player_name = ? where roleId = ?", playerDetail.Data.Detail.Base.Name, roleId)
	return playerDetail, nil
}

// playerData 角色数据
func playerData(roleId string, skland AccountSkland) (string, error) {
	req := SKR()
	return SklandRequestPlayerData(req, "GET", fmt.Sprintf("/api/v1/game/endfield/card/detail?roleId=%s&serverId=1", roleId), skland)
}

// iPlayerData 角色数据
func iPlayerData(roleId string, skland AccountSkland, serverName string) (string, error) {
	serverId := "2"
	if serverName != "Asia" {
		serverId = "3"
	}
	req := SKR().SetHeader("sk-language", "zh_Hans")
	return SkportRequestPlayerData(req, "GET", fmt.Sprintf("/api/v1/game/endfield/card/detail?roleId=%s&serverId=%s", roleId, serverId), skland)

}
