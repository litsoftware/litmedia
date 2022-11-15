package http

import (
	"github.com/gin-gonic/gin"
	"github.com/litsoftware/litmedia/cmd/http/proto"
	"github.com/litsoftware/litmedia/internal/controller"
	"github.com/litsoftware/litmedia/internal/controller/download"
	"github.com/litsoftware/litmedia/internal/g"
	"io"
)

func FileDownload(c *gin.Context) {
	c.Param("media_id")
	rsp := &proto.FileDownloadResponse{}
	st := &download.Statement{}
	req := &proto.FileDownloadRequest{
		MediaId: c.Param("media_id"),
	}

	err := controller.RunProc2(c, st, req, rsp, download.Process)
	if err != nil {
		g.App.Logger.Error(err)
		c.JSON(500, err)
		return
	}

	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+rsp.Info.Filename)
	c.Writer.Header().Set("Content-Type", rsp.Info.Mime)
	c.Writer.Header().Set("Media-Name", rsp.Info.Filename)
	c.Writer.Header().Set("Media-Hash", rsp.Info.Hash)
	c.Writer.Header().Set("Media-Class", rsp.Info.FileClass)
	c.Writer.Header().Set("Media-ID", rsp.Info.Id)
	c.Writer.Header().Set("Media-Ext", rsp.Info.Ext)

	defer func() {
		st.File.Close()
	}()

	_, err = io.Copy(c.Writer, st.File)
	if err != nil {
		g.App.Logger.Error(err)
		c.JSON(500, err)
		return
	}
	return
}
