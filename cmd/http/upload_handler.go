package http

import (
	"github.com/gin-gonic/gin"
	"github.com/litsoftware/litmedia/cmd/http/proto"
	"github.com/litsoftware/litmedia/internal/controller"
	"github.com/litsoftware/litmedia/internal/controller/upload"
)

func FileUpload(c *gin.Context) {
	rsp := &proto.FileUploadResponse{}
	err := controller.RunProc2(c, &upload.Statement{}, &proto.FileUploadRequest{}, rsp, upload.Process)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, rsp)
	return
}
