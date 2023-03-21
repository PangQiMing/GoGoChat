package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	AvatarURL string `gorm:"type:text" json:"avatar_url"`                         //头像URL
	Nickname  string `gorm:"type:varchar(50);not null" json:"nickname"`           //昵称
	Email     string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"` //邮箱，设置为唯一索引
	Password  string `gorm:"size:255;not null" json:"password"`                   //密码
	Sex       string `gorm:"not null;default:未知" json:"sex"`                      //性别
	Age       uint   `gorm:"not null;default:99" json:"age"`                      //年龄
	Status    uint   `gorm:"not null;default:0" json:"status"`                    //离线状态为0，在线状态为1
}
