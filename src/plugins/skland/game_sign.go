package skland

import (
	"fmt"
	"github.com/starudream/go-lib/resty/v2"
	"strconv"
	"strings"
)

type SignGameData struct {
	Ts              string                 `json:"ts"`
	AwardIds        SignGameAwards         `json:"awardIds"`
	ResourceInfoMap map[string]SignGameRes `json:"resourceInfoMap"`
}

type SignGameAward struct {
	Type int    `json:"type"`
	Id   string `json:"id"`
}

type SignGameRes struct {
	Id    string `json:"id"`
	Count int    `json:"count"`
	Name  string `json:"name"`
}

type SignGameAwards []*SignGameAward

func SignGamePlayer(roleId string, account Account, serverName, playerServerName string) (award string, hasSigned bool, err error) {
	account, err = RefreshToken(account, serverName)
	if err != nil {
		return
	}
	var signGameData *SignGameData
	if serverName == "国服" {
		signGameData, err = signGame(roleId, account.Skland)
	} else if serverName == "国际服" {
		signGameData, err = iSignGame(roleId, account.Skland, playerServerName)
	}
	if err != nil {
		e, ok1 := resty.AsRespErr(err)
		if ok1 {
			t, ok2 := e.Response.Error().(*SKBaseResp[interface{}])
			if ok2 && t.Message == "请勿重复签到！" {
				err = nil
				hasSigned = true
			}
		} else {
			err = fmt.Errorf("sign game error: %w", err)
			return
		}
	} else {
		award = shortString(signGameData)
	}
	return
}

// 签到
func signGame(roleId string, skland AccountSkland) (*SignGameData, error) {
	req := SKR().SetHeader("sk-game-role", "3_"+roleId+"_1")
	return SklandRequest[*SignGameData](req, "POST", "/web/v1/game/endfield/attendance", skland)
}

// 签到
func iSignGame(roleId string, skland AccountSkland, serverName string) (*SignGameData, error) {
	serverId := "2"
	if serverName != "Asia" {
		serverId = "3"
	}
	req := SKR().SetHeader("sk-language", "zh_Hans").SetHeader("sk-game-role", "3_"+roleId+"_"+serverId)
	return SkportRequest[*SignGameData](req, "POST", "/web/v1/game/endfield/attendance", skland)
}

func shortString(signGameData *SignGameData) string {
	v := make([]string, len(signGameData.AwardIds))
	for i, a := range signGameData.AwardIds {
		v[i] = signGameData.ResourceInfoMap[a.Id].Name + "*" + strconv.Itoa(signGameData.ResourceInfoMap[a.Id].Count)
	}
	return strings.Join(v, ", ")
}
