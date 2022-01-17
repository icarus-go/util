package identity

import (
	"github.com/icarus-go/utils/date"
	"time"
)

type birthday struct{}

var Birthday = new(birthday)

func (b *birthday) IsValid(birthday string) bool {

	currentYear := time.Now().Year() // 当前年份

	parse, err := b.Parse(birthday)
	if err != nil {
		return false
	}

	year := parse.Year()
	if year < 1900 || year > currentYear {
		return false
	}

	month := parse.Month()
	if month < 1 || month > 12 {
		return false
	}

	day := parse.Day()
	if day > 31 || day < 1 {
		return false
	}

	if day == 31 && (month == 4 || month == 6 || month == 9 || month == 11) {
		return false
	}
	if month == 2 {
		return day < 29 || (day == 29 && date.IsLeapYear(year))
	}

	return true
}

// Parse
//  Description: 解析身份证中的生日
//  Author: Kevin·CC
//  Param: birthday
//  return time.Time
func (*birthday) Parse(birthday string) (time.Time, error) {
	if len(birthday) == 6 {
		birthday = "19" + birthday
	} // 1代身份证,将一代处理为2代格式

	birthdayTime, err := time.ParseInLocation(date.YMDNumber.Value(), birthday, time.Local)
	if err != nil {
		return time.Now(), err
	}

	return birthdayTime, nil
}

// Age
//  Author: Kevin·CC
//  Description: 获取年龄
//  Param birthdate 生日日期
//  Return int 年龄
func (*birthday) Age(birthdate time.Time) int {
	return time.Now().Year() - birthdate.Year()
}
