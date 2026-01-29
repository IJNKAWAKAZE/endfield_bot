package system

import (
	bot "endfield_bot/config"
	"endfield_bot/plugins/messagecleaner"
	"endfield_bot/utils"
	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
)

func WelcomeHandle(update tgbotapi.Update) error {
	chatId := update.Message.Chat.ID
	userId := update.Message.From.ID
	messageId := update.Message.MessageID
	messagecleaner.AddDelQueue(chatId, messageId, 5)

	if bot.Endfield.IsAdmin(chatId, userId) {
		text := update.Message.CommandArguments()
		if text != "" {
			var joined utils.GroupJoined
			utils.GetJoinedByChatId(chatId).Scan(&joined)
			joined.Welcome = text
			bot.DBEngine.Table("group_joined").Save(&joined)
			sendMessage := tgbotapi.NewMessage(chatId, "设置入群欢迎信息成功")
			msg, err := bot.Endfield.Send(sendMessage)
			if err != nil {
				return err
			}
			messagecleaner.AddDelQueue(msg.Chat.ID, msg.MessageID, bot.MsgDelDelay)
		}
		return nil
	}

	sendMessage := tgbotapi.NewMessage(chatId, "无使用权限！")
	sendMessage.ReplyToMessageID = messageId
	msg, err := bot.Endfield.Send(sendMessage)
	if err != nil {
		return err
	}
	messagecleaner.AddDelQueue(msg.Chat.ID, msg.MessageID, bot.MsgDelDelay)
	return nil
}
