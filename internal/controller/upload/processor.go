package upload

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/litsoftware/litmedia/cmd/http/proto"
	"github.com/litsoftware/litmedia/internal/controller"
	"github.com/litsoftware/litmedia/internal/ent"
	"github.com/litsoftware/litmedia/internal/pkg/proc"
	"github.com/trustmaster/goflow"
	"mime/multipart"
)

type Statement struct {
	controller.BaseStatement
	controller.StatementAppImpl
	controller.ClaimsImpl

	Request  *proto.FileUploadRequest
	Response *proto.FileUploadResponse

	User  *ent.User
	App   *ent.App
	Media *ent.Media

	Claims *jwt.MapClaims

	file       multipart.File
	fileHeader *multipart.FileHeader
}

func Process() *goflow.Graph {
	return proc.Build(
		new(initSt),
		new(controller.JwtToken),
		new(fileChecker),
		new(saveMedia),
		new(uploadMediaToOss),
	)
}
