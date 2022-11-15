package sts

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/litsoftware/litmedia/cmd/http/proto"
	"github.com/litsoftware/litmedia/internal/controller"
	"github.com/litsoftware/litmedia/internal/ent"
	"github.com/litsoftware/litmedia/internal/pkg/proc"
	"github.com/trustmaster/goflow"
)

type Statement struct {
	controller.BaseStatement
	controller.StatementAppImpl
	controller.ClaimsImpl

	Request  *proto.StsTokenRequest
	Response *proto.StsTokenResponse

	User *ent.User
	App  *ent.App

	Claims *jwt.MapClaims
}

func Process() *goflow.Graph {
	return proc.Build(
		new(initSt),
		new(controller.JwtToken),
		new(generateToken),
	)
}
