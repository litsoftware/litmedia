package crypter

import (
	"github.com/litsoftware/litmedia/pkg/random"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	var a = "abc"
	var key = random.String(32)

	r, e := AesEncrypt(a, key)
	if e != nil {
		t.Errorf("encrypt err %#v", e)
	}

	oa, e := AesDecrypt(r, key)
	if e != nil {
		t.Errorf("decrypt err %#v", e)
	}

	if oa != a {
		t.Error("aes test err")
	}
}
