package download

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/litsoftware/litmedia/internal/common"
	"github.com/litsoftware/litmedia/internal/controller"
	"github.com/litsoftware/litmedia/internal/errors"
)

type queryMedia struct {
	controller.BaseProc
}

func (c *queryMedia) Process() {
	r := <-c.In
	st := r.Req.(*Statement)
	if st.Err != nil {
		c.Next(r)
		return
	}

	spew.Dump("media_id", st.Request.MediaId, st.App)
	media, err := common.GetMediaById(context.Background(), st.Request.MediaId, st.App)
	if err != nil {
		st.Err = errors.ErrorInternalError("资源未找到", "")
		c.Next(r)
		return
	}

	st.Media = media

	c.Next(r)
	return
}
