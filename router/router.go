package router

import (
	"github.com/YahuiAn/Go_bjut/middleware"
	"github.com/YahuiAn/Go_bjut/service"
	"github.com/YahuiAn/Go_bjut/service/student"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewRouter() *gin.Engine {
	// 设置gin的工作模式
	gin.SetMode(viper.GetString("gin.mode"))

	r := gin.Default()

	// 中间件，顺序不能变
	r.Use(middleware.Session(viper.GetString("session.secret")))

	// 路由
	v1 := r.Group("/api/v1")
	{
		// 健康检查
		v1.GET("ping", service.Ping)

		v1.POST("/stu/register", student.Register)
		v1.POST("/stu/login", student.Login)

		// 需要登陆保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("stu/me", student.Home)
			auth.DELETE("stu/logout", student.Logout)
		}

	}

	return r
}
