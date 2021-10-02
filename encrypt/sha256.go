package encrypt

import (
	sha "crypto/sha256"
	"encoding/hex"
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
	return hex.EncodeToString(bytes) //将字符串编码为16进制格式,返回字符串
}
