package identity

import (
	"pmo-test4.yz-intelligence.com/base/utils/cstring/constant"
	"regexp"
)

type mainland15 struct{}

// Mainland15 大陆15位身份证
var Mainland15 = new(mainland15)

func (m mainland15) IsLocalValidIDCard(idCard string) bool {
	match, err := regexp.Match(constant.Numbers.Value(), []byte(idCard))
	if err != nil {
		return false
	}
	if !match {
		return false
	}

	_, ok := constant.ProvinceCode[m.ProvinceCode(idCard)]
	if !ok {
		return false
	}

	birthdate, err := m.Birthday(idCard)
	if err != nil {
		return false
	}

	return Birthday.IsValid(birthdate)
}

// Age
//  Author: Kevin·CC
//  Description: 年龄
//  Param idCard 身份证
//  Return int 年龄
//  Return error 无法解析年龄的错误
func (m mainland15) Age(idCard string) (int, error) {
	birthdate, err := m.Birthday(idCard)
	if err != nil {
		return 0, err
	}

	parse, err := Birthday.Parse(birthdate)
	if err != nil {
		return 0, err
	}

	return Birthday.Age(parse), nil
}

// Birthday
//  Author: Kevin·CC
//  Description: 获取15位身份证的
//  Return string
func (mainland15) Birthday(idCard string) (string, error) {
	birthdate := "19" + string(idCard[6:12])
	match, err := regexp.Match(constant.Birthday.Value(), []byte(birthdate))
	if err != nil || !match {
		return "", ErrUnparseBirthday
	}
	return birthdate, nil
}

// Gender
//  Author: Kevin·CC
//  Description: 性别
//  Param idCard 身份证
//  Return int 0
//  Return error 无法解析性别
func (m mainland15) Gender(idCard string) (int, error) {
	return 0, ErrUnparseGender
}

// ProvinceCode
//  Author: Kevin·CC
//  Description: 获取省份编码
//  Param idCard 身份证
//  Return string 省份编码
func (mainland15) ProvinceCode(idCard string) string {
	return idCard[0:2]
}

// ProvinceName
//  Author: Kevin·CC
//  Description: 获取省份名称
//  Param provinceCode 省份编码
//  Return string 省份名称
//  Return bool 是否存在
func (mainland15) ProvinceName(provinceCode string) (string, bool) {
	name, ok := constant.ProvinceCode[provinceCode]
	return name, ok
}
