package student

import (
	"net/http"
	"time"

	"github.com/YahuiAn/Go-bjut/model"

	"github.com/YahuiAn/Go-bjut/database"
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
	StuNumber string
	RealName  string
}

// 学生主页
func Home(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("user_id")

	var student HomeInfo

	if err := database.DB.First(&model.Student{}, uid).Scan(&student).Error; err != nil {
		logger.Error.Println("数据库查询失败", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": student})
	return
}
