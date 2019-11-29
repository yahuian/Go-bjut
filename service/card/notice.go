package card

import (
	"github.com/YahuiAn/Go-bjut/model"

	"github.com/YahuiAn/Go-bjut/database"
	"github.com/YahuiAn/Go-bjut/logger"
)

type smsInfo struct {
	ID         uint
	StuNumber  string
	Telephone  string
	Location   string
	Registrant string
}

// TODO 如何防止破坏党多次发送短信(登记丢失卡片时进行筛选和判断)
func Notice() {
	var info []smsInfo

	err := database.DB.Table("students as s").
		Select("c.id, s.stu_number, s.telephone, c.location, c.registrant").
		Joins("join cards as c on s.stu_number = c.stu_number and c.status != ?", model.SuccessfulNotification).
		Scan(&info).Error

	if err != nil {
		logger.Error.Println(err.Error())
		return
	}

	for i := range info {
		var temp model.Card

		if info[i].Telephone == "" {
			database.DB.Table("cards").First(&temp, info[i].ID).Update("status", model.UnboundPhone)
			continue
		}

		if !sendMessage(info[i].Telephone, info[i].StuNumber, info[i].Location, info[i].Registrant) {
			database.DB.Table("cards").First(&temp, info[i].ID).Update("status", model.SmsAPIError)
		} else {
			database.DB.Table("cards").First(&temp, info[i].ID).Update("status", model.SuccessfulNotification)
		}
	}
}
