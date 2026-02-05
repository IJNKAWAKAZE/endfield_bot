package player

import (
	bot "endfield_bot/config"
	"endfield_bot/plugins/account"
	"endfield_bot/plugins/commandoperation"
	"endfield_bot/plugins/skland"
	"endfield_bot/utils"
	"fmt"
	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
	"github.com/spf13/viper"
	"log"
)

type PlayerOperationGacha struct {
	commandoperation.OperationAbstract
}

func (_ PlayerOperationGacha) Run(uid string, userAccount account.UserAccount, chatId int64, message *tgbotapi.Message) error {
	messageId := message.MessageID

	u8Token, err := skland.GetU8Token(userAccount.HypergryphToken, uid, userAccount.ServerName)
	if err != nil {
		return fmt.Errorf("获取抽卡Token失败: %w", err)
	}

	serverId := "1"
	if userAccount.ServerName == "国际服" {
		var player account.UserPlayer
		bot.DBEngine.Table("user_player").Where("uid = ?", uid).First(&player)
		serverId = "2" // Asia
		if player.ServerName != "Asia" {
			serverId = "3"
		}
	}

	var allNewRecords []skland.GachaRecord

	poolTypes := []string{
		"E_CharacterGachaPoolType_Special",
		"E_CharacterGachaPoolType_Standard",
		"E_CharacterGachaPoolType_Beginner",
	}

	for _, pType := range poolTypes {
		gachaResp, err := skland.GetGachaRecords(u8Token, serverId, userAccount.ServerName, "char", pType)
		if err != nil {
			log.Printf("获取角色池 %s 记录失败: %v", pType, err)
			continue
		}
		allNewRecords = append(allNewRecords, gachaResp.List...)
	}

	weaponPools, err := skland.GetWeaponPools(u8Token, serverId, userAccount.ServerName)
	if err == nil {
		for _, wp := range weaponPools {
			gachaResp, err := skland.GetGachaRecords(u8Token, serverId, userAccount.ServerName, "weapon", wp.PoolId)
			if err != nil {
				log.Printf("获取武器池 %s 记录失败: %v", wp.PoolName, err)
				continue
			}
			allNewRecords = append(allNewRecords, gachaResp.List...)
		}
	}

	newCount, err := skland.SyncGachaRecordsToDB(uid, allNewRecords)
	if err != nil {
		log.Printf("同步数据库失败: %v", err)
	} else {
		log.Printf("成功同步 %d 条新记录到数据库", newCount)
	}

	sendAction := tgbotapi.NewChatAction(chatId, "upload_photo")
	bot.Endfield.Send(sendAction)

	port := viper.GetString("http.port")
	analysisUrl := fmt.Sprintf("http://localhost:%s/gacha_analysis?uid=%s", port, uid)

	pic, err := utils.Screenshot(analysisUrl, 0, 1.5)
	if err != nil {
		sendMessage := tgbotapi.NewMessage(chatId, "生成分析报告截图失败: "+err.Error())
		sendMessage.ReplyToMessageID = messageId
		bot.Endfield.Send(sendMessage)
		return nil
	}

	sendDocument := tgbotapi.NewDocument(chatId, tgbotapi.FileBytes{Name: "gacha_analysis.png", Bytes: pic})
	sendDocument.ReplyToMessageID = messageId
	bot.Endfield.Send(sendDocument)

	return nil
}
