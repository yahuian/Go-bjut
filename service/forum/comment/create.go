package comment

import (
	"net/http"

	"github.com/YahuiAn/Go-bjut/service/user"
	"github.com/YahuiAn/Go-bjut/tip"

	"github.com/YahuiAn/Go-bjut/logger"

	"github.com/jinzhu/gorm"

	"github.com/YahuiAn/Go-bjut/model"

	"github.com/gin-gonic/gin"
)

type commentInfo struct {
	DynamicID  uint   `binding:"required"`
	Content    string `binding:"required,max=255"`
	ReplyFloor uint
}

func Create(c *gin.Context) {
	var info commentInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": tip.Warn(err)})
		return
	}

	if _, err := model.GetDynamicByID(info.DynamicID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "动态不存在，评论错误"})
		return
	}

	// 计算楼层
	// 查找评论中的最后一层楼
	var lastComment model.Comment
	err := model.DB.Where("dynamic_id = ?", info.DynamicID).Order("floor").Find(&lastComment).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "楼层查询错误，评论失败"})
		return
	}

	if info.ReplyFloor > lastComment.Floor {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "楼层不存在，评论错误"})
		return
	}

	who := user.CurrentUser(c)
	if who == (model.User{}) {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取当前用户失败"})
		return
	}

	comment := model.Comment{
		DynamicID:   info.DynamicID,
		Floor:       lastComment.Floor + 1,
		Commentator: who.NickName,
		Content:     info.Content,
		ReplyFloor:  info.ReplyFloor,
	}

	err = model.DB.Create(&comment).Error
	if err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库插入错误，评论失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "评论成功"})
}
