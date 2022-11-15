package strconvh

import "strconv"

func ParseFloat32(s string) (f float32, err error) {
	valueFloat64, err := strconv.ParseFloat(s, 32)
	if err == nil {
		f = float32(valueFloat64)
	}
	return
}

func ParseFloat64(s string) (f float64, err error) {
	f, err = strconv.ParseFloat(s, 32)
	return
}
