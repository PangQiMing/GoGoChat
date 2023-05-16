package controller

import (
	"github.com/PangQiMing/GoGoChat/service"
	"github.com/PangQiMing/GoGoChat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMessageByGoGoID(ctx *gin.Context) {
	utils.RequestMethodGet(ctx)
	goGoID := utils.VerificationToken(ctx)
	message := service.GetAllMessageByGoGoID(goGoID)
	ctx.JSON(http.StatusOK, gin.H{
		"data": message,
	})
}
