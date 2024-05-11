package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gochat_re/common"
	"gochat_re/dao"
	"gochat_re/models"
	"time"
)

func UserList(ctx *gin.Context) {
	list, err := dao.GetUserList()
	if err != nil {
		common.RespFail(ctx.Writer, common.GetUserListFail)
		return
	}
	common.RespOKList(ctx.Writer, list, len(list))
	return
}

func GetUserInfo(ctx *gin.Context) {
	userId := ctx.Value("userId").(uint)
	user, err := dao.GetUserByUserid(userId)
	if err != nil {
		zap.S().Error(err)
		common.RespFail(ctx.Writer, common.GetUserInfoFail)
		return
	}
	resUser := models.GetUserResponse(*user)
	common.RespOK(ctx.Writer, resUser, "")
	return
}

func Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		zap.S().Error(common.DataFormatError)
		common.RespFail(ctx.Writer, common.DataFormatError)
		return
	}
	getUser, err := dao.GetUserByNameAndPassword(user)
	if err != nil {
		zap.S().Error(common.UserNameOrPasswordError)
		common.RespFail(ctx.Writer, common.UserNameOrPasswordError)
		return
	}
	claims := make(map[string]interface{})
	claims["user_id"] = getUser.ID
	token, err := common.GenToken(claims)
	common.RespOK(ctx.Writer, token, common.LoginSuccess)
	return
}

func Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		zap.S().Error(common.DataFormatError)
		common.RespFail(ctx.Writer, common.DataFormatError)
		return
	}
	if _, err := dao.GetUserByName(user.Name); err == nil {
		zap.S().Error(common.UserNameExist)
		common.RespFail(ctx.Writer, common.UserNameExist)
		return
	}
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	createUser, err := dao.CreateUser(user)
	if err != nil {
		zap.S().Error(common.RegisterFail)
		common.RespFail(ctx.Writer, common.RegisterFail)
		return
	}
	claims := make(map[string]interface{})
	claims["user_id"] = createUser.ID
	token, err := common.GenToken(claims)
	if err != nil {
		zap.S().Error("token generate fail")
		common.RespFail(ctx.Writer, "token generate fail")
		return
	}
	common.RespOK(ctx.Writer, token, common.RegisterSuccess)
	return
}

func UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		zap.S().Error(common.DataFormatError)
		common.RespFail(ctx.Writer, common.DataFormatError)
		return
	}
	if _, err := dao.UpdateUser(user); err != nil {
		zap.S().Error(common.UserNotExist)
		common.RespFail(ctx.Writer, common.UserNotExist)
		return
	}
	common.RespOK(ctx.Writer, user, common.UpdateUserSuccess)
	return
}

func DeleteUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		zap.S().Error(common.DataFormatError)
		common.RespFail(ctx.Writer, common.DataFormatError)
		return
	}
	common.RespOK(ctx.Writer, nil, common.UpdateUserSuccess)
	return
}

func SendMessage(ctx *gin.Context) {
	dao.Chat(ctx)

}
