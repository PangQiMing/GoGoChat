package service

import (
	"errors"
	"github.com/PangQiMing/GoGoChat/config"
	"github.com/PangQiMing/GoGoChat/dto"
	"github.com/PangQiMing/GoGoChat/entity"
	"log"
)

// GetFriendList 获取好友列表数据
func GetFriendList(goGoID uint64) ([]entity.Friend, error) {
	var friendList []entity.Friend
	err := config.DB.Preload("Friend").Where("go_go_id = ? AND status = ?", goGoID, 1).Find(&friendList).Error
	if err != nil {
		return nil, err
	}
	return friendList, nil
}

// AddFriend 添加好友
func AddFriend(addFriendDTO dto.AddFriendDTO) error {
	// 检查好友关系是否已存在
	var count int64
	config.DB.Model(&entity.Friend{}).Where("go_go_id = ? AND friend_id = ?", addFriendDTO.FriendID, addFriendDTO.GoGoID).Count(&count)
	log.Println(count)
	if count > 0 {
		return errors.New("好友关系已存在")
	}

	// 添加好友关系
	friendList := entity.Friend{
		GoGoID:   addFriendDTO.FriendID,
		FriendID: addFriendDTO.GoGoID,
		Status:   0,
		Remark:   "",
	}
	return config.DB.Create(&friendList).Error
}

// GetFriendRequestList 获取好友请求列表
func GetFriendRequestList(goGoID uint64) ([]entity.Friend, error) {
	var friendRequestList []entity.Friend
	err := config.DB.Preload("Friend").Where("go_go_id = ? AND status = ?", goGoID, 0).Find(&friendRequestList).Error
	if err != nil {
		return []entity.Friend{}, err
	}
	return friendRequestList, nil
}

// AcceptFriendRequest 同意好友请求，设置好友关系状态 Status = 1
func AcceptFriendRequest(acceptFriendDTO dto.AcceptFriendDTO) error {
	//log.Println(acceptFriendDTO.GoGoID)
	//log.Println(acceptFriendDTO.FriendID)
	// 检查好友关系是否已存在
	var count int64
	config.DB.Model(&entity.Friend{}).Where("go_go_id = ? AND friend_id = ? AND status=1", acceptFriendDTO.GoGoID, acceptFriendDTO.FriendID).Count(&count)
	log.Println(count)
	if count > 0 {
		return errors.New("已同意")
	}
	oneself := entity.Friend{
		GoGoID:   acceptFriendDTO.FriendID,
		FriendID: acceptFriendDTO.GoGoID,
		Status:   1,
		Remark:   "",
	}
	return config.DB.Model(&entity.Friend{}).Create(&oneself).Where("go_go_id = ? AND friend_id = ?", acceptFriendDTO.GoGoID, acceptFriendDTO.FriendID).Update("status", 1).Error
}

// RejectFriendRequest 拒绝好友请求,并删除好友关系状态
func RejectFriendRequest(rejectFriendDTO dto.RejectFriendDTO) error {
	return config.DB.Where("go_go_id = ? AND friend_id = ?", rejectFriendDTO.GoGoID, rejectFriendDTO.FriendID).Delete(&entity.Friend{}).Error
}

// DeleteFriend 删除好友
func DeleteFriend(deleteFriendDTO dto.DeleteFriendDTO) error {
	config.DB.Where("go_go_id = ? AND friend_id = ?", deleteFriendDTO.GoGoID, deleteFriendDTO.FriendID).Delete(&entity.Friend{})
	return config.DB.Where("go_go_id = ? AND friend_id = ?", deleteFriendDTO.FriendID, deleteFriendDTO.GoGoID).Delete(&entity.Friend{}).Error
}

// SearchFriendByFriendID 查找好友是否存在
func SearchFriendByFriendID(searchFriendDTO dto.SearchFriendDTO) (entity.User, error) {
	var friend entity.User
	err := config.DB.Where("go_go_id = ?", searchFriendDTO.FriendID).Take(&friend).Error
	if err != nil {
		return entity.User{}, err
	}
	return friend, nil
}
