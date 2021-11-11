package identity

import (
	"pmo-test4.yz-intelligence.com/base/utils/cstring/constant"
	"regexp"
	"strconv"
)

type mainland18 struct{}

// Mainland18大陆18位身份证
var Mainland18 = new(mainland18)

func (mainland18) IsLocalValidIDCard(idCard string) bool {
	if _, ok := constant.ProvinceCode[idCard[0:2]]; !ok {
		return false
	}

	birthday := idCard[6:14]
	match, err := regexp.Match(constant.Birthday.Value(), []byte(birthday))
	if err != nil {
		return false
	}
	if !match {
		return false
	}

	if !Birthday.IsValid(birthday) {
		return false
	}

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

var POWER = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

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
