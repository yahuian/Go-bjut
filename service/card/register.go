package card

import (
	"net/http"

	"github.com/YahuiAn/Go-bjut/database"
	"github.com/YahuiAn/Go-bjut/model"
	"github.com/gin-contrib/sessions"

	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/gin-gonic/gin"
)

type cardInfo struct {
	RealName  string `binding:"max=20"`
	Sex       string // TODO 校验取值只能取male，female
	College   string `binding:"max=20"`
	StuNumber string `binding:"max=20"`
	Location  string `binding:"required,max=50"`
}

// 用户捡到一卡通后，登记相关信息，并入库
func Register(c *gin.Context) {
	var info cardInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		logger.Error.Println("json信息错误", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "json信息错误"}) // TODO 具体化错误信息
		return
	}

	if info.RealName == "" && info.StuNumber == "" {
		logger.Error.Println("请至少填写姓名或学号")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请至少填写姓名或学号"})
		return
	}

	// 查找登记者是谁
	session := sessions.Default(c)
	uid := session.Get("user_id")
	student := model.Student{}
	if err := database.DB.First(&student, uid).Error; err != nil {
		logger.Error.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}

	card := &model.Card{
		Registrant: student.NickName,
		Name:       info.RealName,
		Sex:        info.Sex,
		College:    info.College,
		StuNumber:  info.StuNumber,
		Location:   info.Location,
	}

	if err := database.DB.Create(&card).Error; err != nil {
		logger.Error.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "登记失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "登记成功"})
}
