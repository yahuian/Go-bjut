package router

import (
	"github.com/YahuiAn/Go-bjut/middleware"
	"github.com/YahuiAn/Go-bjut/service"
	"github.com/YahuiAn/Go-bjut/service/card"
	"github.com/YahuiAn/Go-bjut/service/forum/comment"
	"github.com/YahuiAn/Go-bjut/service/forum/dynamic"
	"github.com/YahuiAn/Go-bjut/service/user"
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

		v1.POST("/user/register", user.Register)
		v1.POST("/user/bjut-register", user.BjutRegister)
		v1.POST("/user/login", user.Login)

		v1.GET("card/index", card.Index)
		v1.GET("dynamic/:id", dynamic.GetDynamicById)
		v1.GET("dynamic/", dynamic.Index)

		// 需要登陆保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("user/me", user.Home)
			auth.PUT("user/update", user.Update)
			auth.DELETE("user/logout", user.Logout)

			auth.POST("card/register", card.Register)

			auth.POST("dynamic/create", dynamic.Create)
			auth.POST("comment/create", comment.Create)
		}

	}

	return r
}
