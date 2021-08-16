package network

import (
	"go.uber.org/zap"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	result, err := New("https://www.baidu.com").
		QueryString(map[string]string{}).Get().Do()
	if err != nil {
		return
	}
	print(string(result))
}

func TestDo(t *testing.T) {
	_ = NewHolidaySDK().ListByHolidays(2021)
}

type holidaySDK struct {
	client *netWork
}

func NewHolidaySDK() *holidaySDK {
	holidaySDK := new(holidaySDK)
	return holidaySDK
}

//ListByHolidays 获取节假日
func (h *holidaySDK) ListByHolidays(years int) error {
	//https://api.apihubs.cn/holiday/get?field=year,month,date,week,weekend,workday,holiday&year=2021&workday=2&order_by=1&cn=1&size=366
	result, err := New("https://api.apihubs.cn/holiday/get", SetDebug(true)).QueryString(map[string]string{
		"field":    "year,month,date,week,weekend,workday,holiday,holiday_legal",
		"year":     strconv.Itoa(years),
		"workday":  "2",
		"order_by": "1",
		"cn":       "1",
		"size":     "365",
	}).Get().Do()
	if err != nil {
		zap.L().Error(err.Error())
		return nil

	}
	zap.L().Info(string(result))
	return nil
}
