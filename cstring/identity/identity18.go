package identity

import (
	"errors"
	"pmo-test4.yz-intelligence.com/base/utils/cstring/constant"
	"regexp"
	"strconv"
)

type mainland18 struct{}

// Mainland18 大陆18位身份证
var (
	Mainland18 = new(mainland18)
	POWER      = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
)

// IsLocalValidIDCard
//  Author: Kevin·CC
//  Description: 是否是有效的身份证
//  Param idCard
//  Return bool
func (m mainland18) IsLocalValidIDCard(idCard string) bool {
	if _, ok := constant.ProvinceCode[idCard[0:2]]; !ok {
		return false
	}

	birthdate, err := m.Birthday(idCard)
	if err != nil {
		return false
	} // 身份证生日获取

	if !Birthday.IsValid(birthdate) {
		return false
	} // 是否是有效的生日数字

	last17Code := idCard[0:17]
	matched, err := regexp.Match(constant.Numbers.Value(), []byte(last17Code))
	if err != nil {
		return false
	}
	if matched {
		sum := GetPowerSum(last17Code)

		return GetCheckCode(sum) == idCard[17:]
	}

	return false
}

// Age
//  Author: Kevin·CC
//  Description: 年龄
//  Param idCard 身份证
//  Return int 年龄 (实岁)
//  Return error 错误信息
func (m mainland18) Age(idCard string) (int, error) {
	birthdate, err := m.Birthday(idCard)
	if err != nil {
		return 0, errors.New("身份证获取生日解析失败")
	}

	parse, err := Birthday.Parse(birthdate)
	if err != nil {
		return 0, err
	}

	return Birthday.Age(parse), nil
}

// Gender
//  Author: Kevin·CC
//  Description: 获取性别
//  Param idCard 身份证
//  Return int 1=男,2=女
//  Return bool 是否获取成功
func (m mainland18) Gender(idCard string) (int, error) {
	if len(idCard) != 18 {
		return -1, errors.New("身份证长度不正确")
	}
	v, err := strconv.Atoi(string(idCard[16]))
	if err != nil {
		return 0, errors.New("获取身份证第17位失败")
	}
	return v, nil
}

// GenderCNName
//  Author: Kevin·Cai
//  Description: 性别中文名
//  Param gender 性别int值
//  Return string CN名称
//  Return error 错误
func (m mainland18) GenderCNName(gender int) string {
	if gender%2 == 0 {
		return "女"
	}
	return "男"
}

// Birthday
//  Author: Kevin·CC
//  Description: 获取生日
//  Param idCard 身份证
//  Return string 生日
func (mainland18) Birthday(idCard string) (string, error) {
	birthday := idCard[6:14]
	match, err := regexp.Match(constant.Birthday.Value(), []byte(birthday))
	if err != nil {
		return "", err
	}
	if !match {
		return "", errors.New("不合法年龄")
	}
	return birthday, nil
}

func GetPowerSum(code string) int {
	sum := 0
	for index, i := range code {
		atoi, _ := strconv.Atoi(string(i))
		sum += atoi * POWER[index]
	}
	return sum
}

func GetCheckCode(sum int) string {
	remainder := sum % 11
	switch remainder {
	case 10:
		return "2"
	case 9:
		return "3"
	case 8:
		return "4"
	case 7:
		return "5"
	case 6:
		return "6"
	case 5:
		return "7"
	case 4:
		return "8"
	case 3:
		return "9"
	case 2:
		return "X"
	case 1:
		return "0"
	case 0:
		return "1"
	default:
		return ""
	}
}
