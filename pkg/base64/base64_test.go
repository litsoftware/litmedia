package base64

import "testing"

func TestBase64Encode(t *testing.T) {
	var a = "askdjfalsfjlaskfj"
	ea := Encode(a)
	_a, err := Decode(ea)
	if err != nil || a != _a {
		t.Fatalf("base64 decode err %#v", err)
	}
}
