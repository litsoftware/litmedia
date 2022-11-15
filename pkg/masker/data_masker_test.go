package masker

import (
	"fmt"
	"testing"
)

func TestIdCardNo(t *testing.T) {
	s := "421124202106068888"
	ms := IdCardNo(s)
	if ms != "4211242021"+padStr(8) {
		t.Errorf("mask phone number  err")
	}
}

func TestPadStr(t *testing.T) {
	s := "***"
	if s != padStr(3) {
		t.Errorf("pad str err")
	}
}

func TestBankNo(t *testing.T) {
	s := "6225889977664432"
	ms := BankNo(s)
	ems := fmt.Sprintf("622588%s4432", padStr(6))
	if ms != ems {
		t.Errorf("mask bank number err  %s  !=  %s", ems, ms)
	}
}

func TestEmail(t *testing.T) {
	s1 := "abcdefgx.dds@qq.com"
	es1 := "abc***@qq.com"
	ms1 := Email(s1)
	if es1 != ms1 {
		t.Errorf("mask email err 1  %s  !=  %s", es1, ms1)
	}

	s2 := "gad@qq.com"
	es2 := "gad***@qq.com"
	ms2 := Email(s2)
	if es2 != ms2 {
		t.Errorf("mask email err 2  %s  !=  %s", es2, ms2)
	}

	s3 := "dd@qq.com"
	es3 := "dd***@qq.com"
	ms3 := Email(s3)
	if es3 != ms3 {
		t.Errorf("mask email err 3  %s  !=  %s", es3, ms3)
	}
}

func TestStrLen(t *testing.T) {
	s1 := "ab"
	if strLen(s1) != 2 {
		t.Errorf("strlen err 1")
	}

	s2 := "中国"
	if strLen(s2) != 2 {
		t.Errorf("strlen err 2")
	}
}

func TestNormal(t *testing.T) {
	s1 := "abcdef"
	es1 := "ab**ef"
	ms1 := Normal(s1)
	if es1 != ms1 {
		t.Errorf("Normal err 1. %s != %s", ms1, es1)
	}

	s2 := "abcde"
	es2 := "a***e"
	ms2 := Normal(s2)
	if es2 != ms2 {
		t.Errorf("Normal err 2. %s != %s", ms2, es2)
	}
}

func TestName(t *testing.T) {
	s1 := "刘大宝"
	es1 := "刘**"
	ms1 := Name(s1)
	if es1 != ms1 {
		t.Errorf("Normal err 1. %s != %s", ms1, es1)
	}

	s2 := "雷霆万钧"
	es2 := "雷**"
	ms2 := Name(s2)
	if es2 != ms2 {
		t.Errorf("Normal err 2. %s != %s", ms2, es2)
	}

	s3 := "天使"
	es3 := "天**"
	ms3 := Name(s3)
	if es3 != ms3 {
		t.Errorf("Normal err 3. %s != %s", ms3, es3)
	}

	s4 := "天"
	es4 := "天**"
	ms4 := Name(s4)
	if es4 != ms4 {
		t.Errorf("Normal err 4. %s != %s", ms4, es4)
	}
}
