package entity

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	GroupID      uint64 `gorm:"not null" json:"group_id"`              //群号
	Avatar       string `gorm:"not null" json:"avatar"`                //群头像
	Name         string `gorm:"type:varchar(50);not null" json:"name"` //群名
	Announcement string `gorm:"type:varchar(255)" json:"announcement"` //群公告
	GroupOwnerID uint64 `gorm:"not null" json:"group_owner_id"`        //群主ID
	MemberID     uint64 `gorm:"not null" json:"member_id"`             //群组成员ID
	Status       uint   `gorm:"not null" json:"status"`                //加入群组状态，0待处理，1同意，2拒绝
	Members      User   `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
