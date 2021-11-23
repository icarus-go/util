package encrypt

import (
	"crypto/sha1"
	sha "crypto/sha256"
	base64Encoding "encoding/base64"
	hexEncoding "encoding/hex"
)

type sha256 struct{}

var Sha256 = new(sha256)

//SHA256生成哈希值
func (*sha256) Get(message string) string {
	hash := sha.New()
	//输入数据
	hash.Write([]byte(message))
	//计算哈希值
	bytes := hash.Sum(nil)

	//返回哈希值
	return hexEncoding.EncodeToString(bytes) //将字符串编码为16进制格式,返回字符串
}

//Sha1Encrypt
//  Author: Kevin·CC
//  Description: Shaun1加密
//  Param str
//  Return string
func Sha1Encrypt(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return base64Encoding.URLEncoding.EncodeToString(h.Sum(nil))
}
