package date

import (
	"testing"
	"time"
)

func Test_date_Morning(t *testing.T) {
	value := ""
	times := Zero

	New().SetFormat(YmdCN).Morning().String(&value).SetFormat(YMDHmsCN).SetLocation(CST).Parse(value).Time(&times)

	println(value)

	print(times.Unix())
}

func Test_date_MorningByTime(t *testing.T) {
	temp := Zero

	if err := New().SetLocation(ShangHai).AnyTimeMorning(time.Now()).Time(&temp).Error; err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(temp.Unix())

	endResult := ""
	if err := New().AnyTimeEndNight(time.Now()).Time(&temp).SetFormat(YMDHms).Format(temp).String(&endResult).Error; err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("endNight: " + temp.Format(YMDHms.Value()))
	t.Log("endNight: " + endResult)

	s := ""

	println(s)
}

func Test_date_Add(t *testing.T) {
	s := ""
	temp := Zero

	a := ""

	g := New()
	if err := g.Morning().
		String(&s).
		SetFormat(YMDHms).Parse(s).Time(&temp).
		Add(temp, time.Duration(10)*time.Second).String(&a).
		Error; err != nil {
		t.Log(err.Error())
		return
	}

	t.Log(a)
}
