package router

import (
	"github.com/YahuiAn/Go_bjut/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewRouter() *gin.Engine {
	// 设置gin的工作模式
	gin.SetMode(viper.GetString("gin.mode"))

	r := gin.Default()

	// 路由
	v1 := r.Group("/api/v1")
	{
		// 健康检查
		v1.GET("ping", service.Ping)

	}

	return r
}
