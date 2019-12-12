package main

import (
	"fmt"

	"github.com/YahuiAn/Go-bjut/service/card"

	"github.com/YahuiAn/Go-bjut/config"
	"github.com/YahuiAn/Go-bjut/router"
	"github.com/robfig/cron"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func main() {
	// 初始化配置文件
	config.Init()

	// 加载路由
	r := router.NewRouter()

	// 启动定时任务
	smsTime := viper.GetString("cron.sms")
	spec := "@every " + smsTime
	c := cron.New()
	if err := c.AddFunc(spec, card.Notice); err != nil {
		panic(err)
	}
	c.Start()

	// 启动项目
	addr := viper.GetString("gin.address")
	port := viper.GetString("gin.port")
	if err := r.Run(addr + ":" + port); err != nil {
		panic(fmt.Sprintf("gin启动失败：%s", err))
	}
}
