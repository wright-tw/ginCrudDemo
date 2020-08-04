package encode

import (
	"net/url"
)

func HTTPQueryBuild(dataMap map[string]string, unescape bool) string {
	str := ""
	params := url.Values{}
	for key, value := range dataMap {
		params.Set(key, value)
	}
	str = params.Encode()

	if unescape {
		str, _ = url.QueryUnescape(str)
	}

	return str
}
