package encrypt

import base64Encoding "encoding/base64"

type base64 struct{}

//Base64 base64encoding
var Base64 = new(base64)

//Encrypt
//  Author: Kevin·CC
//  Description: 加密
//  Param value
//  Return string
func (*base64) Encrypt(value []byte) string {
	return base64Encoding.URLEncoding.EncodeToString(value)
}

//Decrypt
//  Author: Kevin·CC
//  Description: 解密
//  Param str
//  Return string
//  Return error
func (a *base64) Decrypt(str string) (string, error) {
	buf, err := base64Encoding.URLEncoding.DecodeString(str)
	return string(buf), err
}
