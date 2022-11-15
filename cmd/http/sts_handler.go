package http

import (
	"github.com/gin-gonic/gin"
	"github.com/litsoftware/litmedia/cmd/http/proto"
	"github.com/litsoftware/litmedia/internal/controller"
	"github.com/litsoftware/litmedia/internal/controller/sts"
)

func StsToken(c *gin.Context) {
	rsp := &proto.StsTokenResponse{}
	err := controller.RunProc2(c, &sts.Statement{}, &proto.StsTokenRequest{}, rsp, sts.Process)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, rsp)
	return
}
