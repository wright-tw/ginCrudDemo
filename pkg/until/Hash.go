package until

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data)) //将[]byte转成16进制
}

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
