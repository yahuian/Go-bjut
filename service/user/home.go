package user

import (
	"net/http"
	"time"

	"github.com/YahuiAn/Go-bjut/model"

	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

// 给用户显示的信息
type HomeInfo struct {
	ID        uint
	CreatedAt time.Time
	NickName  string
	Email     string
	Telephone string
	College   string
	Major     string
	ClassName string
	Number    string
	RealName  string
}

// 用户主页
func Home(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("user_id")

	var user HomeInfo

	if err := model.DB.First(&model.User{}, uid).Scan(&user).Error; err != nil {
		logger.Error.Println("数据库查询失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": user})
	return
}
