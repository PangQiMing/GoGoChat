package service

import (
	"errors"
	"github.com/PangQiMing/GoGoChat/config"
	"github.com/PangQiMing/GoGoChat/dto"
	"github.com/PangQiMing/GoGoChat/entity"
	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

// CreateGroup 创建群组
func CreateGroup(groupOwnerID uint64, createGroupDTO dto.CreateGroupDTO) error {
	var group entity.Group
	err := smapping.FillStruct(&group, smapping.MapFields(&createGroupDTO))
	if err != nil {
		return err
	}
	group.GroupID = uint64(uuid.New().ID())
	group.Avatar = "images/defaultGroup.jpg"
	group.GroupOwnerID = groupOwnerID
	group.MemberID = groupOwnerID
	group.Status = 1
	tx := config.DB.Create(&group)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// UpdateGroup 更新群组信息
func UpdateGroup(groupOwnerID uint64, updateGroupDTO dto.UpdateGroupDTO) error {
	//查询群组是否存在，只有群主才可以更新群信息
	var count int64
	tx := config.DB.Model(&entity.Group{}).Where("group_id = ? AND group_owner_id = ?", updateGroupDTO.GroupID, groupOwnerID).Count(&count)
	if tx.Error != nil {
		return tx.Error
	}

	if count < 1 {
		return errors.New("群组不存在或者你不是群主")
	}

	var group entity.Group
	err := smapping.FillStruct(&group, smapping.MapFields(&updateGroupDTO))
	if err != nil {
		return err
	}
	group.GroupOwnerID = groupOwnerID
	group.MemberID = groupOwnerID
	tx = config.DB.Where("group_id = ?", group.GroupID).Updates(&group)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteGroup 解散群组
func DeleteGroup(groupOwnerID uint64, deleteGroup dto.DeleteGroupDTO) error {
	//查询群组是否存在，只有群主才可以解散群组
	var count int64
	tx := config.DB.Model(&entity.Group{}).Where("group_id = ? AND group_owner_id = ?", deleteGroup.GroupID, groupOwnerID).Count(&count)
	if tx.Error != nil {
		return tx.Error
	}

	if count < 1 {
		return errors.New("群组不存在或者你不是群主")
	}

	var group entity.Group
	group.GroupID = deleteGroup.GroupID
	group.GroupOwnerID = groupOwnerID
	tx = config.DB.Where("group_id = ? AND group_owner_id = ?", group.GroupID, group.GroupOwnerID).Delete(&group)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteGroupMember 群成员退出群组
func DeleteGroupMember(memberID uint64, memberDTO dto.DeleteGroupMemberDTO) error {
	return config.DB.Where("group_id = ? AND member_id = ?", memberDTO.GroupID, memberID).Delete(&entity.Group{}).Error
}

// GetGroupLists 获取群组列表
func GetGroupLists(goGoID uint64) ([]entity.Group, error) {
	var group []entity.Group
	tx := config.DB.Preload("Members").Where("member_id = ? AND status = ?", goGoID, 1).Find(&group)
	if tx.Error != nil {
		return []entity.Group{}, tx.Error
	}
	return group, nil
}

// JoinGroupRequestList 获取入群的申请列表
func JoinGroupRequestList(groupOwnerID uint64) ([]entity.Group, error) {
	var group []entity.Group
	tx := config.DB.Preload("Members").Where("group_owner_id = ? AND status = 0", groupOwnerID).Find(&group)
	if tx.Error != nil {
		return []entity.Group{}, tx.Error
	}
	return group, nil
}

// JoinGroup 申请加入群组
func JoinGroup(joinGroupDTO dto.JoinGroupDTO) error {
	var group entity.Group
	err := smapping.FillStruct(&group, smapping.MapFields(&joinGroupDTO))
	if err != nil {
		return err
	}

	var count int64
	config.DB.Model(&entity.Group{}).Where("group_id = ? AND member_id = ?", group.GroupID, group.MemberID).Count(&count)
	if count > 0 {
		return errors.New("已发送加群请求或已是该群成员")
	}

	var groupInfo entity.Group
	tx := config.DB.Model(&entity.Group{}).Where("group_id = ?", joinGroupDTO.GroupID).Take(&groupInfo)
	if tx.Error != nil {
		return tx.Error
	}

	var joinGroup entity.Group
	joinGroup.GroupID = groupInfo.GroupID
	joinGroup.Avatar = groupInfo.Avatar
	joinGroup.Name = groupInfo.Name
	joinGroup.Announcement = groupInfo.Announcement
	joinGroup.GroupOwnerID = groupInfo.GroupOwnerID
	joinGroup.MemberID = joinGroupDTO.MemberID
	joinGroup.Status = 0

	tx = config.DB.Create(&joinGroup)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// AcceptJoinGroupRequest 同意加入群组请求
func AcceptJoinGroupRequest(acceptJoinGroupDTO dto.AcceptJoinGroupDTO) error {
	return config.DB.Model(&entity.Group{}).Where("group_id = ? AND group_owner_id = ? AND member_id = ?",
		acceptJoinGroupDTO.GroupID, acceptJoinGroupDTO.GroupOwnerID, acceptJoinGroupDTO.MemberID).Update("status", 1).Error
}

// RejectJoinGroupRequest 拒绝加入群组请求
func RejectJoinGroupRequest(groupOwnerID uint64, rejectJoinGroupDTO dto.RejectJoinGroupDTO) error {
	return config.DB.Model(&entity.Group{}).Where("group_id = ? AND group_owner_id = ? AND member_id = ?",
		rejectJoinGroupDTO.GroupID, groupOwnerID, rejectJoinGroupDTO.MemberID).Update("status", 2).Error
}

// GetSearchGroup 查找群组是否存在
func GetSearchGroup(searchGroupDTO dto.SearchGroupDTO) (entity.Group, error) {
	var group entity.Group
	tx := config.DB.Where("group_id = ?", searchGroupDTO.GroupID).Take(&group)
	if tx.Error != nil {
		return entity.Group{}, tx.Error
	}
	return group, nil
}

func GetGroupMemberByToId(groupID string) []entity.Group {
	var group []entity.Group
	config.DB.Where("group_id = ?", groupID).Find(&group)
	return group
}
