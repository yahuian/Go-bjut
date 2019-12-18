package dynamic

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/YahuiAn/Go-bjut/model"

	"github.com/gin-gonic/gin"
)

type DisplayDynamic struct {
	ID       uint
	Time     time.Time
	NickName string
	Title    string
	Content  string
}

func GetDynamicById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "错误的动态ID"})
		return
	}

	var dynamic model.Dynamic
	dynamic, err = model.GetDynamicByID(uint(id))
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "此条动态不存在"})
		}
		return
	}

	var active time.Time
	// 确定活跃时间，如果动态有更新，则显示更新时间，否则显示创建时间
	if dynamic.UpdatedAt.String() != "" {
		active = dynamic.UpdatedAt
	} else {
		active = dynamic.CreatedAt
	}

	display := DisplayDynamic{
		ID:       dynamic.ID,
		Time:     active,
		NickName: dynamic.NickName,
		Title:    dynamic.Title,
		Content:  dynamic.Content,
	}

	c.JSON(http.StatusOK, gin.H{"msg": "获取动态成功", "date": display})
}
