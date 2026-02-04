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
	"math"
	"strconv"
	"strings"
)

type PlayerOperationBoxDetail struct {
	commandoperation.OperationAbstract
}

func (_ PlayerOperationBoxDetail) Run(uid string, userAccount account.UserAccount, chatId int64, message *tgbotapi.Message) error {
	var skAccount skland.Account
	skAccount.Hypergryph.Token = userAccount.HypergryphToken
	skAccount.Skland.Token = userAccount.SklandToken
	skAccount.Skland.Cred = userAccount.SklandCred

	var player account.UserPlayer
	utils.GetPlayerByUserId(userAccount.UserNumber, uid).Scan(&player)

	playerData, err := skland.GetPlayerData(player.RoleId, userAccount.ServerName, player.ServerName, skAccount)
	if err != nil {
		return err
	}

	chars := playerData.Data.Detail.Chars
	if len(chars) == 0 {
		msg := tgbotapi.NewMessage(chatId, "该角色下没有任何干员数据。")
		msg.ReplyToMessageID = message.MessageID
		bot.Endfield.Send(msg)
		return nil
	}

	return sendBoxDetailSelector(chatId, message.MessageID, userAccount, uid, chars, 0)
}

func sendBoxDetailSelector(chatId int64, replyToId int, userAccount account.UserAccount, uid string, chars []struct {
	CharData struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		AvatarSqURL string `json:"avatarSqUrl"`
		AvatarRtURL string `json:"avatarRtUrl"`
		Rarity      struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"rarity"`
		Profession struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"profession"`
		Property struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"property"`
		WeaponType struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"weaponType"`
		Skills []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Type struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"type"`
			Property struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"property"`
			IconURL    string `json:"iconUrl"`
			Desc       string `json:"desc"`
			DescParams struct {
			} `json:"descParams"`
			DescLevelParams struct {
				Num1 struct {
					Level  string `json:"level"`
					Params struct {
						Atb             string `json:"atb"`
						AtkScale        string `json:"atk_scale"`
						DisplayAtkScale string `json:"display_atk_scale"`
						Poise           string `json:"poise"`
					} `json:"params"`
				} `json:"1"`
			} `json:"descLevelParams"`
		} `json:"skills"`
		LabelType       string   `json:"labelType"`
		IllustrationURL string   `json:"illustrationUrl"`
		Tags            []string `json:"tags"`
	} `json:"charData"`
	ID         string `json:"id"`
	Level      int    `json:"level"`
	UserSkills map[string]struct {
		SkillID  string `json:"skillId"`
		Level    int    `json:"level"`
		MaxLevel int    `json:"maxLevel"`
	} `json:"userSkills"`
	BodyEquip struct {
		EquipID   string `json:"equipId"`
		EquipData struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			IconURL string `json:"iconUrl"`
			Rarity  struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"rarity"`
			Type struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"type"`
			Level struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"level"`
			Properties  []string `json:"properties"`
			IsAccessory bool     `json:"isAccessory"`
			Suit        struct {
				ID              string      `json:"id"`
				Name            string      `json:"name"`
				SkillID         string      `json:"skillId"`
				SkillDesc       string      `json:"skillDesc"`
				SkillDescParams interface{} `json:"skillDescParams"`
			} `json:"suit"`
			Function string `json:"function"`
			Pkg      string `json:"pkg"`
		} `json:"equipData"`
	} `json:"bodyEquip,omitempty"`
	ArmEquip struct {
		EquipID   string `json:"equipId"`
		EquipData struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			IconURL string `json:"iconUrl"`
			Rarity  struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"rarity"`
			Type struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"type"`
			Level struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"level"`
			Properties  []string `json:"properties"`
			IsAccessory bool     `json:"isAccessory"`
			Suit        struct {
				ID              string      `json:"id"`
				Name            string      `json:"name"`
				SkillID         string      `json:"skillId"`
				SkillDesc       string      `json:"skillDesc"`
				SkillDescParams interface{} `json:"skillDescParams"`
			} `json:"suit"`
			Function string `json:"function"`
			Pkg      string `json:"pkg"`
		} `json:"equipData"`
	} `json:"armEquip,omitempty"`
	FirstAccessory struct {
		EquipID   string `json:"equipId"`
		EquipData struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			IconURL string `json:"iconUrl"`
			Rarity  struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"rarity"`
			Type struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"type"`
			Level struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"level"`
			Properties  []string `json:"properties"`
			IsAccessory bool     `json:"isAccessory"`
			Suit        struct {
				ID              string      `json:"id"`
				Name            string      `json:"name"`
				SkillID         string      `json:"skillId"`
				SkillDesc       string      `json:"skillDesc"`
				SkillDescParams interface{} `json:"skillDescParams"`
			} `json:"suit"`
			Function string `json:"function"`
			Pkg      string `json:"pkg"`
		} `json:"equipData"`
	} `json:"firstAccessory,omitempty"`
	SecondAccessory struct {
		EquipID   string `json:"equipId"`
		EquipData struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			IconURL string `json:"iconUrl"`
			Rarity  struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"rarity"`
			Type struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"type"`
			Level struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"level"`
			Properties  []string `json:"properties"`
			IsAccessory bool     `json:"isAccessory"`
			Suit        struct {
				ID              string      `json:"id"`
				Name            string      `json:"name"`
				SkillID         string      `json:"skillId"`
				SkillDesc       string      `json:"skillDesc"`
				SkillDescParams interface{} `json:"skillDescParams"`
			} `json:"suit"`
			Function string `json:"function"`
			Pkg      string `json:"pkg"`
		} `json:"equipData"`
	} `json:"secondAccessory,omitempty"`
	TacticalItem struct {
		TacticalItemId   string `json:"tacticalItemId"`
		TacticalItemData struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			IconURL string `json:"iconUrl"`
			Rarity  struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"rarity"`
			ActiveEffectType struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"activeEffectType"`
			ActiveEffect        string      `json:"activeEffect"`
			PassiveEffect       string      `json:"passiveEffect"`
			ActiveEffectParams  interface{} `json:"activeEffectParams"`
			PassiveEffectParams interface{} `json:"passiveEffectParams"`
		} `json:"tacticalItemData"`
	} `json:"tacticalItem,omitempty"`
	EvolvePhase    int `json:"evolvePhase"`
	PotentialLevel int `json:"potentialLevel"`
	Weapon         struct {
		WeaponData struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			IconURL string `json:"iconUrl"`
			Rarity  struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"rarity"`
			Type struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"type"`
			Function    string `json:"function"`
			Description string `json:"description"`
			Skills      []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"skills"`
		} `json:"weaponData"`
		Level             int `json:"level"`
		RefineLevel       int `json:"refineLevel"`
		BreakthroughLevel int `json:"breakthroughLevel"`
		Gem               struct {
			ID   string `json:"id"`
			Icon string `json:"icon"`
		} `json:"gem"`
	} `json:"weapon,omitempty"`
	Gender string `json:"gender"`
	OwnTs  string `json:"ownTs"`
}, page int) error {
	pageSize := 8
	totalChars := len(chars)
	totalPages := int(math.Ceil(float64(totalChars) / float64(pageSize)))

	if page < 0 {
		page = 0
	}
	if page >= totalPages {
		page = totalPages - 1
	}

	start := page * pageSize
	end := start + pageSize
	if end > totalChars {
		end = totalChars
	}

	currentCharSubset := chars[start:end]

	var buttons [][]tgbotapi.InlineKeyboardButton
	// 两列展示角色
	for i := 0; i < len(currentCharSubset); i += 2 {
		var row []tgbotapi.InlineKeyboardButton
		char1 := currentCharSubset[i]
		row = append(row, tgbotapi.NewInlineKeyboardButtonData(
			fmt.Sprintf("%s (Lv.%d)", char1.CharData.Name, char1.Level),
			fmt.Sprintf("box_detail_char,%s,%s", uid, char1.CharData.ID),
		))

		if i+1 < len(currentCharSubset) {
			char2 := currentCharSubset[i+1]
			row = append(row, tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s (Lv.%d)", char2.CharData.Name, char2.Level),
				fmt.Sprintf("box_detail_char,%s,%s", uid, char2.CharData.ID),
			))
		}
		buttons = append(buttons, row)
	}

	// 翻页按钮
	var navRow []tgbotapi.InlineKeyboardButton
	if page > 0 {
		navRow = append(navRow, tgbotapi.NewInlineKeyboardButtonData("⬅️ 上一页", fmt.Sprintf("box_detail_page,%s,%d", uid, page-1)))
	}
	navRow = append(navRow, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d/%d", page+1, totalPages), "none"))
	if page < totalPages-1 {
		navRow = append(navRow, tgbotapi.NewInlineKeyboardButtonData("下一页 ➡️", fmt.Sprintf("box_detail_page,%s,%d", uid, page+1)))
	}
	buttons = append(buttons, navRow)

	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(buttons...)
	msg := tgbotapi.NewMessage(chatId, "请选择要查看详情的干员：")
	msg.ReplyMarkup = inlineKeyboardMarkup
	if replyToId != 0 {
		msg.ReplyToMessageID = replyToId
	}

	bot.Endfield.Send(msg)

	return nil
}

func BoxDetailCallback(update tgbotapi.Update) error {
	callbackQuery := update.CallbackQuery
	data := update.CallbackData()
	d := strings.Split(data, ",")

	if len(d) < 3 {
		return nil
	}

	action := d[0]
	uid := d[1]

	// 验证是否是本人操作（简单通过userId验证）
	// 注意：这里的userId验证可能需要从数据库查询uid所属的userNumber
	var userAccount account.UserAccount
	userId := callbackQuery.From.ID
	res := utils.GetAccountByUid(userId, uid).Scan(&userAccount)
	if res.RowsAffected == 0 {
		callbackQuery.Answer(true, "这不是你的角色！")
		return nil
	}

	if action == "box_detail_page" {
		page, _ := strconv.Atoi(d[2])

		// 重新获取数据以翻页
		var skAccount skland.Account
		skAccount.Hypergryph.Token = userAccount.HypergryphToken
		skAccount.Skland.Token = userAccount.SklandToken
		skAccount.Skland.Cred = userAccount.SklandCred

		var player account.UserPlayer
		utils.GetPlayerByUserId(userAccount.UserNumber, uid).Scan(&player)
		playerData, err := skland.GetPlayerData(player.RoleId, userAccount.ServerName, player.ServerName, skAccount)
		if err != nil {
			return err
		}

		// 编辑原消息实现翻页效果
		pageSize := 8
		totalChars := len(playerData.Data.Detail.Chars)
		totalPages := int(math.Ceil(float64(totalChars) / float64(pageSize)))

		start := page * pageSize
		end := start + pageSize
		if end > totalChars {
			end = totalChars
		}
		currentCharSubset := playerData.Data.Detail.Chars[start:end]

		var buttons [][]tgbotapi.InlineKeyboardButton
		for i := 0; i < len(currentCharSubset); i += 2 {
			var row []tgbotapi.InlineKeyboardButton
			char1 := currentCharSubset[i]
			row = append(row, tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%s (Lv.%d)", char1.CharData.Name, char1.Level),
				fmt.Sprintf("box_detail_char,%s,%s", uid, char1.CharData.ID),
			))
			if i+1 < len(currentCharSubset) {
				char2 := currentCharSubset[i+1]
				row = append(row, tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%s (Lv.%d)", char2.CharData.Name, char2.Level),
					fmt.Sprintf("box_detail_char,%s,%s", uid, char2.CharData.ID),
				))
			}
			buttons = append(buttons, row)
		}

		var navRow []tgbotapi.InlineKeyboardButton
		if page > 0 {
			navRow = append(navRow, tgbotapi.NewInlineKeyboardButtonData("⬅️ 上一页", fmt.Sprintf("box_detail_page,%s,%d", uid, page-1)))
		}
		navRow = append(navRow, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d/%d", page+1, totalPages), "none"))
		if page < totalPages-1 {
			navRow = append(navRow, tgbotapi.NewInlineKeyboardButtonData("下一页 ➡️", fmt.Sprintf("box_detail_page,%s,%d", uid, page+1)))
		}
		buttons = append(buttons, navRow)

		editMsg := tgbotapi.NewEditMessageReplyMarkup(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID, tgbotapi.NewInlineKeyboardMarkup(buttons...))
		bot.Endfield.Send(editMsg)
		callbackQuery.Answer(false, "")
		return nil
	}

	if action == "box_detail_char" {
		charId := d[2]
		callbackQuery.Answer(false, "正在生成干员详情，请稍候...")

		// 删除选择角色的消息
		delMsg := tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID)
		bot.Endfield.Send(delMsg)

		port := viper.GetString("http.port")
		pic, err := utils.Screenshot(fmt.Sprintf("http://localhost:%s/box_detail?userId=%d&uid=%s&sklandId=%s&charId=%s", port, userAccount.UserNumber, uid, userAccount.SklandId, charId), 0, 1.5)
		if err != nil {
			msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "生成失败: "+err.Error())
			bot.Endfield.Send(msg)
			return nil
		}

		sendPhoto := tgbotapi.NewPhoto(callbackQuery.Message.Chat.ID, tgbotapi.FileBytes{Bytes: pic})
		if callbackQuery.Message.ReplyToMessage != nil {
			sendPhoto.ReplyToMessageID = callbackQuery.Message.ReplyToMessage.MessageID
		}
		bot.Endfield.Send(sendPhoto)
		return nil
	}

	return nil
}
