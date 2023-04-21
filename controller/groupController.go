package controller

import (
	"github.com/PangQiMing/GoGoChat/dto"
	"github.com/PangQiMing/GoGoChat/service"
	"github.com/PangQiMing/GoGoChat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateGroup 创建群组
func CreateGroup(ctx *gin.Context) {
	utils.RequestMethodPost(ctx)
	goGoID := utils.VerificationToken(ctx)
	var createGroupDTO dto.CreateGroupDTO

	err := ctx.ShouldBind(&createGroupDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = service.CreateGroup(goGoID, createGroupDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "创建群组成功",
	})
}

// UpdateGroup 更新群组信息
func UpdateGroup(ctx *gin.Context) {
	utils.RequestMethodPut(ctx)
	goGoID := utils.VerificationToken(ctx)
	var updateGroupDTO dto.UpdateGroupDTO
	err := ctx.ShouldBind(&updateGroupDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = service.UpdateGroup(goGoID, updateGroupDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "更新群组信息成功",
	})
}

// DeleteGroup 解散群组
func DeleteGroup(ctx *gin.Context) {
	utils.RequestMethodDelete(ctx)
	goGoID := utils.VerificationToken(ctx)
	var deleteGroupDTO dto.DeleteGroupDTO
	err := ctx.ShouldBind(&deleteGroupDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = service.DeleteGroup(goGoID, deleteGroupDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "解散群组成功",
	})
}

// DeleteGroupMember 群成员退出群组
func DeleteGroupMember(ctx *gin.Context) {
	utils.RequestMethodDelete(ctx)
	_ = utils.VerificationToken(ctx)
	var deleteGroupMemberDTO dto.DeleteGroupMemberDTO
	err := ctx.ShouldBind(&deleteGroupMemberDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = service.DeleteGroupMember(deleteGroupMemberDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "退出改群组成功",
	})
}

// GetGroupLists 获取群组列表
func GetGroupLists(ctx *gin.Context) {
	utils.RequestMethodGet(ctx)
	goGoID := utils.VerificationToken(ctx)
	groupLists, err := service.GetGroupLists(goGoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "获取群组列表成功",
		"data":    groupLists,
	})
}

// JoinGroupRequestList 获取入群的申请列表
func JoinGroupRequestList(ctx *gin.Context) {
	utils.RequestMethodGet(ctx)
	goGoID := utils.VerificationToken(ctx)
	result, err := service.JoinGroupRequestList(goGoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "已获取入群申请列表",
		"data":    result,
	})
}

// JoinGroup 加入群组
func JoinGroup(ctx *gin.Context) {
	utils.RequestMethodPost(ctx)
	goGoID := utils.VerificationToken(ctx)

	var joinGroup dto.JoinGroupDTO
	err := ctx.ShouldBind(&joinGroup)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	joinGroup.MemberID = goGoID
	err = service.JoinGroup(joinGroup)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "已发送入群申请",
	})
}

// AcceptJoinGroupRequest 同意入群申请
func AcceptJoinGroupRequest(ctx *gin.Context) {
	utils.RequestMethodPost(ctx)
	goGoID := utils.VerificationToken(ctx)

	var acceptJoinGroupDTO dto.AcceptJoinGroupDTO
	err := ctx.ShouldBind(&acceptJoinGroupDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	acceptJoinGroupDTO.GroupOwnerID = goGoID

	err = service.AcceptJoinGroupRequest(acceptJoinGroupDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "已同意入群申请",
	})
}

// RejectJoinGroupRequest 拒绝入群申请
func RejectJoinGroupRequest(ctx *gin.Context) {
	utils.RequestMethodPost(ctx)
	goGoID := utils.VerificationToken(ctx)

	var rejectJoinGroupDTO dto.RejectJoinGroupDTO
	err := ctx.ShouldBind(&rejectJoinGroupDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = service.RejectJoinGroupRequest(goGoID, rejectJoinGroupDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "已拒绝入群申请",
	})
}

// GetSearchGroup 查找群组
func GetSearchGroup(ctx *gin.Context) {
	utils.RequestMethodPost(ctx)
	_ = utils.VerificationToken(ctx)
	var searchGroupDTO dto.SearchGroupDTO
	err := ctx.ShouldBind(&searchGroupDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	group, err := service.GetSearchGroup(searchGroupDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    group,
		"message": "查找群组成功",
	})
}
