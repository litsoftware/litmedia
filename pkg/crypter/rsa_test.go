package crypter

import (
	"github.com/litsoftware/litmedia/pkg/rsa"
	"testing"
)

func TestRsaEncode(t *testing.T) {
	var pemPubkey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAq0MSoobqG5slXUuVXtg7
wywzQ6NUfjD4L633RgKyrizhx4JWt+ybOIRlzVy0Hq3GNUf3yWjsgUcUKT+FMg0e
BcmffbRfm0Eaw1yK5524ZX9nkxVWZyLPte8rFbB8Khr1mb9N9Ox8R7bPRwGtnqjE
ZuhPj4jG9pjcCHwJtN3OrBrIVQgTCWDaoV3kHa1Fu5vQRuaYEvZrHv6e0zDS7h0w
cag7Kdnf+JYf192yYqerYgqviGA4+foqoNtn8x8DxCKHS+lyg8AL2NGk3ZZRvG4W
CnMzp2XaRf3+ndzPvrGWich0ZvoLkrlVzGlM5IJX+f+8FQGpevZQEbGRyXEth+8v
3QIDAQAB
-----END PUBLIC KEY-----
`
	var pemPriv = `-----BEGIN PRIVATE KEY-----
MIIEpQIBAAKCAQEAq0MSoobqG5slXUuVXtg7wywzQ6NUfjD4L633RgKyrizhx4JW
t+ybOIRlzVy0Hq3GNUf3yWjsgUcUKT+FMg0eBcmffbRfm0Eaw1yK5524ZX9nkxVW
ZyLPte8rFbB8Khr1mb9N9Ox8R7bPRwGtnqjEZuhPj4jG9pjcCHwJtN3OrBrIVQgT
CWDaoV3kHa1Fu5vQRuaYEvZrHv6e0zDS7h0wcag7Kdnf+JYf192yYqerYgqviGA4
+foqoNtn8x8DxCKHS+lyg8AL2NGk3ZZRvG4WCnMzp2XaRf3+ndzPvrGWich0ZvoL
krlVzGlM5IJX+f+8FQGpevZQEbGRyXEth+8v3QIDAQABAoIBAQCCe5vbEIeVeONC
7a9kj8MYtLhqNCrP6mdtjFH1mWChq7hp1ThU8YRhzx3xFUx2g1eciLSVU8e9x3Xo
52iH9c6GAlPt5lthYn0Nk5iRV22Ch0tWmM83zSIML6jX4Zr4SYoOFd9DWFpoGTRb
mT+6vkLFVQIpDcDpaRVbsYWDs+se3BnXVW+HinLfFq+WjLd/pjeKEVDVjaMkgjLm
Pq/chTW0F8Y3O6MR9JcFxLs9opFKj/aJ7Myxf1ckibmTPLbok6GETd7miQCeOPy1
YSiELbBIEQBjr9bC2Yb1ya0sWZfmpaAFx9y3RYW93C41ubh8y8G7OpR8Bw2dShro
/m6HgvUBAoGBANZOXVsfpLW0fO4xmQ8zYtm2Wq2QEnGUZD2eduyvhezZssDZwGyf
7LqDa/ZvzZKFHp4wz+t7QDKAXlsOwq8jcPCtl730g/ksNDOx1kSAIhGmXNMnfG6U
JvEqUp9m/vAe+Z6/KAuuzaEJ1H9zWm8aKtjJhQHqVZBL+hJEm9Z+ImXRAoGBAMyU
30502a8JvvHgNA0f85wP1asGUBv70frJizZvNafUWNjhPqkcH8hpchPf1PSEZ1XE
ZwcEpI5EC3JB8mq1RUaD03u65X4zfLeKtnG84/+vh4RDCw/DuLEL2UCIjdENSL2F
U539dksQydsm6u2pCIOHOY3gLi8BYCyr713wgZBNAoGAQQR3aEz6YLJIcM+VWzpA
1EJx0lRydAkMPHGWLQq3e/s0MzQdIBeI77EtkQ+sc7Z1apLukcfL20Z82e1pfU8o
Vayk69mCLJePottiYVfqfiZV5S4Gn1nNUr9/X6MU127eVp4yHRFEi4X69Uve9PQg
abMB6mnY2bVjQiWUjgeokfECgYEAp9SprS9nNZWFM+B+UxYXP7wtyXQQXCSmCls0
/Z8WIWyuxp1iTlFoMX1vvKYvibjlHRNWggdpB90CluVK/gdfoY2b/TFB+9o5Qkll
isEvOh7BydtGQ9SIu5XA0JQ1435GYlPyWPYKCWLodgTjTb0R4vUYXi+/M4ipYXxb
jnEzmTECgYEAxh1JHTYpQkLgnoorO2K4FNjkUH9l7K17s/ZvklK7M7HLq4laQ2hU
r0FlKG5DOveRU+PcI8FYcnllfMQo90MIrueULTTnRPAcGfgec7sazKHBR/3mAYB9
iMrurHnUfEmmMSSoxiJfnAzvaiyFB+WWGMWoTS7Q8uu4Fh9EeF5NjWM=
-----END PRIVATE KEY-----
`
	var data = "abc中文"

	ed, err := rsa.PrivateEncrypt(data, pemPriv)
	if err != nil {
		t.Fatalf("encrypt err %#v", err)
		return
	}

	_data, err := rsa.PublicDecrypt(ed, pemPubkey)
	if err != nil {
		t.Fatalf("decrypt err %#v", err)
		return
	}

	if _data != data {
		t.Fatalf("data not eq decrypted data")
		return
	}

	encryptedData, err := rsa.PublicEncrypt(data, pemPubkey)
	if err != nil {
		t.Fatalf("PublicEncrypt err %#v", err)
		return
	}

	priDecryptedData, err := rsa.PrivateDecrypt(encryptedData, pemPriv)
	if err != nil {
		t.Fatalf("PrivateDecrypt err %#v", err)
		return
	}

	if priDecryptedData != data {
		t.Fatalf("PrivateDecrypt data not eq origin data")
		return
	}
}
