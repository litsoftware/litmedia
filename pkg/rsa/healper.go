package rsa

import (
	"bytes"
	"fmt"
	"strings"
)

// IsValidateRsaPublicKeyStr 验证RSA公钥字符串是否合法
func IsValidateRsaPublicKeyStr(s string) (bool, error) {
	_, err := ParseRsaPublicKeyFromPemStr(s)
	if err != nil {
		return false, err
	}

	return true, nil
}

// IsValidateRsaPrivateKeyStr 验证RSA死私钥字符串是否合法
func IsValidateRsaPrivateKeyStr(s string) (bool, error) {
	_, err := ParseRsaPrivateKeyFromPemStr(s)
	if err != nil {
		return false, err
	}

	return true, nil
}

// IsValidateRsaKeyPair 验证一对rsa密钥是否是合法的
func IsValidateRsaKeyPair(publicKey, privateKey string) (bool, error) {
	pub, err := ParseRsaPublicKeyFromPemStr(publicKey)
	if err != nil {
		return false, err
	}

	priv, err := ParseRsaPrivateKeyFromPemStr(privateKey)
	if err != nil {
		return false, err
	}

	data := "a"
	sign, err := PrivateSign(data, priv)
	if err != nil {
		return false, err
	}

	err = PublicVerify(data, sign, pub)
	if err != nil {
		return false, err
	}

	return true, nil
}

func PublicKeyFormat(s string) string {
	headerLine := "-----BEGIN PUBLIC KEY-----"
	footerLine := "-----END PUBLIC KEY-----"

	return FormatPemStr(s, headerLine, footerLine)
}

func PrivateKeyFormat(s string) string {
	headerLine := "-----BEGIN RSA PRIVATE KEY-----"
	footerLine := "-----END RSA PRIVATE KEY-----"

	return FormatPemStr(s, headerLine, footerLine)
}

func FormatPemStr(s, headerLine, footerLine string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	s = strings.ReplaceAll(s, headerLine, "")
	s = strings.ReplaceAll(s, footerLine, "")

	sArr := SplitSubN(s, 64)
	s = strings.Join(sArr, "\n")

	if !strings.Contains(s, headerLine) {
		s = fmt.Sprintf("%s\n%s", headerLine, s)
	}

	if !strings.Contains(s, footerLine) {
		s = fmt.Sprintf("%s\n%s", s, footerLine)
	}

	return s
}

func SplitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

func FormatToAlipay(pem string) string {
	pem = strings.ReplaceAll(pem, "-----BEGIN PUBLIC KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----END PUBLIC KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----BEGIN RSA PRIVATE KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----END RSA PRIVATE KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----BEGIN CERTIFICATE-----", "")
	pem = strings.ReplaceAll(pem, "-----END CERTIFICATE-----", "")
	pem = strings.ReplaceAll(pem, "-----BEGIN PRIVATE KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----END PRIVATE KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----BEGIN ENCRYPTED PRIVATE KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----END ENCRYPTED PRIVATE KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----BEGIN ENCODED RSA PRIVATE KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----END ENCODED RSA PRIVATE KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----BEGIN DSA PRIVATE KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----END DSA PRIVATE KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----BEGIN EC PRIVATE KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----END EC PRIVATE KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----BEGIN PUBLIC KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----END PUBLIC KEY-----", "")
	pem = strings.ReplaceAll(pem, "-----BEGIN CERTIFICATE REQUEST-----", "")
	pem = strings.ReplaceAll(pem, "-----END CERTIFICATE REQUEST-----", "")
	pem = strings.ReplaceAll(pem, "-----BEGIN NEW CERTIFICATE REQUEST-----", "")
	pem = strings.ReplaceAll(pem, "-----END NEW CERTIFICATE REQUEST-----", "")
	pem = strings.TrimLeft(pem, "\n")
	pem = strings.TrimLeft(pem, " ")
	pem = strings.TrimRight(pem, "\n")
	pem = strings.TrimRight(pem, " ")
	pem = strings.ReplaceAll(pem, "\n", "")
	return pem
}
