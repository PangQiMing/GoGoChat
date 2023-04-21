package entity

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	GoGoID   uint64 `gorm:"not null" json:"go_go_id"`                //用户ID
	FriendID uint64 `gorm:"not null" json:"friend_id"`               // 好友ID
	Status   uint   `gorm:"default:0" json:"status"`                 //好友关系状态 0代表待处理，1代表同意，2代表拒绝
	Remark   string `gorm:"type:varchar(20);not null" json:"remark"` // 好友备注
	Friend   User   `gorm:"foreignKey:FriendID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"friend"`
}
