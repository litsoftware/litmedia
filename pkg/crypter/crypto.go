package crypter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/litsoftware/litmedia/internal/pkg/config"
	"strings"
)

//加密
func Encrypt(value string) (string, error) {
	return AesEncrypt(value, getKey())
}

//解密
func Decrypt(value string) (string, error) {
	return AesDecrypt(value, getKey())
}

func checkMAC(message, msgMac, secret string) bool {
	expectedMAC := computeHmacSha256(message, secret)
	fmt.Println(expectedMAC, msgMac)
	return hmac.Equal([]byte(expectedMAC), []byte(msgMac))
}

func computeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

//处理密钥
func getKey() string {
	appKey := config.GetString("app_key")
	if strings.HasPrefix(appKey, "base64:") {
		split := appKey[7:]
		if key, err := base64.StdEncoding.DecodeString(split); err == nil {
			return string(key)
		}
		return split
	}
	return appKey
}
