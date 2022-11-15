package upload

import (
	"fmt"
	"github.com/litsoftware/litmedia/internal/controller"
	"github.com/litsoftware/litmedia/internal/g"
	"strings"
)

type fileChecker struct {
	controller.BaseProc
}

func (c *fileChecker) Process() {
	r := <-c.In
	st := r.Req.(*Statement)
	if st.Err != nil {
		c.Next(r)
		return
	}

	file, header, err := r.GinCtx.Request.FormFile("attachment")
	if err != nil {
		g.App.Logger.Error(fmt.Sprintf("file err : %s", err.Error()))
		st.Err = err
		c.Next(r)
		return
	}

	st.Request.Filename = header.Filename

	if st.Request.Filename == "" {
		st.Err = fmt.Errorf("filename is empty")
		c.Next(r)
		return
	}

	chars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|", " ", "\t", "\r", "\n", "\b", "\f", "\v", "\a",
		"~", "`", "!", "@", "#", "$", "%", "^", "&", "(", ")", "-", "+", "=", "{", "}", "[", "]", ";", "'", ",", ".."}
	for _, char := range chars {
		st.Request.Filename = strings.ReplaceAll(st.Request.Filename, char, "")
	}

	rns := []rune(st.Request.Filename)
	if len(rns) > 64 {
		st.Request.Filename = string(rns[len(rns)-64:])
	}

	st.Request.Size = header.Size
	st.file = file
	st.fileHeader = header

	c.Next(r)
	return
}
