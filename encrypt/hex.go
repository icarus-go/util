package encrypt

import hexEncoding "encoding/hex"

//
type hex struct{}

// Hex编码
var Hex hex

//Encrypt Hex编码加密
func (a *hex) Encrypt(value []byte) string {
	return hexEncoding.EncodeToString(value)
}

//Decrypt 解密
func (a *hex) Decrypt(str string) (string, error) {
	buf, err := hexEncoding.DecodeString(str)
	return string(buf), err
}
