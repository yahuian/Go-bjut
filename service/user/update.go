package user

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/YahuiAn/Go-bjut/model"

	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/gin-gonic/gin"
)

// 可供用户更新的数据
type updateInfo struct {
	NickName    string `binding:"omitempty,min=2,max=30"`
	Password    string `binding:"omitempty,min=8,max=40"`
	NewPassword string `binding:"required_with=Password,omitempty,min=8,max=40"`
	PwdConfirm  string `binding:"eqfield=NewPassword"`
	Email       string
	Telephone   string
	College     string
	Major       string
	ClassName   string
	Number      string
	RealName    string
}

func Update(c *gin.Context) {
	var info updateInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	user := CurrentUser(c)
	if user == (model.User{}) {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "当前用户查询失败"})
		return
	}

	if info.Password != "" {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(info.Password)); err != nil {
			logger.Error.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"msg": "原密码错误"})
			return
		}

		bytesPwd, err := bcrypt.GenerateFromPassword([]byte(info.NewPassword), 10)
		if err != nil {
			logger.Error.Println("密码加密失败", err)
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "密码加密失败"})
			return
		}
		info.Password = string(bytesPwd)
	}

	if info.NickName != "" {
		count := 0
		err := model.DB.Model(&model.User{}).Where("nick_name = ?", info.NickName).Count(&count).Error
		if err != nil {
			logger.Error.Println("数据库查询失败", err)
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
			return
		}
		if count > 0 {
			logger.Error.Println("该昵称已被占用")
			c.JSON(http.StatusBadRequest, gin.H{"msg": "该昵称已被占用"})
			return
		}
	}

	// TODO 对于Email，telephone信息更新时做安全检查和身份认证
	// Update multiple attributes with `struct`, will only update those changed & non blank fields
	// 更新用户信息
	if err := model.DB.Model(&user).Updates(info).Error; err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "信息更新失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "信息更新成功"})
}
