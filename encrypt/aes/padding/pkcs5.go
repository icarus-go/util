package padding

import "bytes"

type pkcs5 struct{}

func NewPKCS5() pkcs5 {
	return pkcs5{}
}

func (pkcs5) WipeOff(source []byte, blockSize int) []byte {
	length := len(source)
	return source[:(length - int(source[length-1]))] // 去掉最后一个字节 unpadding 次
}

func (pkcs5) Append(source []byte, blockSize int) []byte {
	padding := blockSize - len(source)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(source, padtext...)
}
