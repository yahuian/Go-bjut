package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	DynamicID   uint
	Floor       uint
	Commentator string `gorm:"not null"`
	Content     string
	ReplyFloor  uint // ReplyFloor为0表示回复的是当前动态
}
