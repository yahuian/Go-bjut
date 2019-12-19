package dynamic

import (
	"net/http"
	"strconv"

	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/YahuiAn/Go-bjut/model"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	pageIndex, err := strconv.Atoi(c.DefaultQuery("page", "1")) // 页号
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "不合法的页数"})
		return
	}
	if pageIndex < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "不合法的页数"})
		return
	}
	pageSize := 30 // 每页的大小

	offset := (pageIndex - 1) * pageSize
	// gorm在查询多条数据时，即使没有找到数据，也不会报错，这点和查询单条数据时不一样，算是一个小坑
	var dynamics []model.Dynamic
	fields := []string{"id", "created_at", "updated_at", "nick_name", "title"} // 主页不显示动态的正文内容
	err = model.DB.Select(fields).Offset(offset).Limit(pageSize).Order("updated_at desc").Find(&dynamics).Error
	if err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}

	var rows int // 表的总行数
	if err := model.DB.Table("dynamics").Count(&rows).Error; err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}

	pages := rows/pageSize + 1 // 总共有多少页

	c.JSON(http.StatusOK, gin.H{"msg": "查询成功", "data": dynamics, "pages": pages})
}
