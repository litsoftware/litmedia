package controller

import (
	"github.com/litsoftware/litmedia/internal/ent"
	"github.com/litsoftware/litmedia/pkg/reflecth"
)

type BaseStatement struct {
	Err error
	req interface{}
	rsp interface{}
}

func (b *BaseStatement) Error() error {
	return b.Err
}
func (b *BaseStatement) SetError(err error) {
	b.Err = err
}
func (b *BaseStatement) SetOrgRequest(r interface{}) {
	b.req = r
}
func (b *BaseStatement) OrgRequest() interface{} {
	return b.req
}
func (b *BaseStatement) SetOrgResponse(rsp interface{}) {
	b.rsp = rsp
}
func (b *BaseStatement) OrgResponse() interface{} {
	return b.rsp
}

type StatementAppImpl struct {
	BaseStatement
}

func (b *StatementAppImpl) GetApp(v interface{}) interface{} {
	return reflecth.GetStructValeByField(v, "App")
}

func (b *StatementAppImpl) SetApp(v interface{}, app *ent.App) error {
	return reflecth.SetStructValeByField(v, "App", app)
}

type ClaimsImpl struct {
	BaseStatement
}

func (c *ClaimsImpl) SetClaims(v interface{}, claims interface{}) error {
	return reflecth.SetStructValeByField(v, "Claims", claims)
}

func (c *ClaimsImpl) GetClaims(v interface{}) interface{} {
	return reflecth.GetStructValeByField(v, "Claims")
}
