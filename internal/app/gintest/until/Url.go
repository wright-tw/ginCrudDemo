package until

import (
	"net/url"
)

func HttpQueryBuild(dataMap map[string]string) string {
	str := ""
	params := url.Values{}
	for key, value := range dataMap {
		params.Set(key, value)
	}
	str = params.Encode()
	return str
}
