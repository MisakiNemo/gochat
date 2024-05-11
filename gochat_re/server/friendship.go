package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gochat_re/common"
	"gochat_re/dao"
	"gochat_re/models"
	"strconv"
)

func FriendList(ctx *gin.Context) {
	userId := ctx.Value("userId").(uint)
	if userId == 0 {
		zap.S().Error("获取用户id失败")
		common.RespFail(ctx.Writer, "获取用户id失败")
		return
	}
	friendShipList, err := dao.GetFriendList(userId)
	if err != nil {
		zap.S().Error("获取好友列表错误")
		common.RespFail(ctx.Writer, "获取好友列表错误")
		return
	}
	var friendList []models.UserResponse
	for _, v := range friendShipList {
		user, err := dao.GetUserByUserid(v.FriendID)
		if err != nil {
			zap.S().Error("获取用户信息失败")
			common.RespFail(ctx.Writer, "获取用户信息失败")
			return
		}
		friendList = append(friendList, models.GetUserResponse(*user))
	}
	fmt.Printf(strconv.Itoa(len(friendList)))
	common.RespOKList(ctx.Writer, friendList, len(friendList))
	return
}
