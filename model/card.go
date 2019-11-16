package model

import "github.com/jinzhu/gorm"

type Card struct {
	gorm.Model
	Registrant string
	Name       string
	Sex        string
	College    string
	StuNumber  string
	Location   string
}
