package controller

import (
	"github.com/PangQiMing/GoGoChat/entity"
	"github.com/PangQiMing/GoGoChat/service"
	"github.com/PangQiMing/GoGoChat/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func WSController(ctx *gin.Context, hub *entity.Hub) {
	utils.RequestMethodGet(ctx)
	//goGoID := utils.VerificationToken(ctx)
	goGoID := uuid.New().ID()
	service.ServeWS(uint64(goGoID), hub, ctx.Writer, ctx.Request)
}
