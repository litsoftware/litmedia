package upload

import (
	"github.com/litsoftware/litmedia/cmd/http/proto"
	"github.com/litsoftware/litmedia/internal/controller"
)

type initSt struct {
	controller.BaseProc
}

func (c *initSt) Process() {
	r := <-c.In
	st := r.Req.(*Statement)
	if st.Err != nil {
		c.Next(r)
		return
	}

	st.Request = (st.OrgRequest()).(*proto.FileUploadRequest)
	st.Response = (st.OrgResponse()).(*proto.FileUploadResponse)

	c.Next(r)
	return
}
