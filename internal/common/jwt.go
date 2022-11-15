package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/litsoftware/litmedia/internal/ent"
	"github.com/litsoftware/litmedia/pkg/jwth"
	"time"
)

func GenerateStsTokenByApp(app *ent.App) (string, error) {
	priv, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(app.EncryptedAppPrivateKey))
	pub, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(app.EncryptedAppPublicKey))
	jwth.SetPublicKey(pub)
	jwth.SetPrivateKey(priv)

	return jwth.CreateToken(&jwt.MapClaims{
		"iat": time.Now().Unix(),                                      // Token颁发时间
		"nbf": time.Now().Unix(),                                      // Token生效时间
		"exp": time.Now().Add(time.Hour * time.Duration(1000)).Unix(), // Token过期时间
		"iss": "rpa media service",

		"aid": app.ID,    // 应用ID
		"aki": app.AppID, // 应用ID
		"aud": "sts",     // token 类型
	})
}
