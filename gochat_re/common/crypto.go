package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"go.uber.org/zap"
)

func CalculateMD5(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return string(hasher.Sum(nil))
}

func CompareMD5(data string, md5 string) bool {
	return CalculateMD5(data) == md5
}

// 将私钥传进去获取base64编码的公钥
func GetPublicKey(privateKey *rsa.PrivateKey) string {
	publickey := privateKey.PublicKey
	publicKeyBytes, _ := x509.MarshalPKIXPublicKey(&publickey)
	publickeyBlock := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	var buf bytes.Buffer
	_ = pem.Encode(&buf, &publickeyBlock)
	return buf.String()
}

// 使用私钥解析出公共对称密钥
func GetAESKey(privateKey *rsa.PrivateKey, data []byte) []byte {
	fmt.Println("data: ", string(data))
	encryptedAESKey, err := hex.DecodeString(string(data))
	if err != nil {
		zap.S().Error("err", err)
	}
	fmt.Println("Encrypted AES key:", encryptedAESKey)
	aesKey, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedAESKey)
	if err != nil {
		zap.S().Error("err", err)
	}
	fmt.Println("Decrypted AES key:", hex.EncodeToString(aesKey))
	return aesKey
}

func EncryptAES(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, len(plaintext))

	fmt.Println("你知道吗，你正在加密通信⊙⊿⊙☞")
	stream := cipher.NewCTR(block, key)
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

func DecryptAES(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext := make([]byte, len(ciphertext))

	stream := cipher.NewCTR(block, key)
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
