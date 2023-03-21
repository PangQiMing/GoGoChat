package service

import (
	"errors"
	"github.com/PangQiMing/GoGoChat/config"
	"github.com/PangQiMing/GoGoChat/dto"
	"github.com/PangQiMing/GoGoChat/entity"
	"github.com/mashingan/smapping"
)

func RegisterUser(registerUserDTO dto.RegisterUserDTO) error {
	var user entity.User
	err := smapping.FillStruct(&user, smapping.MapFields(&registerUserDTO))
	if err != nil {
		return err
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Login(loginUserDTO dto.LoginUserDTO) error {
	var count int64
	config.DB.Model(entity.User{}).Where("email = ? AND password = ?", loginUserDTO.Email, loginUserDTO.Password).Count(&count)
	if count > 0 {

	}
	return errors.New("用户密码错误")
}
