package initialize

import (
	"crypto/rand"
	"crypto/rsa"
	"gochat_re/global"
)

func InitPrivateKeys() {
	global.PrivateKey, _ = rsa.GenerateKey(rand.Reader, 1024)
}
