package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gochat_re/common"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求头中的authorization
		token := c.GetHeader("Authorization")
		fmt.Printf("token: %v\n", token)
		if token == "" {
			common.RespFail(c.Writer, common.NotLogin)
			c.Abort()
			return
		} else {
			claim, err := common.ParseToken(token, "123456")
			if err != nil {
				common.RespFail(c.Writer, common.TokenExpired)
				c.Abort()
				return
			}
			//AccessExpire, ok := (*claim)["access_expire"].(int64)
			//if !ok {
			//	common.RespFail(c.Writer, common.ToeknError)
			//	c.Abort()
			//	return
			//}
			//if AccessExpire < time.Now().Unix() {
			//	common.RespFail(c.Writer, "token")
			//	c.Abort()
			//	return
			//}
			//遍历claims

			userId, ok := (*claim)["user_id"].(float64)
			if !ok {
				common.RespFail(c.Writer, common.ToeknError)
				c.Abort()
				return
			}
			c.Set("userId", uint(userId))
			c.Next()
		}

	}
}
