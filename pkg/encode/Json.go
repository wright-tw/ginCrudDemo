package encode

import (
	"bytes"
	"encoding/json"
)

func JSONEncode(objcet interface{}) string {
	byteString, _ := json.Marshal(objcet)
	return string(byteString)
}

func JSONDecode(data string) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	r := bytes.NewReader([]byte(data))
	Err := json.NewDecoder(r).Decode(&m)
	return m, Err
}
