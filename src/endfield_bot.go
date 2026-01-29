package main

import (
	"endfield_bot/config"
	"endfield_bot/core/bot"
	"endfield_bot/core/cron"
	"endfield_bot/core/web"
	"log"
)

func main() {
	//初始化数据库连接
	err := config.DB()
	if err != nil {
		panic(err)
	}
	//初始化redis连接
	config.Redis()
	//初始化机器人
	err = config.Bot()
	if err != nil {
		panic(err)
	}
	//开启定时任务
	err = cron.StartCron()
	if err != nil {
		panic(err)
	}
	//开启http服务
	go web.Start()
	//开始消息监听
	bot.Serve()
}

type signHeaders struct {
	Platform  string `json:"platform"`
	Timestamp string `json:"timestamp"`
	DId       string `json:"dId"`
	VName     string `json:"vName"`
}

// 设置日志格式
func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}
