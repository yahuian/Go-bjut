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
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	NickName  string
	Title     string
	Content   string
}

func GetDynamicByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "错误的动态ID"})
		return
	}

	var dynamic model.Dynamic
	dynamic, err = model.GetDynamicByID(uint(id))
	if err != nil {
		// gorm在查询单条数据时，如果数据库中不存在该数据，会返回gorm.ErrRecordNotFound的错
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "此条动态不存在"})
		}
		return
	}

	display := DisplayDynamic{
		ID:        dynamic.ID,
		CreatedAt: dynamic.CreatedAt,
		UpdatedAt: dynamic.UpdatedAt,
		NickName:  dynamic.NickName,
		Title:     dynamic.Title,
		Content:   dynamic.Content,
	}

	c.JSON(http.StatusOK, gin.H{"msg": "获取动态成功", "date": display})
}
