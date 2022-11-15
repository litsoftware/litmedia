package strconvh

import (
	_const "github.com/litsoftware/litmedia/pkg/const"
	"strconv"
)

func ParseInt(s string) (i int, err error) {
	valueInt64, err := strconv.ParseInt(s, 10, _const.Int32Bits)
	if err == nil {
		i = int(valueInt64)
	}
	return
}

func ParseInt8(s string) (i int8, err error) {
	valueInt64, err := strconv.ParseInt(s, 10, _const.Int8Bits)
	if err == nil {
		i = int8(valueInt64)
	}
	return
}

func ParseInt16(s string) (i int16, err error) {
	valueInt64, err := strconv.ParseInt(s, 10, _const.Int16Bits)
	if err == nil {
		i = int16(valueInt64)
	}
	return
}

func ParseInt32(s string) (i int32, err error) {
	valueInt64, err := strconv.ParseInt(s, 10, _const.Int32Bits)
	if err == nil {
		i = int32(valueInt64)
	}
	return
}

func ParseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, _const.Int64Bits)
}

func ParseUInt64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, _const.Int64Bits)
}
