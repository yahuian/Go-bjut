package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/YahuiAn/Go-bjut/model"
	"golang.org/x/crypto/bcrypt"

	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/gin-gonic/gin"
)

// 用于北工大学子的注册函数
// 用户信息从“工大小美”项目的接口中获取
// 返回的json数据格式如下
/*
{
    "class": "160701",
    "college": "信息学部",
    "major": "计算机科学与技术",
    "stuNum": "16010328",
    "sutName": "安亚辉"
}
*/
// https://github.com/YahuiAn/BJUTservice
// TODO 小美 该接口目前仅适用于北工大的本科生，考虑开发其他人群的接口，如老师，研究生等等

const baseInfoAPI = "https://bjut1960.cn/baseinfo"

type baseInfo struct {
	College   string `json:"college"`
	Major     string `json:"major"`
	ClassName string `json:"class"`
	Number    string `json:"stuNum"`
	RealName  string `json:"sutName"` // sutName这个锅该小美API来背
}

// 登录bjut正方教务系统所需信息
type stuInfo struct {
	Number   string
	Password string
}

func BjutRegister(c *gin.Context) {
	var loginInfo stuInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		logger.Error.Println("json信息错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "json信息错误"}) // TODO 具体化错误信息
		return
	}

	// 注册时，用StuNumber作为NickName
	// 检查是否已经注册
	count := 0
	err := model.DB.Model(&model.User{}).Where("number = ?", loginInfo.Number).Count(&count).Error
	if err != nil {
		logger.Error.Println("数据库查询失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "数据库查询失败"})
		return
	}
	if count > 0 {
		logger.Error.Println("该学号已经被注册")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "该学号已经被注册"})
		return
	}

	// 向小美API发起HTTP POST请求
	data := url.Values{"xh": {loginInfo.Number}, "mm": {loginInfo.Password}}
	resp, err := http.PostForm(baseInfoAPI, data)
	if err != nil {
		logger.Error.Printf("%s请求失败,%s\n", baseInfoAPI, err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": baseInfoAPI + "接口请求失败"})
		return
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		logger.Error.Printf(resp.Status)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "请检查教务账号密码是否有误"})
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": baseInfoAPI + "接口请求失败"})
		return
	}

	// 序列化数据
	var info baseInfo
	if err := json.Unmarshal(body, &info); err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": baseInfoAPI + "接口返回数据格式错误"})
		return
	}

	// 使用教务管理系统的密码作为用户密码
	bytesPwd, err := bcrypt.GenerateFromPassword([]byte(loginInfo.Password), 10)
	if err != nil {
		logger.Error.Println("密码加密失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "密码加密失败"})
		return
	}

	user := model.User{
		NickName:  info.Number,
		Password:  string(bytesPwd),
		College:   info.College,
		Major:     info.Major,
		ClassName: info.ClassName,
		Number:    info.Number,
		RealName:  info.RealName,
	}

	// 插入数据
	err = model.DB.Create(&user).Error
	if err != nil {
		logger.Error.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "注册失败"})
		return
	}

	logger.Info.Println("注册成功", info.Number)
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功", "data": info.Number})
}
