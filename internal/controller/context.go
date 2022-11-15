package controller

import (
	"entgo.io/ent/examples/start/ent"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"net/http"
	"time"
)

type Context struct {
	AesKey string
	Unique string
	Req    interface{}
	GinCtx *gin.Context

	Headers    http.Header
	AuthedUser *ent.User
}

func (*Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*Context) Done() <-chan struct{} {
	return nil
}

func (*Context) Err() error {
	return nil
}

func (*Context) Value(key interface{}) interface{} {
	return nil
}

func (c *Context) String() string {
	return "trade/Context"
}

func NewContext() *Context {
	ctx := new(Context)
	ctx.Unique, _ = uuid.GenerateUUID()
	return ctx
}
