package card

import (
	"net/http"

	"github.com/YahuiAn/Go-bjut/service/user"

	"github.com/YahuiAn/Go-bjut/database"
	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/YahuiAn/Go-bjut/model"
	"github.com/gin-gonic/gin"
)

type cardInfo struct {
	RealName string `binding:"max=20"`
	Sex      string `binding:"omitempty,oneof=male female secrecy"`
	College  string `binding:"max=20"`
	Number   string `binding:"required,max=20"`
	Location string `binding:"required,max=50"`
}

// 用户捡到一卡通后，登记相关信息，并入库
func Register(c *gin.Context) {
	var info cardInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		logger.Error.Println("json信息错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "json信息错误"}) // TODO 具体化错误信息
		return
	}

	who := user.CurrentUser(c)
	if who == (model.User{}) {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "查询登录用户失败"})
		return
	}

	count := 0
	err := database.DB.Model(&model.Card{}).
		Where("number = ? and status != ?", info.Number, model.SuccessfulNotification).Count(&count).Error
	if err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "该卡片已经登记，且还未发送通知，请不要重复登记"})
		return
	}

	card := &model.Card{
		Registrant: who.NickName,
		RealName:   info.RealName,
		Sex:        info.Sex,
		College:    info.College,
		Number:     info.Number,
		Location:   info.Location,
		Status:     model.WaitingNotification,
	}

	if err := database.DB.Create(&card).Error; err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "登记失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "登记成功"})
}
