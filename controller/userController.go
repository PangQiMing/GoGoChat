package controller

import (
	"github.com/PangQiMing/GoGoChat/dto"
	"github.com/PangQiMing/GoGoChat/service"
	"github.com/PangQiMing/GoGoChat/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
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

func UpdatePassword(ctx *gin.Context) {
	utils.RequestMethodPut(ctx)
	goGoID := utils.VerificationToken(ctx)
	var updatePasswordDTO dto.UpdateUserPasswdDTO
	err := ctx.ShouldBind(&updatePasswordDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatePasswordDTO.GoGoID = goGoID
	err = service.UpdateUserPasswd(updatePasswordDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "修改密码成功",
	})
}

func UpdateAvatar(ctx *gin.Context) {
	utils.RequestMethodPost(ctx)
	goGoID := utils.VerificationToken(ctx)
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "更新头像失败",
		})
		return
	}

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "更新头像失败",
		})
		return
	}

	//为图片文件生成uuid唯一标识
	fileStr := uuid.New().String()

	basePath := "static/"
	imgPath := "images/" + fileStr + ".jpg"
	savePath := basePath + imgPath
	err = ioutil.WriteFile(savePath, fileData, 0666)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "更新头像失败",
		})
		return
	}
	err = service.UpdateAvatar(goGoID, imgPath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "更新头像失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "更新头像成功",
	})
}
