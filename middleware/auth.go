package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "请先登陆"})
		c.Abort()
	}
}
