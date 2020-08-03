package until

import (
	"encoding/json"
)

func JSONEncode(objcet interface{}) string {
	bytes, _ := json.Marshal(objcet)
	return string(bytes)
}
