package date

import (
	"testing"
	"time"
)

func Test_date_Morning(t *testing.T) {
	s := Now().Morning().String()
	timeValue := Now().Morning().SetFormat(YMDHms).TimeValue()
	timestamp := Now().Morning().SetFormat(YMDHms).Timestamp()

	t.Logf("字符串: %s", s)

	t.Logf("time.Time : %d", timeValue.Unix())

	t.Logf("时间戳: %d", timestamp)

}

func Test_date_MorningByTime(t *testing.T) {

	time := time.Now()

	s := Now().AnyTimeMorning(time).String()
	timestamp := Now().AnyTimeMorning(time).Timestamp()
	timeValue := Now().AnyTimeMorning(time).TimeValue()

	t.Logf("字符串: %s", s)

	t.Logf("time.Time : %d", timeValue.Unix())

	t.Logf("时间戳: %d", timestamp)

}

func Test_date_Add(t *testing.T) {

	timestamp := Now().Add(time.Duration(10) + time.Second).Timestamp()

	timeValue := Now().Add(time.Duration(10) + time.Second).TimeValue()

	s := Now().Add(time.Duration(10) + time.Second).SetFormat(YMDHms).String()

	t.Logf("字符串: %s", s)

	t.Logf("time.Time : %d", timeValue.Unix())

	t.Logf("时间戳: %d", timestamp)

}
