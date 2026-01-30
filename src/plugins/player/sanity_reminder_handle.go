package player

import (
	"crypto/rand"
	bot "endfield_bot/config"
	"endfield_bot/plugins/account"
	"endfield_bot/plugins/skland"
	"endfield_bot/utils"
	"fmt"
	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"log"
	"math/big"
	"strconv"
	"time"
)

type UserSanityReminder struct {
	Id         string    `json:"id" gorm:"primaryKey"`
	UserName   string    `json:"userName"`
	UserNumber int64     `json:"userNumber"`
	IsReminded bool      `json:"isReminded"` // 记录上一次是否已经提醒过，防止回满后重复提醒
	CreateTime time.Time `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"updateTime" gorm:"autoUpdateTime"`
	Remark     string    `json:"remark"`
}

// SanityReminderHandle 理智提醒
func SanityReminderHandle(update tgbotapi.Update) error {
	param := update.Message.CommandArguments()
	chatId := update.Message.Chat.ID
	userId := update.Message.From.ID
	messageId := update.Message.MessageID

	if param == "on" {
		var reminder UserSanityReminder
		res := utils.GetSanityReminderByUserId(userId).Scan(&reminder)
		if res.RowsAffected > 0 {
			sendMessage := tgbotapi.NewMessage(chatId, "您已经开启了理智回满提醒！")
			sendMessage.ReplyToMessageID = messageId
			bot.Endfield.Send(sendMessage)
			return nil
		}
		id, _ := gonanoid.New(32)
		reminder = UserSanityReminder{
			Id:         id,
			UserName:   update.Message.From.FullName(),
			UserNumber: userId,
			IsReminded: false,
		}
		bot.DBEngine.Table("user_sanity_reminder").Create(&reminder)
		sendMessage := tgbotapi.NewMessage(chatId, "开启理智回满提醒成功！每 10 分钟将检查一次理智情况。")
		sendMessage.ReplyToMessageID = messageId
		bot.Endfield.Send(sendMessage)
	} else if param == "off" {
		bot.DBEngine.Exec("delete from user_sanity_reminder where user_number = ?", userId)
		sendMessage := tgbotapi.NewMessage(chatId, "已关闭理智回满提醒！")
		sendMessage.ReplyToMessageID = messageId
		bot.Endfield.Send(sendMessage)
	} else {
		sendMessage := tgbotapi.NewMessage(chatId, "理智提醒指令说明：\n开启：`/sanity on`\n关闭：`/sanity off`")
		sendMessage.ParseMode = tgbotapi.ModeMarkdownV2
		sendMessage.ReplyToMessageID = messageId
		bot.Endfield.Send(sendMessage)
	}
	return nil
}

// CheckSanity 检查所有用户的理智是否回满
func CheckSanity() {
	var reminders []UserSanityReminder
	res := utils.GetSanityReminders().Scan(&reminders)
	if res.RowsAffected > 0 {
		go func() {
			log.Printf("开始检查 %d 位用户的理智回复情况...", len(reminders))
			for _, reminder := range reminders {
				// 获取 0-300 之间的随机秒数，将请求均匀分布在 10 分钟的时间窗口内
				// 避免单位时间内大量请求森空岛 API
				r, _ := rand.Int(rand.Reader, big.NewInt(300))
				random, _ := strconv.Atoi(r.String())
				time.Sleep(time.Second * time.Duration(random))
				checkUserSanity(reminder)
			}
			log.Println("理智检查执行完毕")
		}()
	}
}

func checkUserSanity(reminder UserSanityReminder) {
	var players []account.UserPlayer
	res := utils.GetPlayersByUserId(reminder.UserNumber).Scan(&players)
	if res.RowsAffected > 0 {
		for _, player := range players {
			var skAccount skland.Account
			var userAccount account.UserAccount
			res := utils.GetAccountByUid(reminder.UserNumber, player.Uid).Scan(&userAccount)
			if res.RowsAffected > 0 {
				skAccount.Hypergryph.Token = userAccount.HypergryphToken
				skAccount.Skland.Token = userAccount.SklandToken
				skAccount.Skland.Cred = userAccount.SklandCred

				// 获取玩家实时数据
				playerData, err := skland.GetPlayerData(player.RoleId, userAccount.ServerName, player.ServerName, skAccount)
				if err != nil {
					log.Println("检查理智时获取数据失败:", err)
					continue
				}

				// 获取回满时间戳
				maxTs, _ := strconv.ParseInt(playerData.Data.Detail.Dungeon.MaxTs, 10, 64)
				curStamina, _ := strconv.Atoi(playerData.Data.Detail.Dungeon.CurStamina)
				maxStamina, _ := strconv.Atoi(playerData.Data.Detail.Dungeon.MaxStamina)

				// 如果理智已满且之前未提醒
				if (maxTs <= time.Now().Unix() || curStamina >= maxStamina) && !reminder.IsReminded {
					msg := fmt.Sprintf("⚠️ 理智回满提醒 ⚠️\n\n角色: %s\n当前理智: %d / %d\n您的理智已经回满，请及时上号处理！", player.PlayerName, curStamina, maxStamina)
					sendMessage := tgbotapi.NewMessage(reminder.UserNumber, msg)
					bot.Endfield.Send(sendMessage)

					// 标记为已提醒
					bot.DBEngine.Table("user_sanity_reminder").Where("user_number = ?", reminder.UserNumber).Update("is_reminded", true)
				}

				// 如果理智没有回满，且之前是已提醒状态，则重置提醒状态（用户已经消耗过理智了）
				if curStamina < maxStamina && reminder.IsReminded {
					bot.DBEngine.Table("user_sanity_reminder").Where("user_number = ?", reminder.UserNumber).Update("is_reminded", false)
				}
			}
		}
	}
}
