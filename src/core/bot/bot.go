package bot

import (
	bot "endfield_bot/config"
	"endfield_bot/plugins/account"
	"endfield_bot/plugins/gatekeeper"
	"endfield_bot/plugins/lottery"
	"endfield_bot/plugins/player"
	"endfield_bot/plugins/sign"
	"endfield_bot/plugins/system"
	"github.com/spf13/viper"
	"log"
)

// Serve TG机器人阻塞监听
func Serve() {
	log.Println("机器人启动成功")
	b := bot.Endfield.AddHandle()
	bot.Endfield.Debug = viper.GetBool("bot.debug")
	bot.Endfield.IgnoreChannelCMD = true
	b.JoinRequestProcessor(gatekeeper.JoinRequestHandle)
	b.NewMemberProcessor(gatekeeper.NewMemberHandle)
	b.LeftMemberProcessor(gatekeeper.LeftMemberHandle)

	// callback
	b.NewCallBackProcessor("verify", gatekeeper.CallBackData)
	b.NewCallBackProcessor("request", gatekeeper.RequestCallBackData)
	b.NewCallBackProcessor("chooseServer", account.ChooseServer)
	b.NewCallBackProcessor("bind", account.ChoosePlayer)
	b.NewCallBackProcessor("unbind", account.UnbindPlayer)
	b.NewCallBackProcessor("sign", sign.SignPlayer)
	b.NewCallBackProcessor("player", player.PlayerData)
	b.NewCallBackProcessor("report", system.Report)

	// 私聊
	b.NewPrivateCommandProcessor("start", system.HelpHandle)
	b.NewPrivateCommandProcessor("cancel", account.CancelHandle)
	b.NewPrivateCommandProcessor("bind", account.BindHandle)
	b.NewPrivateCommandProcessor("unbind", account.UnbindHandle)
	b.NewPrivateCommandProcessor("reset_token", account.SetTokenHandle)

	// wait
	b.NewWaitMessageProcessor("setToken", account.SetToken)
	b.NewWaitMessageProcessor("resetToken", account.ResetToken)

	// 普通
	b.NewCommandProcessor("help", system.HelpHandle)
	b.NewCommandProcessor("ping", system.PingHandle)
	b.NewCommandProcessor("sign", sign.SignHandle)
	b.NewCommandProcessor("state", player.PlayerHandle)
	b.NewCommandProcessor("report", system.ReportHandle)
	b.NewCommandProcessor("join_lottery", lottery.JoinLotteryHandle)
	b.NewCommandProcessor("lottery_detail", lottery.LotteryDetailHandle)

	// 权限
	b.NewCommandProcessor("news", system.NewsHandle)
	b.NewCommandProcessor("request_mode", system.RequestModeHandle)
	b.NewCommandProcessor("reg", system.RegulationHandle)
	b.NewCommandProcessor("welcome", system.WelcomeHandle)
	b.NewCommandProcessor("clear", system.ClearHandle)
	b.NewCommandProcessor("start_lottery", lottery.StartLotteryHandle)
	b.NewCommandProcessor("stop_lottery", lottery.StopLotteryHandle)
	b.NewCommandProcessor("end_lottery", lottery.EndLotteryHandle)
	b.NewCommandProcessor("lottery", lottery.LotteryHandle)

	// 仅拥有者
	b.NewCommandProcessor("update", system.UpdateHandle)
	b.NewCommandProcessor("sign_all", sign.SignAllHandle)
	b.NewCommandProcessor("kill", system.KillHandle)
	b.Run()
}
