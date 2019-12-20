package comment

import (
	"net/http"
	"strconv"
	"time"

	"github.com/YahuiAn/Go-bjut/logger"

	"github.com/YahuiAn/Go-bjut/model"

	"github.com/gin-gonic/gin"
)

type DisplayComment struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Floor       uint
	Commentator string
	Content     string
	ReplyFloor  uint
}

func GetCommentByDynamicID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "错误的动态ID"})
		return
	}

	var display []DisplayComment

	if err := model.DB.Table("comments").Where("dynamic_id = ?", id).Scan(&display).Order("floor").Error; err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "获取评论成功", "data": display})
}
