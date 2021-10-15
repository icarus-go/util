package date

import (
	"testing"
	"time"
)

func Test_generate_TodayEndNight(t *testing.T) {
	date := New()

	val, ok := date.Morning().String()

	if !ok {
		if err := date.Problem(); err != nil {
			t.Errorf(err.Error())
		}

	}

	println(val)

	result, err := time.ParseInLocation(YMDHms.Value(), val, time.Local)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	t.Logf("val = %s, timestamp = %d", val, result.Unix())

	unix, ok := date.SetLocation(Local).SetFormat(YMDHms).ParseInLocation(val).Time()
	if !ok {
		if err := date.Problem(); err != nil {
			t.Errorf(err.Error())
		}
	}

	println(unix.Unix())

}

func Test_generate_TodayMorning(t *testing.T) {

}
