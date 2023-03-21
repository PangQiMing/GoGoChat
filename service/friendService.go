package service

import (
	"errors"
	"github.com/PangQiMing/GoGoChat/config"
	"github.com/PangQiMing/GoGoChat/entity"
)

// GetFriendList 获取好友列表数据
func GetFriendList(userID uint) ([]entity.FriendList, error) {
	var friendList []entity.FriendList
	err := config.DB.Where("user_id = ?", userID).Find(&friendList).Error
	if err != nil {
		return nil, err
	}
	return friendList, nil
}

// AddFriend 添加好友
func AddFriend(userID, friendID uint) error {
	// 检查好友关系是否已存在
	var count int64
	config.DB.Model(&entity.FriendList{}).Where("user_id = ? AND friend_id = ?", userID, friendID).Count(&count)
	if count > 0 {
		return errors.New("好友关系已存在")
	}
	// 添加好友关系
	friendList := entity.FriendList{
		UserID:   userID,
		FriendID: friendID,
	}
	return config.DB.Create(&friendList).Error
}

// DeleteFriend 删除好友
func DeleteFriend(userID, friendID uint) error {
	return config.DB.Where("user_id = ? AND friend_id = ?", userID, friendID).Delete(&entity.FriendList{}).Error
}
