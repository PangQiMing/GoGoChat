package service

import (
	"github.com/PangQiMing/GoGoChat/config"
	"github.com/PangQiMing/GoGoChat/dto"
	"github.com/PangQiMing/GoGoChat/entity"
	"github.com/mashingan/smapping"
)

func CreateFriendCircle(momentCreateDTO dto.CircleCreateDTO) error {
	var circle entity.FriendCircle
	err := smapping.FillStruct(&circle, smapping.MapFields(&momentCreateDTO))
	if err != nil {
		return err
	}
	return config.DB.Create(&circle).Error
}

func DeleteFriendCircle(goGoID uint64, circleDTO dto.DeleteFriendCircleDTO) error {
	return config.DB.Where("go_go_id = ? AND friend_circle_id = ?", goGoID, circleDTO.FriendCircleID).Delete(&entity.FriendCircle{}).Error
}

func GetFriendCircleAll(goGoID uint64) ([]entity.FriendCircle, error) {
	var friendCircle []entity.FriendCircle
	err := config.DB.Where("go_go_id = ?", goGoID).Find(&friendCircle).Error
	if err != nil {
		return []entity.FriendCircle{}, err
	}
	return friendCircle, nil
}
