package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5V
//  Author: Kevin·CC
//  Description: MD5加密
//  Param str
//  Return string
func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
