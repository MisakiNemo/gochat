package server

import (
	"github.com/gin-gonic/gin"
	"gochat_re/common"
	"gochat_re/dao"
	"gochat_re/global"
)

func GetPublicKey(ctx *gin.Context) {
	common.RespOK(ctx.Writer, common.GetPublicKey(global.PrivateKey), "")
	return
}

func UpdateAESKey(ctx *gin.Context) {
	userId := ctx.Value("userId").(uint)
	aesKey := []byte(ctx.Query("aeskey"))
	dao.BindAESKey(userId, aesKey)
	common.RespOK(ctx.Writer, "", "")
	return
}
