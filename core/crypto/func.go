package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func Hmac256(base string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(base))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
