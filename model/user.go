package model

import (
	"github.com/YahuiAn/Go-bjut/logger"
	"github.com/jinzhu/gorm"
)

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

func ExistUserByUniqueField(field, value string) (bool, error) {
	var user User
	err := DB.Select("id").Where(gorm.ToColumnName(field)+" = ?", value).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error.Println(err)
		return false, err
	}
	if user.ID > 0 {
		logger.Info.Printf("field:%s,value:%s.\n", field, value)
		return true, nil
	}
	return false, nil
}
