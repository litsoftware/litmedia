package controller

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/litsoftware/litmedia/internal/common"
	"github.com/litsoftware/litmedia/internal/errors"
	"github.com/litsoftware/litmedia/internal/g"
	"github.com/litsoftware/litmedia/pkg/base64"
	"github.com/litsoftware/litmedia/pkg/jwth"
	"github.com/litsoftware/litmedia/pkg/rsa"
)

type JwtToken struct {
	BaseProc
}

func (c *JwtToken) Process() {
	r := <-c.In
	st := r.Req.(StatementInterface)
	if st.Error() != nil {
		c.Next(r)
		return
	}

	var token string
	if r.GinCtx.GetHeader("authorization") != "" {
		token = r.GinCtx.GetHeader("authorization")
	}
	if r.GinCtx.Query("t") != "" {
		token = r.GinCtx.Query("t")
	}

	if token == "" {
		st.SetError(errors.ErrorUnauthorized("missing authorization header"))
		c.Next(r)
		return
	}

	payload, err := jwth.GetPayload(token)
	if err != nil {
		g.App.Logger.Errorf("GetPayload err: %s", err.Error())
		st.SetError(errors.ErrorUnauthorized(""))
		c.Next(r)
		return
	}

	if (*payload)["aki"] == nil {
		g.App.Logger.Errorf("aki is nil")
		st.SetError(errors.ErrorUnauthorized(""))
		c.Next(r)
		return
	}

	appEnt, err := common.GetAppByAppId(context.Background(), (*payload)["aki"].(string))
	if err != nil {
		g.App.Logger.Errorf("GetAppByAppId err: %s", err.Error())
		st.SetError(errors.ErrorUnauthorized(""))
		c.Next(r)
		return
	}

	priv, err1 := jwt.ParseRSAPrivateKeyFromPEM([]byte(appEnt.EncryptedAppPrivateKey))
	pub, err2 := jwt.ParseRSAPublicKeyFromPEM([]byte(appEnt.EncryptedAppPublicKey))
	if err1 != nil || err2 != nil {
		g.App.Logger.Errorf("ParseRSAPrivateKeyFromPEM err: %s", err.Error())
		st.SetError(errors.ErrorUnauthorized(""))
		c.Next(r)
		return
	}

	aud, ok := (*payload)["aud"]
	if ok && aud == "sts" {
		g.App.Logger.Infof("token is sts token, verify use app public key")
		jwth.SetPublicKey(pub)
	} else {
		s, err := base64.Decode(appEnt.EncryptedOperatorRsaPublicKey)
		if err != nil {
			g.App.Logger.Errorf("Decode err: %s", err.Error())
			st.SetError(errors.ErrorUnauthorized(""))
			c.Next(r)
			return
		}

		g.App.Logger.Infof("token is app token, verify use user public key ", s)
		userPub, err := rsa.ParseRsaPublicKeyFromPemStrBytes([]byte(s))
		if err != nil {
			g.App.Logger.Errorf("ParseRSAPublicKeyFromPEM err: %s", err.Error())
			st.SetError(errors.ErrorUnauthorized(""))
			c.Next(r)
			return
		}
		g.App.Logger.Infof("token is normal token ")
		jwth.SetPublicKey(userPub)
	}

	jwth.SetPrivateKey(priv)
	claims, err := jwth.ParseToken(token)
	if err != nil {
		g.App.Logger.Errorf("ParseToken err: %s", err.Error())
		st.SetError(err)
		c.Next(r)
		return
	}

	err = st.(AppParamInterface).SetApp(st, appEnt)
	if err != nil {
		g.App.Logger.Errorf("setApp err: %s", err.Error())
		st.SetError(err)
		c.Next(r)
		return
	}

	err = st.(ClaimsParamInterface).SetClaims(st, &claims)
	if err != nil {
		g.App.Logger.Errorf("setClaims err: %s", err.Error())
		c.Next(r)
		return
	}

	c.Next(r)
	return
}
