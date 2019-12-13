package card

import (
	"github.com/YahuiAn/Go-bjut/model"

	"github.com/YahuiAn/Go-bjut/logger"
)

type smsInfo struct {
	ID         uint
	Number     string
	Telephone  string
	Location   string
	Registrant string
}

// TODO 如何防止破坏党多次发送短信(登记丢失卡片时进行筛选和判断)
func Notice() {
	var info []smsInfo

	err := model.DB.Table("users as u").
		Select("c.id, u.number, u.telephone, c.location, c.registrant").
		Joins("join cards as c on u.number = c.number and c.status != ?", model.SuccessfulNotification).
		Scan(&info).Error

	if err != nil {
		logger.Error.Println(err)
		return
	}

	for i := range info {
		var temp model.Card

		if info[i].Telephone == "" {
			model.DB.Table("cards").First(&temp, info[i].ID).Update("status", model.UnboundPhone)
			continue
		}

		if !sendMessage(info[i].Telephone, info[i].Number, info[i].Location, info[i].Registrant) {
			model.DB.Table("cards").First(&temp, info[i].ID).Update("status", model.SmsAPIError)
		} else {
			model.DB.Table("cards").First(&temp, info[i].ID).Update("status", model.SuccessfulNotification)
		}
	}
}
