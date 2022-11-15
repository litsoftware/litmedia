package aes

import (
	"testing"
)

var key = []byte("LKHlhb899Y09olUi")

func TestEncrypt(t *testing.T) {
	encryptMsg, err := Encrypt(key, "Hello World")
	if err != nil {
		t.Fatal(err)
	}

	msg, err := Decrypt(key, encryptMsg)
	if err != nil {
		t.Fatal(err)
	}
	if msg != "Hello World" {
		t.Fatal("decrypt error")
	}
}
