package student

import (
	"net/http"
	"time"

	"github.com/YahuiAn/Go_bjut/database"
	"github.com/YahuiAn/Go_bjut/logger"
	"github.com/YahuiAn/Go_bjut/model"
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

	student := model.Student{}

	if err := database.DB.First(&student, uid).Error; err != nil {
		logger.Error.Println("数据库查询失败", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}

	homeInfo := HomeInfo{
		ID:        student.ID,
		CreatedAt: student.CreatedAt,
		NickName:  student.NickName,
		Email:     student.Email,
		Telephone: student.Telephone,
		College:   student.College,
		Major:     student.Major,
		ClassName: student.ClassName,
		StuNumber: student.StuNumber,
		RealName:  student.RealName,
	}

	c.JSON(http.StatusOK, gin.H{"msg": homeInfo})
	return
}
