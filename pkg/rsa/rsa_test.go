package rsa

import (
	"testing"
)

// 公钥生成 https://miniu.alipay.com/keytool/create

var (
	pub  = `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlIg/mNdHk3RYPIXMekUIRvmPwIiJmhmBtNgABzbGnubi0F5SoWN+qk/VNmfueMhBEMv74spQFTpewYV+ygcoQ2snwzSQ5YLqxN3ZtpIZSNyWG+Ax0WY5hdUnOKy6IVOPs9Vdi/rYyGDwyaEkzXjWyGzYHsnjjAPFCJSzg1Q1emi3paHiKCBPawrOVpCqL36/bKbMejV6IetFOED7itIG/cMmjk9m6uTt1JYvy8nSFEwX9DQFSUpQ3rmyW7RtiH+Jf5WpbH12y6E/X5kGiq80TqXHto4W9FKo9UcwXKR4ZQJ1qeoMSQwSJEdzdhXM2uFzuwUKL6JL3Fz4MY87uGbrSwIDAQAB`
	priv = `MIIEowIBAAKCAQEAlIg/mNdHk3RYPIXMekUIRvmPwIiJmhmBtNgABzbGnubi0F5SoWN+qk/VNmfueMhBEMv74spQFTpewYV+ygcoQ2snwzSQ5YLqxN3ZtpIZSNyWG+Ax0WY5hdUnOKy6IVOPs9Vdi/rYyGDwyaEkzXjWyGzYHsnjjAPFCJSzg1Q1emi3paHiKCBPawrOVpCqL36/bKbMejV6IetFOED7itIG/cMmjk9m6uTt1JYvy8nSFEwX9DQFSUpQ3rmyW7RtiH+Jf5WpbH12y6E/X5kGiq80TqXHto4W9FKo9UcwXKR4ZQJ1qeoMSQwSJEdzdhXM2uFzuwUKL6JL3Fz4MY87uGbrSwIDAQABAoIBAAcT8wvifluteKLxsUvGTF2teMcw+nWob4Dhpiax2ocp83cAd6mPJzMQeNWN4FRLHqahQdCN2YEZdfh81wzjiGymB5AKyjjVo0BcXgqQLbpGILUVWupil60j+il+OaRq6fck3L9V5cyuqerBzhIohvuoDChzcG90oogFJgTFF4NMtWOfuihuQA56tqSgeTZhhbL7K9x/pDkH8U6SeFXI5iQldLzIvcPc0a6TSuNR4+a/9Yyom3iBCMV3dyzvS2XdJn9s5opRbluedPXO91WxWcQ6zSQIGbpgC5z31g7OudXOY3v7sv4Gi4od8wrM9xKUyjmKZbT6P0EC8UHYybyHi9ECgYEA4YMhqG8nt2mZ6QuBJKmLefUbEIV+4lCx03YLZ1DNpyZgBh/JtUHzrYn2C9blxp1nlyHKY/n6MglojTUvD00UPbFndpm+rFG+AjKutvqaH8hdj0l2H4DiQCaJPtXWtW83+D2YUZYh+H77n/NAbpzMttiLW3MQtEt1xuvFr8R3MmMCgYEAqJzgUKlx34H+H8ssjZGHt4gIbV8SgcslFmYsKOdeodNSVPTV+2K54AYERdfPCqyYknoyiLV9wavmxjGMePEwaQW5ovGRu42Tg6ZnjY/cujmsXEMFn92HsMc+4k8UiHha0dMEBXr3FmT9XZK8S4IMJYlR/NNcX3nFUNlqfifNQ/kCgYEAlzbf1dqbDiCQZLUjJptNfDy/pidtuoGTjBDmeqOzErbnwpOEJLeRlzcgNjYmIzUe5jdxR/KyMuUeJzmXUbJEtU2E7AER2uiA+WZUwzttJ4yqN89xlpRMoel1NB4dd9GY/SsIPQTnyIIVZd6twL4Bg0XWxD523/6NfhJQn5ikuLMCgYBt3aJwDKjNWTeL4ehK6ovDM9lB4tP9TLAC01ps4K2Rxk98WfhgD+lLpr/7/m3C58mWkBcDazbhDjCV7c05NPc1R2y5Vqx9x/dZrgEFhbrugQvlJiqxCzkSIBHChoK7laif1d5l78S2i+FnqmgUz0043CxXECW9dGv0jWZNg4PFiQKBgFOe97RGrseGSRuNdN9fL2E1EalScmH7HTXvX/CGa4B7LL08FS765mXz+JvoTA3SSR99WZDP/n/U9slU3k+t/preCvsX6qdhm8eFxLgNbV6Eu52s1DqyFMhkLgCTiMpoaYd9SyLbSzUshJg7KOhBiookJ4yK7cxOSvpzcnCawfsL`

	data = "a"
	sign = ""
)

func TestPrivateSign(t *testing.T) {
	priv = PrivateKeyFormat(priv)

	p, err := ParseRsaPrivateKeyFromPemStr(priv)
	if err != nil {
		t.Fatalf("parse err %#v", err)
	}

	sign, err = PrivateSign(data, p)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPublicVerify(t *testing.T) {
	pub = PublicKeyFormat(pub)

	p, err := ParseRsaPublicKeyFromPemStr(pub)
	if err != nil {
		t.Fatalf("parse err %#v", err)
	}

	err = PublicVerify(data, sign, p)
	if err != nil {
		t.Fatal(err)
	}
}
