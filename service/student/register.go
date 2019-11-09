package student

import (
	"net/http"

	"github.com/YahuiAn/Go_bjut/logger"

	"golang.org/x/crypto/bcrypt"

	"github.com/YahuiAn/Go_bjut/model"

	"github.com/YahuiAn/Go_bjut/database"

	"github.com/gin-gonic/gin"
)

type RegisterInfo struct {
	Nickname   string `json:"nickname" binding:"required,min=2,max=30"`
	Password   string `json:"password" binding:"required,min=8,max=40"`
	PwdConfirm string `json:"pwdConfirm" binding:"required,min=8,max=40"`
}

// 学生注册
func Register(c *gin.Context) {
	var info RegisterInfo
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

	count := 0
	err := database.DB.Model(&model.Student{}).Where("nick_name = ?", info.Nickname).Count(&count).Error
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

	bytesPwd, err := bcrypt.GenerateFromPassword([]byte(info.Password), 10)
	if err != nil {
		logger.Error.Println("密码加密失败", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "密码加密失败"})
		return
	}

	student := model.Student{
		User: model.User{
			NickName: info.Nickname,
			Password: string(bytesPwd),
		},
	}

	// 插入数据
	err = database.DB.Create(&student).Error
	if err != nil {
		logger.Error.Println("注册失败", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "注册失败"})
		return
	}

	logger.Info.Println("注册成功", info.Nickname)
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功", "data": student.NickName})
}
