package urlh

import "net/url"

func AppendQueryParams(urlStr string, params map[string]string) string {
	if params == nil {
		return urlStr
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}

	q := u.Query()
	for k, v := range params {
		ev := url.QueryEscape(v)
		q.Set(k, ev)
	}

	u.RawQuery = q.Encode()
	return u.String()
}

func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
