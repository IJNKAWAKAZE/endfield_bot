package web

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/assets", "./assets")
	Help(r)
	State(r)
	Box(r)
	Lottery(r)
	port := viper.GetString("http.port")
	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
