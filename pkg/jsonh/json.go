package jsonh

import (
	"bytes"
	"encoding/json"
	"os"
)

func Marshal(d interface{}) string {
	b, err := json.Marshal(d)
	if err != nil {
		return "{}"
	}

	return string(b)
}

func UnMarshal(s string, d interface{}) {
	err := json.Unmarshal([]byte(s), d)
	if err != nil {
		d = nil
	}
}

func ConvertTo(d interface{}, t interface{}) {
	b, err := json.Marshal(d)
	if err != nil {
		return
	}

	_ = json.Unmarshal(b, t)
}

func PrettyString(content []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, content, "", "    "); err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}

func Load(path string, o interface{}) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, o)
	if err != nil {
		return err
	}

	return nil
}
