package qr

import (
	"encoding/base64"
	"github.com/mdp/qrterminal/v3"
	"github.com/skip2/go-qrcode"
	"net/http"
	"os"
)

func GeneratePng(content string, size int) ([]byte, error) {
	if size == 0 {
		size = 256
	}

	return qrcode.Encode(content, qrcode.Medium, size)
}

func GenerateDataURI(content string) (string, error) {
	png, err := GeneratePng(content, 256)
	if err != nil {
		return "", err
	}

	var base64Encoding string
	mimeType := http.DetectContentType(png)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(png)
	return base64Encoding, err
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func DisplayQrTerminal(content string) {
	qrterminal.Generate(content, qrterminal.L, os.Stdout)
}
