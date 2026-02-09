package player

import (
	bot "endfield_bot/config"
	"endfield_bot/plugins/account"
	"endfield_bot/plugins/commandoperation"
	"endfield_bot/utils"
	"fmt"
	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
	"github.com/spf13/viper"
)

type PlayerOperationExplore struct {
	commandoperation.OperationAbstract
}

func (_ PlayerOperationExplore) Run(uid string, userAccount account.UserAccount, chatId int64, message *tgbotapi.Message) error {
	sendAction := tgbotapi.NewChatAction(chatId, "upload_photo")
	bot.Endfield.Send(sendAction)

	port := viper.GetString("http.port")
	pic, err := utils.Screenshot(fmt.Sprintf("http://localhost:%s/explore?userId=%d&uid=%s&sklandId=%s", port, userAccount.UserNumber, uid, userAccount.SklandId), 0, 1.5)

	if err != nil {
		msg := tgbotapi.NewMessage(chatId, "生成失败: "+err.Error())
		msg.ReplyToMessageID = message.MessageID
		bot.Endfield.Send(msg)
		return nil
	}

	sendPhoto := tgbotapi.NewPhoto(chatId, tgbotapi.FileBytes{Bytes: pic})
	sendPhoto.ReplyToMessageID = message.MessageID
	bot.Endfield.Send(sendPhoto)
	return nil
}
