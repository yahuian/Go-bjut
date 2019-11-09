package student

import (
	"net/http"

	"github.com/YahuiAn/Go_bjut/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	if err := s.Save(); err != nil {
		logger.Error.Println("session设置错误", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "session设置错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "退出成功"})
	return
}
