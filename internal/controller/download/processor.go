package download

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/litsoftware/litmedia/cmd/http/proto"
	"github.com/litsoftware/litmedia/internal/controller"
	"github.com/litsoftware/litmedia/internal/ent"
	"github.com/litsoftware/litmedia/internal/pkg/proc"
	"github.com/trustmaster/goflow"
	"io"
)

type Statement struct {
	controller.BaseStatement
	controller.StatementAppImpl
	controller.ClaimsImpl

	Request  *proto.FileDownloadRequest
	Response *proto.FileDownloadResponse

	User  *ent.User
	App   *ent.App
	Media *ent.Media

	Claims *jwt.MapClaims
	File   io.ReadCloser
}

func Process() *goflow.Graph {
	return proc.Build(
		new(initSt),
		new(controller.JwtToken),
		new(queryMedia),
		new(downloadFromOss),
	)
}
