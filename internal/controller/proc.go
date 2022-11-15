package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/litsoftware/litmedia/internal/ent"
	"github.com/litsoftware/litmedia/internal/g"
	"github.com/trustmaster/goflow"
)

type BaseProc struct {
	In  <-chan *Context
	Out chan<- *Context
}

func (b *BaseProc) Next(r *Context) {
	b.Out <- r
}

type StatementInterface interface {
	Error() error
	SetError(err error)
	SetOrgRequest(r interface{})
	OrgRequest() interface{}
	SetOrgResponse(rsp interface{})
	OrgResponse() interface{}
}

type AppParamInterface interface {
	GetApp(v interface{}) interface{}
	SetApp(v interface{}, app *ent.App) error
}

type ClaimsParamInterface interface {
	GetClaims(v interface{}) interface{}
	SetClaims(v interface{}, claims interface{}) error
}

func RunProc2(ctx *gin.Context, st StatementInterface, r interface{}, rsp interface{}, f func() *goflow.Graph) error {
	c := NewContext()
	st.SetOrgRequest(r)
	st.SetOrgResponse(rsp)

	net := f()
	in := make(chan *Context)
	_ = net.SetInPort("In", in)
	wait := goflow.Run(net)

	c.GinCtx = ctx
	c.Headers = ctx.Request.Header
	c.Req = st

	in <- c
	close(in)
	<-wait

	if st.Error() != nil {
		switch st.Error().(type) {
		case *errors.Error:
			e := st.Error().(*errors.Error)
			g.App.Logger.Error("RunProc 出错了 ", e)
			return e
		default:
			g.App.Logger.Error("RunProc 出错了 %#v", st.Error())
			return errors.FromError(st.Error())
		}
	}

	return nil
}
