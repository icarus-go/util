package identity

import (
	"regexp"
	"strconv"
	"strings"
)

type hongKong struct{}

var HongKong = hongKong{}

func (h hongKong) IsLocalValidIDCard(idCard string) bool {
	idCardBytes := []byte(idCard)
	match, err := regexp.Match("^[A-Z]{1,2}[0-9]{6}\\(?[0-9A]\\)?$", idCardBytes)
	if err != nil {
		return false
	}
	if !match {
		return false
	}

	return h.isValidHKCard(idCard)
}

func (h hongKong) isValidHKCard(idCard string) (ok bool) {
	card := strings.ReplaceAll(idCard, "[()]", "")
	var sum int
	firstCode := strings.ToUpper(string(card[0]))[0]
	if len(idCard) == 9 {
		secondCode := strings.ToUpper(string(card[1]))[0]
		sum = (int(firstCode)-55)*8 + (int(secondCode)-55)*8
	} else {
		sum = 522 + (int(firstCode)-55)*8
	}

	mid := card[1:7]
	end := card[7:8]

	iFlag := 7
	for _, v := range mid {
		atoi, err := strconv.Atoi(string(v))
		if err != nil {
			return
		}
		sum = sum + atoi*iFlag
		iFlag--
	}

	if "A" == strings.ToUpper(end) {
		sum += 10
	} else {
		endInt, err := strconv.Atoi(end)
		if err != nil {
			return
		}
		sum += endInt
	}
	if sum%11 == 0 {
		ok = true
		return
	}
	return

}

func (h hongKong) Birthday(idCard string) (string, error) {
	return "", ErrUnparseBirthday
}

func (h hongKong) Age(idCard string) (int, error) {
	return 0, ErrUnparseAge
}

func (h hongKong) Gender(idCard string) (int, error) {
	return 0, ErrUnparseGender
}
