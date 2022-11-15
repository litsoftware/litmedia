package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/litsoftware/litmedia/pkg/base64"
	"github.com/wenzhenxi/gorsa"
)

// GenerateRsaKeyPair generate key pair
func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, []byte, []byte, error) {
	reader := rand.Reader
	bitSize := 2048

	// privateKey
	privateKey, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var privateKeyBlock = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	privateKeyPEM := pem.EncodeToMemory(privateKeyBlock)

	// public key
	publicKey := &privateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	var publicKeyBlock = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicKeyPEM := pem.EncodeToMemory(publicKeyBlock)

	return privateKey, publicKey, privateKeyPEM, publicKeyPEM, nil
}

// PrivateEncrypt 私钥加密
func PrivateEncrypt(data string, privt string) (string, error) {
	ciphertext, err := gorsa.PriKeyEncrypt(data, privt)
	if err != nil {
		return "", nil
	}

	return ciphertext, nil
}

// PublicDecrypt 公钥解密
func PublicDecrypt(ciphertext string, pub string) (string, error) {
	data, err := gorsa.PublicDecrypt(ciphertext, pub)
	if err != nil {
		return "", err
	}

	return data, nil
}

// PrivateSign 签名
func PrivateSign(data string, priv *rsa.PrivateKey) (string, error) {
	hashFunc := crypto.SHA256
	h := hashFunc.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	signData, err := rsa.SignPKCS1v15(rand.Reader, priv, hashFunc, hashed)
	if err != nil {
		return "", err
	}

	return base64.Encode(string(signData)), nil
}

func PublicVerify(data string, sign string, pub *rsa.PublicKey) error {
	signData, err := base64.Decode(sign)
	if err != nil {
		return err
	}

	hashFunc := crypto.SHA256
	h := hashFunc.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed, []byte(signData))
	if err != nil {
		return err
	}

	return nil
}

// PublicEncrypt 公钥加密
func PublicEncrypt(data string, publicKeyPemStr string) (string, error) {
	publicKey, err := ParseRsaPublicKeyFromPemStr(publicKeyPemStr)
	if err != nil {
		fmt.Println("ParseRsaPublicKeyFromPemStr", 1111)
		return "", err
	}

	body, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(data))
	if err != nil {
		fmt.Println("EncryptPKCS1v15", 22222)
		return "", err
	}
	return base64.Encode(string(body)), nil
}

// PrivateDecrypt 私钥解密
func PrivateDecrypt(ciphertext string, privateKeyPem string) (string, error) {
	privateKey, err := ParseRsaPrivateKeyFromPemStr(privateKeyPem)
	if err != nil {
		return "", err
	}

	resultTemp, err := base64.NewBase64().Decode(ciphertext)
	if err != nil {
		return "", err
	}

	body, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, resultTemp)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ParseRsaPrivateKeyFromPemBase64Str(privateKeyPEM string) (*rsa.PrivateKey, error) {
	privBytes, err := base64.NewBase64().Decode(privateKeyPEM)
	if err != nil {
		return nil, err
	}

	return ParseRsaPrivateKeyFromPemStrBytes(privBytes)
}

func ParseRsaPrivateKeyFromPemStr(privateKeyPEM string) (*rsa.PrivateKey, error) {
	return ParseRsaPrivateKeyFromPemStrBytes([]byte(privateKeyPEM))
}

func ParseRsaPrivateKeyFromPemStrBytes(privBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privBytes)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	var pk *rsa.PrivateKey
	var ok bool

	pk8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		pk, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	}

	if pk == nil {
		pk, ok = pk8.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("parse PKCS8 key error")
		}
	}

	return pk, nil
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}

	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
}

func ParseRsaPublicKeyFromPemBase64Str(publicKeyPEM string) (*rsa.PublicKey, error) {
	publicKeyBytes, err := base64.NewBase64().Decode(publicKeyPEM)
	if err != nil {
		return nil, err
	}
	return ParseRsaPublicKeyFromPemStrBytes(publicKeyBytes)
}

func ParseRsaPublicKeyFromPemStr(publicKeyPEM string) (*rsa.PublicKey, error) {
	return ParseRsaPublicKeyFromPemStrBytes([]byte(publicKeyPEM))
}

func ParseRsaPublicKeyFromPemStrBytes(publicKeyBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pub.(*rsa.PublicKey), nil
}

func ParseCertFromPemBase64Str(certPEM string) (*x509.Certificate, error) {
	fmt.Printf("certPEM \n")
	fmt.Printf(certPEM)

	certBytes, err := base64.NewBase64().Decode(certPEM)
	if err != nil {
		return nil, err
	}

	return ParseCertFromPemStrBytes(certBytes)
}

func ParseCertFromPemStr(certPEM string) (*x509.Certificate, error) {
	return ParseCertFromPemStrBytes([]byte(certPEM))
}

func ParseCertFromPemStrBytes(certBytes []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(certBytes)
	if block == nil {
		return nil, errors.New("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse certificate: " + err.Error())
	}

	return cert, nil
}
