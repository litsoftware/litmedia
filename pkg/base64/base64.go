package base64

import "encoding/base64"

// Base64 base64
type Base64 struct{}

func NewBase64() *Base64 {
	return &Base64{}
}

// Encode base64 encode
func (b *Base64) Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func (b *Base64) Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func (b *Base64) DecodeBytes(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(data))
}

func Encode(s string) string {
	return NewBase64().Encode([]byte(s))
}

func Decode(s string) (string, error) {
	r, err := NewBase64().Decode(s)
	return string(r), err
}
