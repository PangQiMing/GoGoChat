package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	GoGoID       uint64 `gorm:"primaryKey" json:"go_go_id"`                                 //用户账号
	Password     string `gorm:"size:255;not null" json:"password"`                          //密码
	AvatarURL    string `gorm:"type:varchar(255);default:'images/1.jpg'" json:"avatar_url"` //头像URL
	Nickname     string `gorm:"type:varchar(50);not null" json:"nickname"`                  //昵称
	Sex          string `gorm:"not null;default:保密" json:"sex"`                             //性别
	Age          string `gorm:"not null;default:0" json:"age"`                              //年龄
	Introduction string `gorm:"not null;default:''" json:"introduction"`                    //个人简介
	Status       uint   `gorm:"not null;default:0" json:"status"`                           //离线状态为0，在线状态为1
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
