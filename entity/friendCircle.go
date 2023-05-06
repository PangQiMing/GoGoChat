package entity

import (
	"gorm.io/gorm"
)

type FriendCircle struct {
	gorm.Model
	FriendCircleID uint64 `gorm:"not null" json:"friend_circle_id"`
	GoGoID         uint64 `gorm:"not null" json:"-"`
	Nickname       string `gorm:"varchar(255)" json:"nickname"`
	Content        string `gorm:"type:text" json:"content"`
	Picture        string `gorm:"type:longtext" json:"picture"`
}
