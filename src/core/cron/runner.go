package cron

import (
	"endfield_bot/plugins/datasource"
	"endfield_bot/plugins/endfieldnews"
	"endfield_bot/plugins/lottery"
	"endfield_bot/plugins/messagecleaner"
	"endfield_bot/plugins/player"
	"endfield_bot/plugins/sign"
	"github.com/robfig/cron/v3"
	"log"
)

func StartCron() error {
	crontab := cron.New(cron.WithSeconds())

	//终末地bilibili动态 0/30 * * * * ?
	_, err := crontab.AddFunc("0/30 * * * * ?", endfieldnews.BilibiliNews)
	if err != nil {
		return err
	}

	//每周五凌晨2点33更新数据源 0 33 02 ? * FRI
	_, err = crontab.AddFunc("0 33 02 ? * FRI", datasource.UpdateDataSource)
	if err != nil {
		return err
	}

	//每日1点执行自动签到 0 0 1 * * ?
	_, err = crontab.AddFunc("0 0 1 * * ?", sign.AutoSign)
	if err != nil {
		return err
	}

	//清理消息 0/1 * * * * ?
	_, err = crontab.AddFunc("0/1 * * * * ?", messagecleaner.DelMsg)
	if err != nil {
		return err
	}

	//每分钟检查抽奖是否停止报名 0 0/1 * * * ?
	_, err = crontab.AddFunc("0 0/1 * * * ?", lottery.CheckStopLottery)
	if err != nil {
		return err
	}

	//每10分钟检查一次玩家理智 0 0/10 * * * ?
	_, err = crontab.AddFunc("0 0/10 * * * ?", player.CheckSanity)
	if err != nil {
		return err
	}

	//启动定时任务
	crontab.Start()
	log.Println("定时任务已启动")
	return nil
}
