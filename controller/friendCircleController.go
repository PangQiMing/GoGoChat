package controller

import (
	"github.com/PangQiMing/GoGoChat/dto"
	"github.com/PangQiMing/GoGoChat/service"
	"github.com/PangQiMing/GoGoChat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFriendCircleAll(ctx *gin.Context) {
	utils.RequestMethodDelete(ctx)
	goGoID := utils.VerificationToken(ctx)
	all, err := service.GetFriendCircleAll(goGoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": all,
	})

}

func CreateFriendCircle(ctx *gin.Context) {
	utils.RequestMethodDelete(ctx)
	goGoID := utils.VerificationToken(ctx)

	var createFriendCircleDTO dto.CircleCreateDTO
	err := ctx.ShouldBind(&createFriendCircleDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	createFriendCircleDTO.GoGoID = goGoID

	err = service.CreateFriendCircle(createFriendCircleDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "发布成功",
	})
}

func DeleteFriendCircle(ctx *gin.Context) {
	utils.RequestMethodDelete(ctx)
	goGoID := utils.VerificationToken(ctx)

	var deleteFriendCircle dto.DeleteFriendCircleDTO
	err := ctx.ShouldBind(&deleteFriendCircle)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = service.DeleteFriendCircle(goGoID, deleteFriendCircle)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "删除动态失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "删除动态成功",
	})
}
