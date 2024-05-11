package main

import (
	"gochat_re/initialize"
	"gochat_re/router"
)

func init() {
	initialize.InitDB()
	initialize.InitLogger()
	initialize.InitPrivateKeys()
}

func main() {
	rt := router.Router()
	err := rt.Run(":8080")
	if err != nil {
		return
	}

}
