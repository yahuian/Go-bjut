package dynamic

import (
	"net/http"

	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/YahuiAn/Go-bjut/pkg"

	"github.com/YahuiAn/Go-bjut/model"
	"github.com/YahuiAn/Go-bjut/service/user"

	"github.com/gin-gonic/gin"
)

type dynamicInfo struct {
	Title   string `binding:"required,max=80"`
	Content string
}

func Create(c *gin.Context) {
	var info dynamicInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": pkg.Warn(err)})
		return
	}

	who := user.CurrentUser(c)
	if who == (model.User{}) {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取当前用户失败"})
		return
	}

	dynamic := model.Dynamic{
		NickName: who.NickName,
		Title:    info.Title,
		Content:  info.Content,
	}

	if err := model.DB.Create(&dynamic).Error; err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库操作失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "发布成功", "date": dynamic.ID})
}
