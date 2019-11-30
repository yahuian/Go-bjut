package user

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/YahuiAn/Go-bjut/model"

	"github.com/YahuiAn/Go-bjut/database"

	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/gin-gonic/gin"
)

// 可供用户更新的数据
type updateInfo struct {
	NickName   string `null,binding:"min=2,max=30"`
	Password   string `null,binding:"min=8,max=40"`
	PwdConfirm string `null,binding:"min=8,max=40"`
	Email      string
	Telephone  string
	College    string
	Major      string
	ClassName  string
	Number     string
	RealName   string
}

func Update(c *gin.Context) {
	var info updateInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		logger.Error.Println("json信息错误", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "json信息错误"}) // TODO 具体化错误信息
		return
	}

	if info.Password != info.PwdConfirm {
		logger.Error.Println("两次输入密码不一致")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "两次输入密码不一致"})
		return
	}

	if info.NickName != "" {
		count := 0
		err := database.DB.Model(&model.User{}).Where("nick_name = ?", info.NickName).Count(&count).Error
		if err != nil {
			logger.Error.Println("数据库查询失败", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
			return
		}
		if count > 0 {
			logger.Error.Println("该昵称已被占用")
			c.JSON(http.StatusBadRequest, gin.H{"msg": "该昵称已被占用"})
			return
		}
	}

	if info.Password != "" {
		bytesPwd, err := bcrypt.GenerateFromPassword([]byte(info.Password), 10)
		if err != nil {
			logger.Error.Println("密码加密失败", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "密码加密失败"})
			return
		}
		info.Password = string(bytesPwd)
	}

	// TODO 对于Email，telephone，password信息更新时做安全检查和身份认证

	user := CurrentUser(c)
	if user == (model.User{}) {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "查询登录用户失败"})
		return
	}

	// Update multiple attributes with `struct`, will only update those changed & non blank fields
	// 更新用户信息
	if err := database.DB.Model(&user).Updates(info).Error; err != nil {
		logger.Error.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "信息更新失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "信息更新成功"})
}
