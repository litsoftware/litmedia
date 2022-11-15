package helpers

import (
	"testing"
)

func TestValidateIPv4(t *testing.T) {
	iplist := map[string]bool{
		"127.0.0.1": true,
		"0.0.0.0":   true,
		"1.1.1":     false,
	}

	for ip, r := range iplist {
		if r != ValidateIP(ip) {
			t.Errorf("ValidateIP not work: %s", ip)
		}
	}
}

func TestValidateIPv4Cidr(t *testing.T) {
	iplist := map[string]bool{
		"192.168.5.0/24": false,
		"192.168.5.0/32": false,
	}

	for ip, r := range iplist {
		if r != ValidateIP(ip) {
			t.Errorf("ValidateIP validate ip v4 not work: %s", ip)
		}
	}
}

func TestValidateIPv6(t *testing.T) {
	iplist := map[string]bool{
		"2001:db8::68":     true,
		"2001:db8:52:68":   false,
		"::ffff:192.0.2.1": true,
	}

	for ip, r := range iplist {
		if r != ValidateIP(ip) {
			t.Errorf("ValidateIP validate ip v6 not work: %s", ip)
		}
	}
}

func TestValidateEmail(t *testing.T) {
	emails := map[string]bool{
		"a@qq.com":             true,
		"1@qq.com":             true,
		"cc_@qq.baidu.com":     true,
		"google+1@apple.cn":    true,
		"_2020@apple.cn":       true,
		"20-20.33_44@apple.cn": true,
		"@apple.cn":            false,
		"1@qq":                 false,
	}

	for email, r := range emails {
		if r != ValidateEmail(email) {
			t.Errorf("ValidateEmail not work: %s", email)
		}
	}
}

func TestValidateURL(t *testing.T) {
	urls := map[string]bool{
		"http://baidu.com":        true,
		"https://qq.com":          true,
		"//qq.com":                true,
		"//a.b.c.qq.com":          true,
		"//a.b.c.qq.com?ab=1":     true,
		"//a.b.c.qq.com?ab=1&c=3": true,
		"//a.b.c.qq.com/a":        true,
		"//a.b.c.qq.com/b/c/d":    true,
		"a.b.c.qq.com/b/c/d?kkkk": true,
		"a.b.c.qq.com/a.html":     true,
		"localhost":               true,
		"com":                     true,
		"docker-mysql":            true,
		"128.1.1.1":               true,
		"0.0.0.0":                 true,
		":":                       false,
		":80":                     false,
	}

	for url, r := range urls {
		if r != ValidateURL(url) {
			t.Errorf("ValidateURL not work: %s", url)
		}
	}
}

func TestValidateHostname(t *testing.T) {
	urls := map[string]bool{
		"nginx":        true,
		"com":          true,
		"docker-mysql": true,
		"128.1.1.1":    true,
		"0.0.0.0":      true,
		":":            false,
		":80":          false,
	}

	for url, r := range urls {
		if r != ValidateHostname(url) {
			t.Errorf("ValidateHostname not work: %s", url)
		}
	}
}

func TestValidateHostList(t *testing.T) {
	if ValidateHostList([]string{
		"docker_nginx:980",
		"0.0.0.0:80",
	}) != true {
		t.Errorf("ValidateHostList not work 1")
	}

	if ValidateHostList([]string{
		"128.1.1.1/a",
		"docker.b",
		":80",
	}) != false {
		t.Errorf("ValidateHostList not work 2")
	}
}

func TestValidateHostPort(t *testing.T) {
	if true != ValidateHostPort("0.0.0.0:80", false) {
		t.Errorf("ValidateHostPort not work 1")
	}

	if true != ValidateHostPort("0.0.0.0:80", true) {
		t.Errorf("ValidateHostPort not work 2")
	}

	if true != ValidateHostPort(":80", true) {
		t.Errorf("ValidateHostPort not work 3")
	}

	if false != ValidateHostPort(":80", false) {
		t.Errorf("ValidateHostPort not work 4")
	}

	if false != ValidateHostPort(":800000", false) {
		t.Errorf("ValidateHostPort not work 4")
	}
}
