package senstive_parameter_processor

import (
	"github.com/litsoftware/litmedia/internal/g"
	"github.com/litsoftware/litmedia/pkg/aes"
	"github.com/litsoftware/litmedia/pkg/reflecth"
)

type AesProcessor struct {
	Key string
}

func (f *AesProcessor) Encode(params map[string]string, data interface{}) error {
	if len(params) > 0 {
		for k, v := range params {
			if v == "" {
				continue
			}

			encryptedStr, err := aes.Encrypt([]byte(f.Key), v)
			if err != nil {
				g.App.Logger.Errorf("aes encrypt %s err: %#v", k, err)
				return err
			}

			err = reflecth.SetStructValeByField(data, k, encryptedStr)
			if err != nil {
				g.App.Logger.Errorf("SetStructValeByField %s err: %#v", k, err)
				return err
			}
		}
	}

	return nil
}

func (f *AesProcessor) Decode(params map[string]string, data interface{}) error {
	if len(params) > 0 {
		for k, v := range params {
			if v == "" {
				continue
			}

			decryptStr, err := aes.Decrypt([]byte(f.Key), v)
			if err != nil {
				g.App.Logger.Errorf("aes decrypt %s err: %#v", k, err)
				return err
			}

			err = reflecth.SetStructValeByField(data, k, decryptStr)
			if err != nil {
				g.App.Logger.Errorf("SetStructValeByField %s err: %#v", k, err)
				return err
			}
		}
	}

	return nil
}
