package sts

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/litsoftware/litmedia/internal/controller"
	"github.com/litsoftware/litmedia/internal/g"
	"github.com/litsoftware/litmedia/pkg/jwth"
	"time"
)

type generateToken struct {
	controller.BaseProc
}

func (c *generateToken) Process() {
	r := <-c.In
	st := r.Req.(*Statement)
	if st.Err != nil {
		return
	}

	//if (*st.Claims)["aud"].(string) == "sts" {
	//	st.Err = errors.ErrorWrongTokenType("token is sts")
	//	return
	//}

	priv, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(st.App.EncryptedAppPrivateKey))
	pub, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(st.App.EncryptedAppPublicKey))
	jwth.SetPublicKey(pub)
	jwth.SetPrivateKey(priv)

	if st.Request.Ttl < 1 {
		st.Request.Ttl = 3600
	}

	token, err := jwth.CreateToken(&jwt.MapClaims{
		"iat": time.Now().Unix(),                                    // Token颁发时间
		"nbf": time.Now().Unix(),                                    // Token生效时间
		"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(), // Token过期时间
		"iss": "rpa media service sts token from server",

		"aid": st.App.ID,            // 应用ID
		"aki": st.App.AppID,         // 应用ID
		"aud": "sts",                // token 类型
		"uid": st.Request.UserId,    // 客户系统的用户ID
		"cid": st.Request.CompanyId, // 客户系统的公司ID
		"res": st.Request.Resource,
	})

	if err != nil {
		g.App.Logger.Errorf("generate token error: %s", err.Error())
		st.Err = err
		return
	}

	st.Response.StsToken = token
	return
}
