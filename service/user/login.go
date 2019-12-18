package user

import (
	"net/http"

	"github.com/gin-contrib/sessions"

	"golang.org/x/crypto/bcrypt"

	"github.com/YahuiAn/Go-bjut/model"
	"github.com/YahuiAn/Go-bjut/tip"

	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	Nickname string `binding:"required,min=2,max=30"`
	Password string `binding:"required,min=8,max=40"`
}

// 用户登录
func Login(c *gin.Context) {
	var info LoginInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": tip.Warn(err)})
		return
	}

	user := model.User{}
	if err := model.DB.Where("nick_name = ?", info.Nickname).First(&user).Error; err != nil {
		logger.Error.Println("用户名错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名错误"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(info.Password)); err != nil {
		logger.Error.Println("密码错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "密码错误"})
		return
	}

	// 设置session
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	if err := s.Save(); err != nil {
		logger.Error.Println("session设置错误", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "session设置错误"})
		return
	}

	logger.Info.Println("登录成功", user.NickName)
	c.JSON(http.StatusOK, gin.H{"msg": "登录成功", "data": user.NickName})
}
