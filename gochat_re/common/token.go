package common

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"time"
)

type (
	TokenOptions struct {
		AccessSecret string
		AccessExpire int64
		Fields       map[string]interface{}
	}

	Token struct {
		AccessToken  string `json:"access_token"`
		AccessExpire int64  `json:"access_expire"`
	}
)

func ParseToken(tokenString string, secretKey string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrInvalidKeyType
	}

	return &claims, nil
}

func BuildTokens(opt TokenOptions) (Token, error) {
	var token Token
	now := time.Now().Add(-time.Minute).Unix()
	accessToken, err := genToken(now, opt.AccessSecret, opt.Fields, opt.AccessExpire)
	if err != nil {
		return token, err
	}
	token.AccessToken = accessToken
	token.AccessExpire = now + opt.AccessExpire

	return token, nil
}

func genToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

func GenToken(Field map[string]interface{}) (Token, error) {
	err := godotenv.Load("config/.env")
	if err != nil {
		zap.S().Error("读取配置文件失败")
	}
	secretKey := os.Getenv("SECRET_KEY")
	secretExpired := os.Getenv("SECRET_EXPIRED")
	secretExpiredInt, err := time.ParseDuration(secretExpired)
	fmt.Println(secretKey)
	fmt.Printf(secretExpired)
	if err != nil {
		zap.S().Error("转换过期时间失败")
		return Token{}, err
	}
	opt := TokenOptions{
		AccessSecret: secretKey,
		AccessExpire: time.Now().Add(secretExpiredInt).Unix(),
		Fields:       Field,
	}
	return BuildTokens(opt)
}
