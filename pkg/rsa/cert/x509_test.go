package cert

import (
	"testing"
)

func TestX509Cert(t *testing.T) {
	// 根证书
	rootKey, _ := genPrivateKey()
	rootCert, _ := GenCARoot(rootKey)
	//fmt.Println("rootCert\n", string(rootCertPEM))

	// 商户证书
	DCAPrivateKey, _ := genPrivateKey()
	DCACert, _ := GenDCA("1234", rootCert, rootKey, DCAPrivateKey)
	//fmt.Println("DCACert\n", string(DCACertPEM))
	verifyDCA(rootCert, DCACert)

	// 服务器证书
	ServerCert, _, _, _ := GenServerCert(DCACert, DCAPrivateKey)
	//fmt.Println("ServerPEM\n", string(ServerPEM))
	verifyLow(rootCert, DCACert, ServerCert)
}
