package user

import (
	"github.com/YahuiAn/Go-bjut/database"
	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/YahuiAn/Go-bjut/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 获取当前登录用户的信息
func CurrentUser(c *gin.Context) model.User {
	session := sessions.Default(c)
	uid := session.Get("user_id")
	var user model.User
	if err := database.DB.First(&user, uid).Error; err != nil {
		logger.Error.Println(err)
	}
	return user
}
