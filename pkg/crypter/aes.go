package crypter

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"github.com/litsoftware/litmedia/pkg/base64"
)

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// goAes 加密
type goAes struct {
	Key []byte
}

func newGoAes(key []byte) *goAes {
	return &goAes{Key: key}
}

// Encrypt 加密数据
func (a *goAes) Encrypt(origData []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, a.Key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// Decrypt 解密数据
func (a *goAes) Decrypt(crypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, a.Key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pKCS7UnPadding(origData)
	return origData, nil
}

func AesEncrypt(text string, key string) (string, error) {
	if text == "" || len([]rune(text)) < 1 {
		return text, nil
	}

	a := newGoAes([]byte(key))
	t, err := a.Encrypt([]byte(text))
	return base64.NewBase64().Encode(t), err
}

func AesDecrypt(ciphertext string, key string) (string, error) {
	if ciphertext == "" || len([]rune(ciphertext)) < 1 {
		return ciphertext, nil
	}

	a := newGoAes([]byte(key))
	cb, err := base64.NewBase64().Decode(ciphertext)
	t, err := a.Decrypt(cb)
	return string(t), err
}
