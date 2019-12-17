package model

import (
	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/jinzhu/gorm"
)

type Dynamic struct {
	gorm.Model
	NickName string `gorm:"not null"`
	Title    string
	Content  string `gorm:"type:text"`
}

func GetDynamicByID(ID uint) (Dynamic, error) {
	var dynamic Dynamic
	err := DB.Where("id = ?", ID).First(&dynamic).Error
	if err != nil {
		logger.Error.Println(err)
		return dynamic, err
	}
	return dynamic, nil
}
