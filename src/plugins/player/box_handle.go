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

type PlayerOperationBox struct {
	commandoperation.OperationAbstract
}

func (_ PlayerOperationBox) Run(uid string, userAccount account.UserAccount, chatId int64, message *tgbotapi.Message) error {
	messageId := message.MessageID
	sendAction := tgbotapi.NewChatAction(chatId, "upload_photo")
	bot.Endfield.Send(sendAction)

	port := viper.GetString("http.port")
	pic, err := utils.Screenshot(fmt.Sprintf("http://localhost:%s/box?userId=%d&uid=%s&sklandId=%s", port, userAccount.UserNumber, uid, userAccount.SklandId), 0, 1.5)
	if err != nil {
		sendMessage := tgbotapi.NewMessage(chatId, err.Error())
		sendMessage.ReplyToMessageID = messageId
		bot.Endfield.Send(sendMessage)
		return nil
	}
	sendPhoto := tgbotapi.NewPhoto(chatId, tgbotapi.FileBytes{Bytes: pic})
	sendPhoto.ReplyToMessageID = messageId
	bot.Endfield.Send(sendPhoto)
	return nil
}
