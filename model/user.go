package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	NickName  string `gorm:"unique_index;not null"`
	Password  string
	Email     *string `gorm:"unique_index"`
	Telephone *string `gorm:"unique_index"`
	College   string
	Major     string
	ClassName string
	Number    *string `gorm:"unique_index"` // 学号或职工号
	RealName  string
}
