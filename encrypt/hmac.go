package encrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	base64Encoding "encoding/base64"
)

// HmacEncrypt 获得Hmac-Sha1加密
func HmacEncrypt(key string, data []byte) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write(data)
	return base64Encoding.URLEncoding.EncodeToString(h.Sum(nil))
}
