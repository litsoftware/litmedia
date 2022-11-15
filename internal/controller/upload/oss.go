package upload

import (
	"context"
	"github.com/litsoftware/litmedia/cmd/http/proto"
	"github.com/litsoftware/litmedia/internal/common/types"
	"github.com/litsoftware/litmedia/internal/controller"
	"github.com/litsoftware/litmedia/internal/g"
	"github.com/litsoftware/litmedia/internal/gateway/core"
	"github.com/litsoftware/litmedia/internal/proxy"
	"github.com/minio/minio-go/v7"
	"io"
)

type uploadMediaToOss struct {
	controller.BaseProc
}

func (c *uploadMediaToOss) Process() {
	r := <-c.In
	st := r.Req.(*Statement)
	if st.Err != nil {
		st.Response.Code = proto.UploadStatusCode_Failed
		st.Response.Message = st.Err.Error()
		return
	}

	client := proxy.GetClient(st.App.Conf)
	ctx, cancel := context.WithCancel(r)
	defer cancel()
	_, err := st.file.Seek(io.SeekStart, 0)
	if err != nil {
		g.App.Logger.Error("uploadMediaToOss.Process Seek", err)
		st.Err = err
		c.Next(r)
		return
	}

	_, err = client.(core.ObjectLayer).PutObject(ctx, st.Media.FullPath, st.file, -1, minio.PutObjectOptions{})
	if err != nil {
		st.Response.Code = proto.UploadStatusCode_Failed
		st.Response.Message = err.Error()
		return
	}

	st.Response.Code = proto.UploadStatusCode_Ok
	st.Response.Message = "ok"
	st.Response.Media = &proto.MediaInfo{
		Id:        st.Media.Sn,
		Filename:  st.Media.OrgFileName,
		Ext:       st.Media.Ext,
		FileClass: types.MediaType(st.Media.Type).String(),
		Mime:      st.Media.Mime,
		Hash:      st.Media.Hash,
	}

	return
}
