package padding

import "bytes"

type zero struct{}

func NewZero() zero {
	return zero{}
}

//WipeOff
//  Author: Kevin·CC
//  Description: 去除位
//  Param content
//  Return []byte
func (zero) WipeOff(content []byte, blockSize int) []byte {
	index := bytes.IndexByte(content, 0)
	if index < 0 {
		return []byte{}
	}

	return content[0:index]
}

//Append
//  Author: Kevin·CC
//  Description: 补位
//  Param content
//  Param blockSize
//  Return []byte
func (zero) Append(content []byte, blockSize int) []byte {
	bit := blockSize - len(content)%blockSize // 补多少位

	paddingText := bytes.Repeat([]byte{0}, bit) // 返回一个新的字节片，其中包含bit的拷贝数。

	return append(content, paddingText...) // 将补位后的数组追加到原内容中
}
