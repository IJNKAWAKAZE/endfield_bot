package player

import (
	bot "endfield_bot/config"
	"endfield_bot/plugins/account"
	"endfield_bot/plugins/commandoperation"
	"endfield_bot/utils"
	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
	"log"
)

var inited = false

// PlayerHandle 角色信息查询
func PlayerHandle(update tgbotapi.Update) error {
	if !inited {
		initFactory()
		inited = true
	}
	chatId := update.Message.Chat.ID
	messageId := update.Message.MessageID
	var userAccount account.UserAccount
	var players []account.UserPlayer
	var operationP *commandoperation.OperationI
	userAccountP, playersP, err := getAccountAndPlayers(update)
	if err != nil || userAccountP == nil || playersP == nil {
		return err
	}
	command := update.Message.Command()
	if commandoperation.HaveNextStep(chatId) {
		return commandoperation.GetStep(chatId).Run(update)
	}
	if len(command) != 0 {
		operationP = playerOperationFactory(command)
	}
	if operationP == nil {
		log.Printf("Unmatched Handle %s", update.Message.Command())
		return nil
	}
	operation := *operationP
	players = playersP
	userAccount = *userAccountP
	if players == nil || len(players) == 0 {
		log.Printf("Code reach impossible point players = %v after getPlayer warp", players)
	}
	if !operation.CheckRequirementsAndPrepare(update) {
		msg, isMarkDown := operation.HintOnRequirementsFailed()
		sendMessage := tgbotapi.NewMessage(chatId, msg)
		sendMessage.ReplyToMessageID = messageId
		if isMarkDown {
			sendMessage.ParseMode = tgbotapi.ModeMarkdownV2
		}
		bot.Endfield.Send(sendMessage)
		return nil
	}
	if len(players) > 1 {
		return playerSelector(update, userAccount, players, operation, command)
	}
	utils.GetAccountByUid(userAccount.UserNumber, players[0].Uid).Scan(&userAccount)
	return operation.Run(players[0].Uid, userAccount, chatId, update.Message)
}
