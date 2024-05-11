package common

import (
	"fmt"
	"gochat_re/common"
	"testing"
	"time"
)

func TestBuildTokens(t *testing.T) {
	opt := common.TokenOptions{
		AccessSecret: "your_secret_key",
		AccessExpire: time.Now().Add(time.Hour * 1).Unix(),
		Fields: map[string]interface{}{
			"user_id": 123,
		},
	}

	token, err := common.BuildTokens(opt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %v\n", token)

	if token.AccessToken == "" {
		t.Error("Expected access token to be set, got empty string")
	}

	if token.AccessExpire <= time.Now().Unix() {
		t.Error("Expected access token to expire in the future, got past or present time")
	}
}

func TestParseToken(t *testing.T) {
	opt := common.TokenOptions{
		AccessSecret: "your_secret_key",
		AccessExpire: time.Now().Add(time.Hour * 1).Unix(),
		Fields: map[string]interface{}{
			"user_id": 123,
		},
	}
	token, err := common.BuildTokens(opt)
	if err != nil {
		t.Fatal(err)
	}
	claims, err := common.ParseToken(token.AccessToken, opt.AccessSecret)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(claims)

}

func TestGenToken(t *testing.T) {
	claims := make(map[string]interface{})
	claims["user_id"] = 123

	token, err := common.GenToken(claims)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)

}
