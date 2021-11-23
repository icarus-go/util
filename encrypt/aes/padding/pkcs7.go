package padding

import (
	"bytes"
	"crypto/aes"
)

type pkcs7 struct{}

func NewPKCS7() pkcs7 {
	return pkcs7{}
}

func (pkcs7) WipeOff(source []byte, blockSize int) []byte {
	for i := len(source) - 1; i >= 0; i-- {
		if source[i] > 15 {
			return source[:i+1]
		}
	} // fix: 修复死循环的问题
	return nil
}

func (pkcs7) Append(source []byte, blockSize int) []byte {
	//填充个数
	paddingCount := aes.BlockSize - len(source)%aes.BlockSize
	if paddingCount == 0 {
		return source
	}
	//填充数据
	return append(source, bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)...)
}
