package service

import (
	"github.com/PangQiMing/GoGoChat/config"
	"github.com/PangQiMing/GoGoChat/entity"
)

// SaveMessage 保存message到数据库
func SaveMessage(message entity.Message) error {
	return config.DB.Create(&message).Error
}

func GetAllMessageByGoGoID(goGoGoID uint64) []entity.Message {
	var message []entity.Message
	config.DB.Find(&message)
	return message
}
