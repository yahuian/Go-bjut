package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 健康检查，看服务是否正常
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
