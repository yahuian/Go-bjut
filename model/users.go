package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	NickName  string `gorm:"unique;not null"`
	Password  string
	Email     string
	Telephone string
}

type Student struct {
	User
	College   string
	Major     string
	ClassName string
	StuNumber string `gorm:"unique"`
	RealName  string
}
