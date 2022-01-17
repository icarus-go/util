package pattern

import (
	"crypto/aes"
	"github.com/icarus-go/utils/encrypt/aes/padding"
)

type ecb struct {
	key []byte

	padding padding.Padding
}

func NewECB(key string) *ecb {
	return &ecb{
		key:     []byte(key),
		padding: padding.NewZero(),
	}
}

//SetPadding
//  Author: Kevin·CC
//  Description: 设置补位方式
//  Param padding
//  Return *ecb
func (e *ecb) SetPadding(padding padding.Padding) *ecb {
	e.padding = padding
	return e
}

//SetKey
//  Author: Kevin·CC
//  Description: 密钥
//  Param k
//  Return *ecb
func (e *ecb) SetKey(k string) *ecb {
	e.key = []byte(k)
	return e
}

//Encrypt
//  Author: Kevin·CC
//  Description: ecb 方式加密
//  Param value 加密前
//  Return string 加密后
//  Return error 错误
func (e *ecb) Encrypt(source string) ([]byte, error) {
	sourceBytes := []byte(source)

	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	sourceBytes = e.padding.Append(sourceBytes, block.BlockSize())

	blockSize := block.BlockSize()
	//返回加密结果
	encryptData := make([]byte, len(sourceBytes))
	//存储每次加密的数据
	tmpData := make([]byte, blockSize)

	//分组分块加密
	for index := 0; index < len(sourceBytes); index += blockSize {
		block.Encrypt(tmpData, sourceBytes[index:index+blockSize])
		copy(encryptData, tmpData)
	}
	return encryptData, nil

}

func (e *ecb) Decrypt(source string) ([]byte, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}
	sourceBytes := []byte(source)

	//返回加密结果
	decryptData := make([]byte, len(sourceBytes))
	//存储每次加密的数据
	tmpData := make([]byte, block.BlockSize())

	//分组分块加密
	for index := 0; index < len(sourceBytes); index += block.BlockSize() {
		length := index + block.BlockSize()
		var bytes []byte

		if length >= len(sourceBytes) {
			bytes = sourceBytes[index:]
		} else {
			bytes = sourceBytes[index:length]
		}

		block.Decrypt(tmpData, bytes)

		copy(decryptData, tmpData[0:])
	}

	return e.padding.WipeOff(decryptData, block.BlockSize()), nil
}
