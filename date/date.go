package date

import "time"

type Format string

const (
	//YMDHmsCN 中文年月日时分秒
	YMDHmsCN Format = "2006年01月02日 15时04分05秒"
	//YmdCN 中文年月日
	YmdCN Format = "2006年01月02日"
	//YMDHms 年-月-日 时:分:秒
	YMDHms Format = "2006-01-02 15:04:05"
	//Ymd 年-月-日
	Ymd Format = "2006-01-02"
	//YMDHmsSlash 年/月/日 时:分:秒
	YMDHmsSlash Format = "2006/01/02 15:04:05"
	//YmdSlash 年/月/日
	YmdSlash Format = "2006/01/02"
)

//Value 值
func (d Format) Value() string {
	return string(d)
}

type Location string

const (
	CST      Location = "CST"
	ShangHai Location = "Asia/Shanghai"
	UTC      Location = "UTC"
	Local    Location = "Local"
)

func (l Location) Value() string {
	return string(l)
}

var Zero = time.Time{}

type OneTime string

const (
	OneDayMorning OneTime = " 00:00:00"

	OneDayEndNight OneTime = " 23:59:59"
)

func (o OneTime) Value() string {
	return string(o)
}

func (o OneTime) Join(val string) string {
	return val + o.Value()
}

const (
	gregorianCutOverYear       = 1582
	gregorianCutoverYearJulian = 1582
)

func IsLeapYear(year int) bool {
	if year&3 != 0 {
		return false
	}

	//todo 暂未支持Julian
	return (year%100 != 0) || (year%400 == 0) // Gregorian
}
