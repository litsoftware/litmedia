package senstive_parameter_processor

import (
	"github.com/litsoftware/litmedia/pkg/base64"
	"github.com/litsoftware/litmedia/pkg/reflecth"
)

type Base64Processor struct {
}

func (f *Base64Processor) Encode(params map[string]string, data interface{}) error {
	if len(params) > 0 {
		for k, v := range params {
			if v == "" {
				continue
			}

			encodeStr := base64.Encode(v)
			err := reflecth.SetStructValeByField(data, k, encodeStr)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (f *Base64Processor) Decode(params map[string]string, data interface{}) error {
	if len(params) > 0 {
		for k, v := range params {
			if v == "" {
				continue
			}

			decodeStr, err := base64.Decode(v)
			if err != nil {
				return err
			}

			err = reflecth.SetStructValeByField(data, k, decodeStr)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
