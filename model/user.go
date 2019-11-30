package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	NickName  string `gorm:"unique;not null"`
	Password  string
	Email     string
	Telephone string
	College   string
	Major     string
	ClassName string
	Number    string `gorm:"unique"` // 学号或职工号
	RealName  string
}
