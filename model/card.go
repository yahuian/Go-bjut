package model

import "github.com/jinzhu/gorm"

type Card struct {
	gorm.Model
	Registrant string
	RealName   string
	Sex        string
	College    string
	StuNumber  string
	Location   string
	Flag       bool // true表示已经给失主发送了短信
}
