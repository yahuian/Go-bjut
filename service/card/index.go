package card

import (
	"net/http"
	"strconv"
	"time"

	"github.com/YahuiAn/Go-bjut/database"
	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/gin-gonic/gin"
)

type cardDisplay struct {
	RealName   string
	Sex        string
	College    string
	StuNumber  string
	Location   string
	Registrant string
	CreatedAt  time.Time // TODO 返回前端可读性良好的时间格式
	Status     string
}

func Index(c *gin.Context) {
	pageIndex, err := strconv.Atoi(c.Query("page")) // 页号
	if err != nil {
		logger.Error.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "不合法的页数"})
		return
	}
	if pageIndex < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "不合法的页数"})
		return
	}
	pageSize := 10 // 每页的大小

	offset := (pageIndex - 1) * pageSize
	var cards []cardDisplay
	if err := database.DB.Table("cards").Offset(offset).Limit(pageSize).Scan(&cards).Error; err != nil {
		logger.Error.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}

	var rows int // 表的总行数
	if err := database.DB.Table("cards").Count(&rows).Error; err != nil {
		logger.Error.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}

	pages := rows/pageSize + 1 // 总共有多少页

	c.JSON(http.StatusOK, gin.H{"msg": "查询成功", "data": cards, "pages": pages})
}
