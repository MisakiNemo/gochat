package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"gochat_re/common"

	"testing"
)

func TestCalculateMD5(t *testing.T) {
	h := "hello"
	md5h := common.CalculateMD5(h)
	fmt.Printf("md5 hash: %v\n", md5h)
	if md5h == "" {
		t.Error("Expected md5 hash to be set, got nil")
	}
	flag := common.CompareMD5(h, md5h)
	if flag == false {
		t.Error("Expected md5 hash to be equal, got different")
	}
}

func TestGetPublicKey(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	publicKey := common.GetPublicKey(privateKey)
	fmt.Printf("public key: %v\n", string(publicKey))
	if publicKey == "" {
		t.Error("Expected public key to be set, got nil")
	}
}

func TestGetAESKey(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	publicKey := common.GetPublicKey(privateKey)
	fmt.Println(string(publicKey))
	pkblock, _ := pem.Decode([]byte(publicKey))
	fmt.Println("pkblock: ", pkblock)
	if pkblock == nil {
		t.Error("Expected public key to be decoded, got nil")
	}
	publicInterface, _ := x509.ParsePKIXPublicKey(pkblock.Bytes)
	publicKeyt := publicInterface.(*rsa.PublicKey)
	var AESKey = make([]byte, 16)
	_, _ = rand.Read(AESKey)
	fmt.Println("AESKey: ", hex.EncodeToString(AESKey))
	encryptedAESKey, _ := rsa.EncryptPKCS1v15(rand.Reader, publicKeyt, AESKey)
	fmt.Println("encrypted aes key: ", encryptedAESKey)
	fmt.Println(string(encryptedAESKey))
	aesKey := common.GetAESKey(privateKey, []byte(hex.EncodeToString(encryptedAESKey)))
	fmt.Println("aes key: ", hex.EncodeToString(aesKey))
}

func TestGetAESKey2(t *testing.T) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	publicKey := common.GetPublicKey(privateKey)
	fmt.Println(string(publicKey))
	pkblock, _ := pem.Decode([]byte(publicKey))
	fmt.Println("pkblock: ", pkblock)
	if pkblock == nil {
		t.Error("Expected public key to be decoded, got nil")
	}
	publicInterface, _ := x509.ParsePKIXPublicKey(pkblock.Bytes)
	publicKeyt := publicInterface.(*rsa.PublicKey)
	var AESKey = make([]byte, 16)
	_, _ = rand.Read(AESKey)
	fmt.Println("AESKey: ", len(hex.EncodeToString(AESKey)))
	encryptedAESKey, _ := rsa.EncryptPKCS1v15(rand.Reader, publicKeyt, AESKey)
	fmt.Println("encrypted aes key: ", encryptedAESKey)
	fmt.Println(string(encryptedAESKey))
	aesKey := common.GetAESKey(privateKey, encryptedAESKey)
	fmt.Println("aes key: ", hex.EncodeToString(aesKey))
}

func TestEncryptAES(t *testing.T) {
	key := []byte("5ddfbce7ba906a0cdd7ba5e738362f43")
	plaintext := []byte("hello")
	ciphertext, err := common.EncryptAES(key, plaintext)
	if err != nil {
		t.Error("Expected ciphertext to be encrypted, got error")
	}
	fmt.Printf("ciphertext: %v\n", ciphertext)
	res, err := common.DecryptAES(key, ciphertext)
	if err != nil {
		t.Error("Expected plaintext to be decrypted, got error")
	}
	fmt.Printf("res: %v\n", res)
}
