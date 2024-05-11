package router

import (
	"github.com/gin-gonic/gin"
	"gochat_re/middleware"
	"gochat_re/server"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware)
	v1 := router.Group("v1")
	user := v1.Group("user")
	{
		user.POST("register", server.Register)
		user.POST("login", server.Login)
	}
	auth := v1.Group("user")
	auth.Use(middleware.JWT())
	{
		auth.GET("info", server.GetUserInfo)
		auth.GET("friends", server.FriendList)
		auth.GET("publckey", server.GetPublicKey)
		auth.GET("aesKey", server.UpdateAESKey)
	}

	chat := v1.Group("chat")
	//chat.Use(middleware.JWT())
	{
		chat.GET("ws", server.SendMessage)
	}
	return router

}
