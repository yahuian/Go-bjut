package main

import (
	"fmt"

	"github.com/YahuiAn/Go_bjut/config"
	"github.com/YahuiAn/Go_bjut/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func main() {
	// 初始化配置文件
	config.Init()

	// 加载路由
	r := router.NewRouter()

	// 启动项目
	addr := viper.GetString("gin.address")
	port := viper.GetString("gin.port")
	if err := r.Run(addr + ":" + port); err != nil {
		panic(fmt.Sprintf("gin启动失败：%s", err))
	}
}
