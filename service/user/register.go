package user

import (
	"net/http"

	"github.com/YahuiAn/Go-bjut/tip"

	"github.com/YahuiAn/Go-bjut/logger"

	"golang.org/x/crypto/bcrypt"

	"github.com/YahuiAn/Go-bjut/model"

	"github.com/gin-gonic/gin"
)

type RegisterInfo struct {
	NickName   string `binding:"required,min=2,max=30"`
	Password   string `binding:"required,min=8,max=40"`
	PwdConfirm string `binding:"eqfield=Password"`
}

// 用户注册
func Register(c *gin.Context) {
	var info RegisterInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": tip.Warn(err)})
		return
	}

	// 检查是否已经注册
	exist, err := model.ExistUserByUniqueField("nick_name", info.NickName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}
	if exist {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "该昵称已被占用"})
		return
	}

	bytesPwd, err := bcrypt.GenerateFromPassword([]byte(info.Password), 10)
	if err != nil {
		logger.Error.Println("密码加密失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "密码加密失败"})
		return
	}

	user := model.User{
		NickName: info.NickName,
		Password: string(bytesPwd),
	}

	// 插入数据
	err = model.DB.Create(&user).Error
	if err != nil {
		logger.Error.Println("注册失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "注册失败"})
		return
	}

	logger.Info.Println("注册成功", info.NickName)
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功", "data": user.NickName})
}

// TODO 增加通过微信注册的方式
