package date

import (
	"time"
)

type generate struct {
	Error error

	result *result

	need Need
}

func New() *generate {
	i := new(generate)
	i.need.dateFormat = format{
		dateType: Ymd,
		isSet:    false,
	}
	i.need.location = Local
	i.result = &result{}
	return i
}

// Need 所需参数
//  Author:  Kevin·CC
func (g *generate) Need(need Need) *generate {
	g.need = need
	return g
}

// SetLocation Parse地区
//  Author:  Kevin·CC
func (g *generate) SetLocation(local Location) *generate {
	g.need.location = local
	return g
}

// SetFormat 设置日期格式
//  Author:  Kevin·CC
func (g *generate) SetFormat(format Format) *generate {
	g.need.Format(format)
	return g
}

// Morning 今天的第一秒
//  Author:  Kevin·CC
func (g *generate) Morning() *generate {
	if !g.need.dateFormat.isSet {
		g.need.dateFormat.dateType = Ymd
	}

	val := OneDayMorning.Join(time.Now().In(g.need.GetLocation()).Format(g.need.dateFormat.dateType.Value()))
	g.result.stringVal = &val
	return g
}

// Parse 根据设置的地区解析
//  Author:  Kevin·CC
func (g *generate) Parse(date string) *generate {
	val, err := time.ParseInLocation(g.need.dateFormat.dateType.Value(), date, g.need.GetLocation())
	if err != nil {
		g.Error = err
		return g
	}

	if err == nil {
		g.result.time = &val
	}

	return g
}

// EndNight 今天最后一秒
//  Author:  Kevin·CC
func (g *generate) EndNight() *generate {
	val := OneDayEndNight.Join(time.Now().Format(Ymd.Value()))
	g.result.stringVal = &val
	return g
}

// AnyTimeMorning 任何一天的第一秒
//  Author:  Kevin·CC
func (g *generate) AnyTimeMorning(t time.Time) *generate {
	val := t.Format(Ymd.Value())
	parse, err := time.ParseInLocation(YMDHms.Value(), OneDayMorning.Join(val), g.need.GetLocation())
	g.Error = err
	g.result.time = &parse
	return g
}

// AnyTimeEndNight 任何一天的最后一秒
//  Author:  Kevin·CC
func (g *generate) AnyTimeEndNight(t time.Time) *generate {
	val := t.Format(Ymd.Value())
	parse, err := time.ParseInLocation(YMDHms.Value(), OneDayEndNight.Join(val), g.need.GetLocation())
	g.Error = err
	g.result.time = &parse
	return g
}

// AnyStringMorning 任何日期字符串的第一秒
//  Author:  Kevin·CC
func (g *generate) AnyStringMorning(val string) *generate {
	t := Zero
	return g.Parse(val).Time(&t).AnyTimeMorning(t)
}

// AnyStringEndNight 任何日期字符串的最后一秒
//  Author:  Kevin·CC
func (g *generate) AnyStringEndNight(val string) *generate {
	t := Zero
	return g.Parse(val).Time(&t).AnyTimeEndNight(t)
}

// AnyTimestampMorning 任何时间的一天最后一秒
//  Author:  Kevin·CC
func (g *generate) AnyTimestampMorning(timestamp, nanosecond int64) *generate {
	return g.AnyTimeMorning(time.Unix(timestamp, nanosecond))
}

// AnyTimestampEndNight 任何时间戳的第一秒
//  Author:  Kevin·CC
func (g *generate) AnyTimestampEndNight(timestamp int64, nanosecond int64) *generate {
	return g.AnyTimeMorning(time.Unix(timestamp, nanosecond))
}

// Format 根据SetFormat的日期格式化格式日期
//  Author:  Kevin·CC
func (g *generate) Format(t time.Time) *generate {
	val := t.Format(g.need.dateFormat.dateType.Value())
	g.result.stringVal = &val
	return g
}

// Add 添加 time.Duration
//  Author:  Kevin·CC
func (g *generate) Add(t time.Time, d time.Duration) *generate {
	after := t.Add(d)

	g.result.time = &after

	return g
}
