package encrypt

import (
	"pmo-test4.yz-intelligence.com/base/utils/encrypt/aes/padding"
	"pmo-test4.yz-intelligence.com/base/utils/encrypt/aes/pattern"
)

//AesECBEncrypt
//  Author: Kevin·CC
//  Description: AES加密算法 ECB 模式 Base64Encode 编码
//  Param value 要加密的值
//  Param key 加密秘钥
//  Param padding 填充类算法类型 padding.NewZero()、padding.NewPKCS5()、padding.NewPKCS7()
//  Return string 加密后的值 并且 base64 URLEncode
//  Return error 错误信息
func AesECBEncrypt(value, key string, padding padding.Padding) (string, error) {
	encrypt, err := pattern.NewECB(key).SetPadding(padding).Encrypt(value)
	value = Base64.Encrypt(encrypt)
	return value, err
}

//AesECBDecrypt
//  Author: Kevin·CC
//  Description: Aes加密算法 ECB 模式 Base64Encode 编码
//  Param value 要解密的值
//  Param key 秘钥
//  Param padding 填充类算法类型 padding.NewZero()、padding.NewPKCS5()、padding.NewPKCS7()
//  Return string 解密后的值
//  Return error 错误
func AesECBDecrypt(value, key string, padding padding.Padding) (string, error) {
	value, err := Base64.Decrypt(value)
	if err != nil {
		return "", err
	}

	decrypt, err := pattern.NewECB(key).SetPadding(padding).Decrypt(value)
	return string(decrypt), err
}

//AesECBHexEncrypt
//  Author: Kevin·CC
//  Description: Aes加密算法 ECB 模式 Hex 编码
//  Param value 要加密的值
//  Param key 秘钥
//  Param padding 填充类算法类型 padding.NewZero()、padding.NewPKCS5()、padding.NewPKCS7()
//  Return string 加密后的值
//  Return error 错误信息
func AesECBHexEncrypt(value, key string, padding padding.Padding) (string, error) {
	encrypt, err := pattern.NewECB(key).SetPadding(padding).Encrypt(value)
	value = Hex.Encrypt(encrypt)
	return value, err
}

//AesECBHexDecrypt
//  Author: Kevin·CC
//  Description: Aes加密算法
//  Param value
//  Param key
//  Param padding 填充类算法类型 padding.NewZero()、padding.NewPKCS5()、padding.NewPKCS7()
//  Return string 加密后的值
//  Return error 错误信息
func AesECBHexDecrypt(value, key string, padding padding.Padding) (string, error) {
	decrypt, err := Hex.Decrypt(value)
	if err != nil {
		return "", err

	}
	decryptBytes, err := pattern.NewECB(key).SetPadding(padding).Decrypt(decrypt)
	if err != nil {
		return "", err
	}
	return string(decryptBytes), nil
}
