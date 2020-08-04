package encode

import (
	b64 "encoding/base64"
)

func Base64EncodeByByte(byte []byte) string {
	return b64.StdEncoding.EncodeToString(byte)
}

func Base64Encode(str string) string {
	return b64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) string {
	byte, _ := b64.StdEncoding.DecodeString(str)
	return string(byte)
}
