package download

import (
	"context"
	"github.com/litsoftware/litmedia/cmd/http/proto"
	"github.com/litsoftware/litmedia/internal/common/types"
	"github.com/litsoftware/litmedia/internal/controller"
	"github.com/litsoftware/litmedia/internal/gateway/core"
	"github.com/litsoftware/litmedia/internal/proxy"
	"github.com/minio/minio-go/v7"
	"io"
)

type downloadFromOss struct {
	controller.BaseProc
}

func (c *downloadFromOss) Process() {
	r := <-c.In
	st := r.Req.(*Statement)
	if st.Err != nil {
		return
	}

	client := proxy.GetClient(st.App.Conf)
	ctx, cancel := context.WithCancel(r)
	defer cancel()
	res, err := client.(core.ObjectLayer).GetObject(ctx, st.Media.FullPath, minio.GetObjectOptions{})
	if err != nil {
		st.Response.Code = proto.DownloadStatusCode_Download_Failed
		st.Response.Message = err.Error()
		return
	}

	f := res.(io.ReadCloser)
	st.File = f

	st.Response.Code = proto.DownloadStatusCode_Download_Ok
	st.Response.Info = &proto.MediaInfo{
		Id:        st.Media.Sn,
		Filename:  st.Media.OrgFileName,
		Ext:       st.Media.Ext,
		FileClass: types.MediaType(st.Media.Type).String(),
		Mime:      st.Media.Mime,
		Hash:      st.Media.Hash,
	}

	return
}
