package identity

import (
	"pmo-test4.yz-intelligence.com/base/utils/cstring/constant"
	"regexp"
	"strconv"
	"strings"
)

type GangAoTai struct{}

func (g GangAoTai) IsLocalValidIDCard(idCard string) bool {
	if idCard == "" {
		return false
	}

	card := strings.ReplaceAll(idCard, "[()]", "")

	length := len(card)
	if length != 8 && length != 9 && length != 10 {
		return false
	}
	idCardBytes := []byte(idCard)
	match, err := regexp.Match("^[a-zA-Z][0-9]{9}$", idCardBytes)
	if err != nil {
		return false
	}

	var info []string

	if match {
		info = g.GetInfo(idCard)
	} // 台湾

	match, err = regexp.Match("^[157][0-9]{6}\\(?[0-9A-Z]\\)?$", idCardBytes)
	if err != nil {
		return false
	}

	if match {
		info[0] = "澳门"
		info[1] = "N"
		info[2] = "true"
	} // 香港

	return false
}

func (g GangAoTai) GetInfo(idCard string) []string {

	var info = make([]string, 3)
	info[0] = "台湾"
	firstCode := string(idCard[1])
	if "1" == firstCode {
		info[1] = "M"
	} else if "2" == firstCode {
		info[1] = "F"
	} else {
		info[1] = "N"
		info[2] = "false"
	}

	if g.IsValidTaiWan(idCard) {
		info[2] = "true"
	} else {
		info[2] = "false"
	}

	return info
}

func (GangAoTai) IsValidTaiWan(idCard string) bool {
	if len(idCard) != 10 {
		return false
	}
	secondCode := idCard[0]
	iStart, ok := constant.TaiWanFirstCode[string(secondCode)]
	if !ok {
		return false
	}

	sum := iStart/10 + (iStart%10)*9

	mid := idCard[1:9]

	iFlag := 8
	for _, i := range mid {
		atoi, err := strconv.Atoi(string(i))
		if err != nil {
			return false
		}

		sum += atoi * iFlag
		iFlag--
	}

	end, err := strconv.Atoi(idCard[9:10])
	if err != nil {
		return false
	}

	if sum%10 == 0 {
		return 0 == end
	}
	return 10-sum%10 == end
}
