package encrypt

import (
	"pmo-test4.yz-intelligence.com/base/utils/encrypt/aes/padding"
	"pmo-test4.yz-intelligence.com/base/utils/encrypt/aes/pattern"
	"testing"
)

func Test_AES_Encrypt_PKCS7(t *testing.T) {

	encrypt, err := pattern.NewECB("Hcdipgaf38gFPg1z").SetPadding(padding.NewPKCS7()).Encrypt("13739092481")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	base64Val := Base64.Encrypt(encrypt)

	println(base64Val)

	bytes, _ := Base64.Decrypt(base64Val)

	decrypt, err := pattern.NewECB("Hcdipgaf38gFPg1z").SetPadding(padding.NewPKCS7()).Decrypt(bytes)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	println(string(decrypt))
}

func Test_AES_Encrypt_PKCS5(t *testing.T) {

	encrypt, err := pattern.NewECB("Hcdipgaf38gFPg1z").SetPadding(padding.NewPKCS5()).Encrypt("13739092481")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	base64Val := Base64.Encrypt(encrypt)

	println(base64Val)

	bytes, _ := Base64.Decrypt(base64Val)

	decrypt, err := pattern.NewECB("Hcdipgaf38gFPg1z").SetPadding(padding.NewPKCS5()).Decrypt(bytes)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	println(string(decrypt))
}

func Test_AES_Encrypt_Zero(t *testing.T) {

	encrypt, err := pattern.NewECB("Hcdipgaf38gFPg1z").SetPadding(padding.NewZero()).Encrypt("13739092481")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	base64Val := Base64.Encrypt(encrypt)

	println(base64Val)

	bytes, _ := Base64.Decrypt(base64Val)

	decrypt, err := pattern.NewECB("Hcdipgaf38gFPg1z").SetPadding(padding.NewZero()).Decrypt(bytes)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	println(string(decrypt))
}

func Test_AES_Encrypt_ZERO_HEX(t *testing.T) {
	encrypt, err := pattern.NewECB("Hcdipgaf38gFPg1z").SetPadding(padding.NewZero()).Encrypt("13739092481")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	base64Val := Hex.Encrypt(encrypt)

	println(base64Val)

	bytes, _ := Hex.Decrypt(base64Val)

	decrypt, err := pattern.NewECB("Hcdipgaf38gFPg1z").SetPadding(padding.NewZero()).Decrypt(bytes)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	println(string(decrypt))
}

func Test_AES_Encrypt_PKCS5_HEX(t *testing.T) {
	encrypt, err := pattern.NewECB("Hcdipgaf38gFPg1z").SetPadding(padding.NewPKCS5()).Encrypt("13739092481")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	base64Val := Hex.Encrypt(encrypt)

	println(base64Val)

	bytes, _ := Hex.Decrypt(base64Val)

	decrypt, err := pattern.NewECB("Hcdipgaf38gFPg1z").SetPadding(padding.NewPKCS5()).Decrypt(bytes)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	println(string(decrypt))
}

func Test_AES_Encrypt_PKCS7_HEX(t *testing.T) {
	encrypt, err := pattern.NewECB("Hcdipgaf38gFPg1z").SetPadding(padding.NewPKCS7()).Encrypt("13739092481")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	base64Val := Hex.Encrypt(encrypt)

	println(base64Val)

	bytes, _ := Hex.Decrypt(base64Val)

	decrypt, err := pattern.NewECB("Hcdipgaf38gFPg1z").SetPadding(padding.NewPKCS7()).Decrypt(bytes)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	println(string(decrypt))
}
