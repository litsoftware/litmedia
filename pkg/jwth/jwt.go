package jwth

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/litsoftware/litmedia/internal/g"
	"strings"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func SetPrivateKey(priv *rsa.PrivateKey) {
	privateKey = priv
}

func SetPublicKey(pub *rsa.PublicKey) {
	publicKey = pub
}

func CreateToken(claim *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	return token.SignedString(privateKey)
}

func GetPayload(tokenStr string) (*jwt.MapClaims, error) {
	arr := strings.Split(tokenStr, ".")
	if len(arr) < 3 {
		return nil, errors.New("token is invalid 1")
	}

	b, err := jwt.DecodeSegment(arr[1])
	if err != nil {
		return nil, errors.New("token is invalid 2")
	}

	claims := &jwt.MapClaims{}
	err = json.Unmarshal(b, claims)
	return claims, err
}

func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			g.App.Logger.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}

		if token.Method.(*jwt.SigningMethodRSA).Alg() != jwt.SigningMethodRS256.Alg() {
			return nil, jwt.ErrSignatureInvalid
		}

		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
