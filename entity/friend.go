package entity

import (
	"gorm.io/gorm"
	"time"
)

type Friend struct {
	GoGoID    uint64 `gorm:"not null"`                  //用户ID
	FriendID  uint64 `gorm:"not null"`                  // 好友ID
	Status    uint   `gorm:"default:0"`                 //好友关系状态 0代表待处理，1代表同意，2代表拒绝
	Remark    string `gorm:"type:varchar(20);not null"` // 好友备注
	Friend    User   `gorm:"foreignKey:FriendID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"friend"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
