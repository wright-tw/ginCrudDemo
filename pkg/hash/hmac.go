package hash

import (
	"crypto/hmac"
	"crypto/sha1"
)

func HmacWithSha1(str string, key string) []byte {
	fMethod := sha1.New
	byteKey := []byte(key)

	hashHandler := hmac.New(fMethod, byteKey)
	_, writeErr := hashHandler.Write([]byte(str))
	if writeErr != nil {
		return []byte{}
	}

	return hashHandler.Sum(nil)
}
