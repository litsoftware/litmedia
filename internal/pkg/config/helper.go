package config

import (
	"encoding/json"
)

func GetConf(key string, conf interface{}) error {
	b, err := json.Marshal(GetMap(key))
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, conf)
	return err
}
