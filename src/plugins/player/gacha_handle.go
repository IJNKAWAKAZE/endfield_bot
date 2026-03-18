package player

import (
	bot "endfield_bot/config"
	"endfield_bot/plugins/account"
	"endfield_bot/plugins/commandoperation"
	"endfield_bot/plugins/messagecleaner"
	"endfield_bot/plugins/skland"
	"endfield_bot/utils"
	"fmt"
	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
	"github.com/spf13/viper"
	"log"
	"sync"
)

type PlayerOperationGacha struct {
	commandoperation.OperationAbstract
}

var (
	gachaTaskMu      sync.Mutex
	gachaTaskRunning = make(map[string]bool)
	gachaTaskQueue   = make(chan struct{}, 1)
)

func (_ PlayerOperationGacha) Run(uid string, userAccount account.UserAccount, chatId int64, message *tgbotapi.Message) error {
	messageId := message.MessageID

	if !tryStartGachaTask(uid) {
		sendMessage := tgbotapi.NewMessage(chatId, "该角色的抽卡记录正在同步中，请稍后再试。")
		sendMessage.ReplyToMessageID = messageId
		bot.Endfield.Send(sendMessage)
		return nil
	}

	sendMessage := tgbotapi.NewMessage(chatId, "抽卡分析已加入队列，正在后台同步数据，完成后会自动发送分析图。")
	sendMessage.ReplyToMessageID = messageId
	send, _ := bot.Endfield.Send(sendMessage)
	messagecleaner.AddDelQueue(chatId, send.MessageID, bot.MsgDelDelay)

	go func() {
		gachaTaskQueue <- struct{}{}
		defer func() {
			<-gachaTaskQueue
			finishGachaTask(uid)
		}()

		if err := runGachaTask(uid, userAccount, chatId, messageId); err != nil {
			log.Printf("gacha task failed uid=%s: %v", uid, err)
			failMessage := tgbotapi.NewMessage(chatId, "抽卡分析生成失败："+err.Error())
			failMessage.ReplyToMessageID = messageId
			bot.Endfield.Send(failMessage)
		}
	}()

	return nil
}

func runGachaTask(uid string, userAccount account.UserAccount, chatId int64, messageId int) error {
	sendAction := tgbotapi.NewChatAction(chatId, "typing")
	bot.Endfield.Send(sendAction)

	u8Token, err := skland.GetU8Token(userAccount.HypergryphToken, uid, userAccount.ServerName)
	if err != nil {
		return fmt.Errorf("获取抽卡 token 失败: %w", err)
	}

	serverId := "1"
	if userAccount.ServerName == "国际服" {
		var player account.UserPlayer
		bot.DBEngine.Table("user_player").Where("uid = ?", uid).First(&player)
		serverId = "2"
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
		log.Printf("同步抽卡记录失败: %v", err)
	} else {
		log.Printf("成功同步 %d 条新抽卡记录到数据库", newCount)
	}

	sendAction = tgbotapi.NewChatAction(chatId, "upload_photo")
	bot.Endfield.Send(sendAction)

	port := viper.GetString("http.port")
	analysisURL := fmt.Sprintf("http://localhost:%s/gacha_analysis?uid=%s", port, uid)

	pic, err := utils.Screenshot(analysisURL, 0, 1.5)
	if err != nil {
		return fmt.Errorf("生成分析截图失败: %w", err)
	}

	sendDocument := tgbotapi.NewDocument(chatId, tgbotapi.FileBytes{Name: "gacha_analysis.png", Bytes: pic})
	sendDocument.ReplyToMessageID = messageId
	bot.Endfield.Send(sendDocument)

	return nil
}

func tryStartGachaTask(uid string) bool {
	gachaTaskMu.Lock()
	defer gachaTaskMu.Unlock()

	if gachaTaskRunning[uid] {
		return false
	}

	gachaTaskRunning[uid] = true
	return true
}

func finishGachaTask(uid string) {
	gachaTaskMu.Lock()
	defer gachaTaskMu.Unlock()
	delete(gachaTaskRunning, uid)
}
