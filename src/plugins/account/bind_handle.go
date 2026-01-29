package account

import (
	"encoding/json"
	bot "endfield_bot/config"
	"endfield_bot/plugins/skland"
	"endfield_bot/utils"
	"fmt"
	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// BindHandle 绑定角色
func BindHandle(update tgbotapi.Update) error {
	chatId := update.Message.Chat.ID
	var buttons [][]tgbotapi.InlineKeyboardButton
	buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("国服", fmt.Sprintf("%s,%s,%s", "chooseServer", "国服", "setToken")),
		tgbotapi.NewInlineKeyboardButtonData("国际服", fmt.Sprintf("%s,%s,%s", "chooseServer", "国际服", "setToken")),
	))
	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(
		buttons...,
	)
	sendMessage := tgbotapi.NewMessage(chatId, "请选择要绑定的服务器")
	sendMessage.ReplyMarkup = inlineKeyboardMarkup
	bot.Endfield.Send(sendMessage)
	return nil
}

// SetToken 设置token
func SetToken(update tgbotapi.Update) error {
	message := update.Message
	chatId := message.Chat.ID
	userId := message.From.ID
	token := message.Text

	sendAction := tgbotapi.NewChatAction(chatId, "typing")
	bot.Endfield.Send(sendAction)

	var userToken UserToken
	err := json.Unmarshal([]byte(token), &userToken)
	if err == nil {
		token = userToken.Data.Content
	}
	account, err := skland.Login(token, serverNameMap[chatId])
	if err != nil {
		sendMessage := tgbotapi.NewMessage(chatId, "登录失败！请检查token是否正确。")
		bot.Endfield.Send(sendMessage)
		return err
	}
	// 查询账户是否存在
	var userAccount UserAccount
	res := utils.GetAccountByUserIdAndSklandId(userId, account.UserId).Scan(&userAccount)
	if res.RowsAffected > 0 {
		// 更新账户信息
		userAccount.HypergryphToken = token
		userAccount.SklandToken = account.Skland.Token
		userAccount.SklandCred = account.Skland.Cred
		bot.DBEngine.Table("user_account").Save(&userAccount)
	} else {
		// 不存在 新增账户
		id, _ := gonanoid.New(32)
		userAccount = UserAccount{
			Id:              id,
			UserName:        message.From.FullName(),
			UserNumber:      userId,
			HypergryphToken: token,
			SklandToken:     account.Skland.Token,
			SklandCred:      account.Skland.Cred,
			SklandId:        account.UserId,
			ServerName:      serverNameMap[chatId],
		}
		bot.DBEngine.Table("user_account").Create(&userAccount)
	}
	delete(tgbotapi.WaitMessage, chatId)
	// 获取角色列表
	players, err := skland.EndfieldPlayers(account.Skland, userAccount.ServerName)
	if err != nil || len(players) == 0 {
		sendMessage := tgbotapi.NewMessage(chatId, "未查询到绑定角色！")
		bot.Endfield.Send(sendMessage)
		return err
	}

	sklandIdMap[chatId] = account.UserId
	var buttons [][]tgbotapi.InlineKeyboardButton
	for _, player := range players {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s(%s)", player.DefaultRole.Nickname, player.DefaultRole.ServerName), fmt.Sprintf("%s,%s,%s,%s,%s", "bind", player.Uid, player.DefaultRole.ServerName, player.DefaultRole.Nickname, player.DefaultRole.RoleId)),
		))
	}
	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(
		buttons...,
	)
	sendMessage := tgbotapi.NewMessage(chatId, "请选择要绑定的角色")
	sendMessage.ReplyMarkup = inlineKeyboardMarkup
	bot.Endfield.Send(sendMessage)
	return nil
}

// CancelHandle 取消操作
func CancelHandle(update tgbotapi.Update) error {
	chatId := update.Message.Chat.ID
	delete(tgbotapi.WaitMessage, chatId)
	sendMessage := tgbotapi.NewMessage(chatId, "已取消操作")
	bot.Endfield.Send(sendMessage)
	return nil
}
