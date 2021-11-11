package identity

import (
	"pmo-test4.yz-intelligence.com/base/utils/cstring/constant"
	"regexp"
)

var POWER = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

type Mainland15 struct{}

func (Mainland15) IsLocalValidIDCard(idCard string) bool {
	match, err := regexp.Match(constant.Numbers.Value(), []byte(idCard))
	if err != nil {
		return false
	}
	if !match {
		return false
	}

	_, ok := constant.ProvinceCode[idCard[0:2]]
	if !ok {
		return false
	}

	birthday := "19" + idCard[6:12]

	match, err = regexp.Match(constant.Birthday.Value(), []byte(birthday))
	if err != nil {
		return false
	}
	if !match {
		return false
	}

	if Birthday.IsValid(birthday) {
		return true
	}

	return false
}
