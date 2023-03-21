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
