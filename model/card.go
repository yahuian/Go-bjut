package model

import "github.com/jinzhu/gorm"

type Card struct {
	gorm.Model
	Registrant string
	RealName   string
	Sex        string
	College    string
	Number     string `gorm:"unique_index;not null"`
	Location   string
	Status     string
}

// Status字段的值
const (
	WaitingNotification    = "等待通知"
	SuccessfulNotification = "成功通知"
	UnboundPhone           = "未绑定电话"
	SmsAPIError            = "短信API错误"
)
