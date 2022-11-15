package d

import "github.com/litsoftware/litmedia/pkg/jsonh"

type H = map[string]interface{}

func ConvertMSStoH(d map[string]string) *H {
	s := jsonh.Marshal(d)
	h := new(H)
	jsonh.UnMarshal(s, h)
	return h
}

func ConvertToH(d interface{}) *H {
	s := jsonh.Marshal(d)
	h := new(H)
	jsonh.UnMarshal(s, h)
	return h
}

func S2H(d interface{}, h *H) {
	s := jsonh.Marshal(d)
	jsonh.UnMarshal(s, h)
}

func H2I(h *H) *map[string]interface{} {
	i := &map[string]interface{}{}
	s := jsonh.Marshal(h)
	jsonh.UnMarshal(s, i)
	return i
}
