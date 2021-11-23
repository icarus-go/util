package encrypt

import (
	"crypto/md5"
	hexEncoding "encoding/hex"
)

//MD5V
//  Author: Kevin·CC
//  Description: MD5加密
//  Param str
//  Return string
func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hexEncoding.EncodeToString(h.Sum(nil))
}
