package until

import (
	"encoding/json"
)

func JsonEncode(objcet interface{}) string {
	bytes, _ := json.Marshal(objcet)
	return string(bytes)
}
