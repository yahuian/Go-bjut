package card

import (
	"net/http"
	"strconv"
	"time"

	"github.com/YahuiAn/Go-bjut/model"

	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/gin-gonic/gin"
)

type cardDisplay struct {
	RealName   string
	Sex        string
	College    string
	Number     string
	Location   string
	Registrant string
	CreatedAt  time.Time
	Status     string
}

// golang标准库struct time，默认是以rfc3339的格式进行序列化反序列化的
// 如果想要用别的时间格式，就得重写MarshalJSON、UnmarshalJSON方法，不过这样就要动别处的代码
// 权衡后规定：前后端统一用rfc3339的格式传递时间，对于时间的友好显示操作由前端来处理

func Index(c *gin.Context) {
	pageIndex, err := strconv.Atoi(c.Query("page")) // 页号
	if err != nil {
		logger.Error.Println(err)
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
	if err := model.DB.Table("cards").Offset(offset).Limit(pageSize).Order("created_at desc").Scan(&cards).Error; err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}

	var rows int // 表的总行数
	if err := model.DB.Table("cards").Count(&rows).Error; err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}

	pages := rows/pageSize + 1 // 总共有多少页

	c.JSON(http.StatusOK, gin.H{"msg": "查询成功", "data": cards, "pages": pages})
}
