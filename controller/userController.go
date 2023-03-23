package controller

import (
	"github.com/PangQiMing/GoGoChat/dto"
	"github.com/PangQiMing/GoGoChat/service"
	"github.com/PangQiMing/GoGoChat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterUser 注册用户
func RegisterUser(ctx *gin.Context) {
	utils.RequestMethodPost(ctx)
	var registerUserDTO dto.RegisterUserDTO
	err := ctx.ShouldBind(&registerUserDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = service.RegisterUser(registerUserDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}

// LoginUser 用户登录
func LoginUser(ctx *gin.Context) {
	utils.RequestMethodPost(ctx)
	var loginUser dto.LoginUserDTO
	err := ctx.ShouldBind(&loginUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = service.Login(loginUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(loginUser.GoGoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
	})
}

// GetUserInfo 获取用户个人信息
func GetUserInfo(ctx *gin.Context) {
	utils.RequestMethodGet(ctx)
	goGoID := utils.VerificationToken(ctx)
	user, err := service.GetUserInfo(goGoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// LogoutUser 用户退出操作
func LogoutUser(ctx *gin.Context) {
	utils.RequestMethodPost(ctx)
	_ = utils.VerificationToken(ctx)
	ctx.Redirect(http.StatusSeeOther, "/")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "用户已退出",
	})
}

// UpdateUser 更新用户信息：头像，性别，年龄
func UpdateUser(ctx *gin.Context) {
	utils.RequestMethodPut(ctx)
	utils.RequestMethodPost(ctx)
	goGoID := utils.VerificationToken(ctx)

	var updateUserDTO dto.UpdateUserDTO
	err := ctx.ShouldBind(&updateUserDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = service.UpdateUser(goGoID, updateUserDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "用户信息已更新",
	})
}
