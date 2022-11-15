package masker

import (
	"fmt"
	"github.com/apaxa-go/helper/stringsh"
	"math"
	"strings"
)

func Phone(s string) string {
	return overlay(s, "****", 3, 7)
}

func Tel(s string) string {
	l := len(s)
	return overlay(s, padStr(6), l-2, l)
}

func Email(s string) string {
	l := strLen(s)
	if l == 0 {
		return ""
	}

	tmp := strings.Split(s, "@")
	addr := tmp[0]
	domain := tmp[1]

	if strLen(addr) > 3 {
		addr = overlay(addr, padStr(3), 3, 30)
	} else {
		addr = addr + padStr(3)
	}

	return addr + "@" + domain
}

func IdCardNo(s string) string {
	l := strLen(s)
	if l == 0 {
		return ""
	}

	o := padStr(8)
	return overlay(s, o, 10, 18)
}

func BankNo(s string) string {
	l := strLen(s)
	if l == 0 {
		return ""
	}

	o := padStr(l - 10)
	return overlay(s, o, 6, l-4)
}

func Name(s string) string {
	l := strLen(s)
	if l == 0 {
		return ""
	}

	if strs := strings.Split(s, " "); len(strs) > 1 {
		tmp := make([]string, len(strs))
		for idx, str := range strs {
			tmp[idx] = Name(str)
		}
		return strings.Join(tmp, " ")
	}

	if l > 1 {
		return overlay(s, padStr(2), 1, l)
	}

	return s + padStr(2)
}

func strLen(s string) int {
	r := []rune(s)
	return len([]rune(r))
}

func padStr(l int) string {
	return stringsh.PadLeft("", "*", l)
}

func overlay(s string, overlay string, start int, end int) string {
	l := strLen(s)
	r := []rune(s)
	if l == 0 {
		return ""
	}

	if start < 0 {
		start = 0
	}
	if start > l {
		start = l
	}
	if end < 0 {
		end = 0
	}
	if end > l {
		end = l
	}
	if start > end {
		tmp := start
		start = end
		end = tmp
	}

	var overlayed = ""
	overlayed += string(r[:start])
	overlayed += overlay
	overlayed += string(r[end:])

	return overlayed
}

func Normal(s string) string {
	l := len(s)
	ml := int(math.Floor(float64(l) / 3.0))
	fmt.Println("ml ============= ", ml)
	o := padStr(l - ml*2)
	return overlay(s, o, ml, l-ml)
}

func Str(s string, o string, start int, end int) string {
	return overlay(s, o, start, end)
}
