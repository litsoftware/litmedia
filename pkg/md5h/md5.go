package md5h

import (
	"crypto/md5"
	"fmt"
	"io"
)

func New(s string) string {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	b := h.Sum(nil)
	return fmt.Sprintf("%x", b)
}
