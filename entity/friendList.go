package entity

import "gorm.io/gorm"

type FriendList struct {
	gorm.Model
	UserID   uint   // 用户ID，外键，关联用户表
	FriendID uint   // 好友ID，外键，关联用户表
	Remark   string // 好友备注
}
