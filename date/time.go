package date

import (
	"time"
)

type Generate struct {
	Error error

	result *result

	need Need
}

func Now() *Generate {
	i := new(Generate)

	i.need.dateFormat = format{
		dateType: Ymd,
		isSet:    false,
	}

	i.need.location = Local

	i.result = &result{}

	now := time.Now()

	i.result.time = &now

	return i
}

// Need 所需参数
//  Author:  Kevin·CC
func (g *Generate) Need(need Need) *Generate {
	g.need = need
	return g
}

// SetLocation Parse地区
//  Author:  Kevin·CC
func (g *Generate) SetLocation(local Location) *Generate {
	g.need.location = local
	return g
}

// SetFormat 设置日期格式
//  Author:  Kevin·CC
func (g *Generate) SetFormat(format Format) *Generate {
	g.need.Format(format)
	return g
}

// Morning 今天的第一秒
//  Author:  Kevin·CC
func (g *Generate) Morning() *Generate {
	val := OneDayMorning.Join(g.result.time.In(g.need.GetLocation()).Format(g.need.dateFormat.dateType.Value()))

	g.result.stringVal = &val

	g.result.time = nil

	return g
}

// EndNight 今天最后一秒
//  Author:  Kevin·CC
func (g *Generate) EndNight() *Generate {

	val := OneDayEndNight.Join(g.result.time.Format(g.need.dateFormat.dateType.Value()))

	g.result.stringVal = &val

	g.result.time = nil

	return g
}

// Parse 根据设置的地区解析
//  Author:  Kevin·CC
func (g *Generate) Parse(date string) *Generate {

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

// AnyTimeMorning 任何一天的第一秒
//  Author:  Kevin·CC
func (g *Generate) AnyTimeMorning(t time.Time) *Generate {
	val := t.Format(Ymd.Value())
	parse, err := time.ParseInLocation(YMDHms.Value(), OneDayMorning.Join(val), g.need.GetLocation())
	g.Error = err
	g.result.time = &parse
	return g
}

// AnyTimeEndNight 任何一天的最后一秒
//  Author:  Kevin·CC
func (g *Generate) AnyTimeEndNight(t time.Time) *Generate {
	val := t.Format(Ymd.Value())
	parse, err := time.ParseInLocation(YMDHms.Value(), OneDayEndNight.Join(val), g.need.GetLocation())
	g.Error = err
	g.result.time = &parse
	return g
}

// AnyStringMorning 任何日期字符串的第一秒
//  Author:  Kevin·CC
func (g *Generate) AnyStringMorning(val string) *Generate {
	g.Parse(val)

	return g.Format()
}

// AnyStringEndNight 任何日期字符串的最后一秒
//  Author:  Kevin·CC
func (g *Generate) AnyStringEndNight(val string) *Generate {
	timeVal := g.Parse(val).TimeValue()

	return g.AnyTimeEndNight(timeVal)
}

// AnyTimestampMorning 任何时间的一天最后一秒
//  Author:  Kevin·CC
func (g *Generate) AnyTimestampMorning(timestamp, nanosecond int64) *Generate {
	return g.AnyTimeMorning(time.Unix(timestamp, nanosecond))
}

// AnyTimestampEndNight 任何时间戳的第一秒
//  Author:  Kevin·CC
func (g *Generate) AnyTimestampEndNight(timestamp int64, nanosecond int64) *Generate {
	return g.AnyTimeMorning(time.Unix(timestamp, nanosecond))
}

// Format 根据SetFormat的日期格式化格式日期
//  Author:  Kevin·CC
func (g *Generate) Format() *Generate {
	val := g.result.time.Format(g.need.dateFormat.dateType.Value())

	g.result.time = nil

	g.result.stringVal = &val
	return g
}

// Add 添加 time.Duration
//  Author:  Kevin·CC
func (g *Generate) Add(d time.Duration) *Generate {
	after := g.result.time.Add(d)

	g.result.time = &after

	return g
}
