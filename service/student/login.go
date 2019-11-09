package student

import (
	"net/http"

	"github.com/gin-contrib/sessions"

	"golang.org/x/crypto/bcrypt"

	"github.com/YahuiAn/Go_bjut/database"
	"github.com/YahuiAn/Go_bjut/model"

	"github.com/YahuiAn/Go_bjut/logger"
	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	Nickname string `json:"nickname" binding:"required,min=2,max=30"`
	Password string `json:"password" binding:"required,min=8,max=40"`
}

// 学生登录
func Login(c *gin.Context) {
	var info LoginInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		logger.Error.Println("json信息错误", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "json信息错误"}) // TODO 具体化错误信息
		return
	}

	student := model.Student{}
	if err := database.DB.Where("nick_name = ?", info.Nickname).First(&student).Error; err != nil {
		logger.Error.Println("用户名错误", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名错误"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(info.Password)); err != nil {
		logger.Error.Println("密码错误", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "密码错误"})
		return
	}

	// 设置session
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", student.ID)
	if err := s.Save(); err != nil {
		logger.Error.Println("session设置错误", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "session设置错误"})
		return
	}

	logger.Info.Println("登录成功", student.NickName)
	c.JSON(http.StatusOK, gin.H{"msg": "登录成功", "data": student.NickName})
}
