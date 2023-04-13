package controller

import (
	"github.com/PangQiMing/GoGoChat/dto"
	"github.com/PangQiMing/GoGoChat/service"
	"github.com/PangQiMing/GoGoChat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetFriendList 获取好友列表
func GetFriendList(ctx *gin.Context) {
	utils.RequestMethodGet(ctx)
	goGoID := utils.VerificationToken(ctx)

	friendList, err := service.GetFriendList(goGoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": friendList,
	})
}

// AddFriend 添加好友，向好友发送好友请求
func AddFriend(ctx *gin.Context) {
	utils.RequestMethodPost(ctx)
	goGoID := utils.VerificationToken(ctx)

	var addFriendDTO dto.AddFriendDTO
	err := ctx.ShouldBind(&addFriendDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	addFriendDTO.GoGoID = goGoID
	err = service.AddFriend(addFriendDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "已发送好友请求",
	})
}

// GetFriendRequestList 获取好友请求列表
func GetFriendRequestList(ctx *gin.Context) {
	utils.RequestMethodGet(ctx)
	goGoID := utils.VerificationToken(ctx)

	friendRequestList, err := service.GetFriendRequestList(goGoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "好友申请列表",
		"data":    friendRequestList,
	})
}

// AcceptFriendRequest 同意好友申请
func AcceptFriendRequest(ctx *gin.Context) {
	utils.RequestMethodPost(ctx)
	goGoID := utils.VerificationToken(ctx)

	var acceptFriendRequest dto.AcceptFriendDTO
	err := ctx.ShouldBind(&acceptFriendRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	acceptFriendRequest.GoGoID = goGoID
	err = service.AcceptFriendRequest(acceptFriendRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "同意好友申请",
	})
}

// RejectFriendRequest 拒绝好友申请
func RejectFriendRequest(ctx *gin.Context) {
	utils.RequestMethodPost(ctx)
	goGoID := utils.VerificationToken(ctx)

	var rejectFriendRequest dto.RejectFriendDTO
	err := ctx.ShouldBind(&rejectFriendRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	rejectFriendRequest.GoGoID = goGoID
	err = service.RejectFriendRequest(rejectFriendRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "拒绝好友申请",
	})
}

// DeleteFriend 删除好友
func DeleteFriend(ctx *gin.Context) {
	utils.RequestMethodDelete(ctx)
	goGoID := utils.VerificationToken(ctx)

	var deleteFriendDTO dto.DeleteFriendDTO
	err := ctx.ShouldBind(&deleteFriendDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	deleteFriendDTO.GoGoID = goGoID
	err = service.DeleteFriend(deleteFriendDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "删除好友成功",
	})
}

// GetSearchFriend 获取好友信息，判断好友是否存在
func GetSearchFriend(ctx *gin.Context) {
	utils.RequestMethodGet(ctx)
	_ = utils.VerificationToken(ctx)

	var searchFriendDTO dto.SearchFriendDTO
	err := ctx.ShouldBind(&searchFriendDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	friend, err := service.SearchFriendByFriendID(searchFriendDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": friend,
	})
}
