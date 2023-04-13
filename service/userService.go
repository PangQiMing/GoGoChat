package service

import (
	"errors"
	"github.com/PangQiMing/GoGoChat/config"
	"github.com/PangQiMing/GoGoChat/dto"
	"github.com/PangQiMing/GoGoChat/entity"
	"github.com/PangQiMing/GoGoChat/utils"
	"github.com/mashingan/smapping"
)

// RegisterUser 注册用户
func RegisterUser(registerUserDTO dto.RegisterUserDTO) error {
	//校验User字段
	err := validateUserFieldIsEmpty(registerUserDTO)
	if err != nil {
		return err
	}

	var user entity.User
	err = smapping.FillStruct(&user, smapping.MapFields(&registerUserDTO))
	if err != nil {
		return err
	}

	//给密码加密
	hashPasswd, err := utils.HashPassword([]byte(user.Password))
	if err != nil {
		return err
	}
	user.Password = hashPasswd

	//创建User
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Login 用户登录
func Login(loginUserDTO dto.LoginUserDTO) error {
	user, err := GetUserInfo(loginUserDTO.GoGoID)
	if err != nil {
		return err
	}

	//校验密码是否正确
	_, err = utils.ComparePassword([]byte(user.Password), []byte(loginUserDTO.Password))
	if err != nil {
		return err
	}
	return nil
}

// GetUserInfo 获取用户个人信息
func GetUserInfo(goGoID uint64) (entity.User, error) {
	var user entity.User
	tx := config.DB.Where("go_go_id= ?", goGoID).Take(&user)
	return user, tx.Error
}

// UpdateUser 更新用户信息,
func UpdateUser(goGoID uint64, updateUserDTO dto.UpdateUserDTO) error {
	var user entity.User
	err := smapping.FillStruct(&user, smapping.MapFields(&updateUserDTO))
	if err != nil {
		return err
	}
	tx := config.DB.Where("go_go_id = ?", goGoID).Updates(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// UpdateUserPasswd 更新用户密码
func UpdateUserPasswd(updateUserPasswd dto.UpdateUserPasswdDTO) error {
	//获取用户信息，包括Password
	user, err := GetUserInfo(updateUserPasswd.GoGoID)
	if err != nil {
		return err
	}

	// 校验旧密码是否正确
	_, err = utils.ComparePassword([]byte(user.Password), []byte(updateUserPasswd.OldPassword))
	if err != nil {
		return err
	}

	// 加密密码
	hashPassword, err := utils.HashPassword([]byte(updateUserPasswd.NewPassword))
	if err != nil {
		return err
	}

	// 把加密的新密码赋值给用户
	user.Password = hashPassword

	tx := config.DB.Updates(user)
	if tx.Error != nil {
		return errors.New("旧密码不正确")
	}
	return nil
}

// 验证User字段是否为空
func validateUserFieldIsEmpty(registerUserDTO dto.RegisterUserDTO) error {
	if registerUserDTO.Password == "" || registerUserDTO.Nickname == "" {
		return errors.New("昵称,邮箱,密码不能为空")
	}
	return nil
}
