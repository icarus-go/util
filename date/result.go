package date

import "time"

type result struct {
	stringVal *string
	timestamp *int64
	time      *time.Time
}

// Timestamp 时间戳类型结果
//  Author:  Kevin·CC
func (g *generate) Timestamp(timestamp *int64) *generate {
	if g.result.timestamp != nil {
		*timestamp = *g.result.timestamp
		g.result.timestamp = nil
	}

	if g.result.time != nil {
		*timestamp = g.result.time.Unix()
		g.result.timestamp = nil
	}

	if g.result.stringVal != nil {
		g.Parse(*g.result.stringVal)
		g.result.stringVal = nil
		g.Timestamp(timestamp)
	}

	return g
}

// Time 时间类型结果
//  Author:  Kevin·CC
func (g *generate) Time(val *time.Time) *generate {
	if g.result.time != nil {
		*val = *g.result.time
		g.result.time = nil
	}
	if g.result.timestamp != nil {
		unix := time.Unix(*g.result.timestamp, 0)
		*val = unix
		g.result.timestamp = nil
	}

	if g.result.stringVal != nil {
		g.Parse(*g.result.stringVal)
		g.result.stringVal = nil
		g.Time(val)
		return g
	}

	return g
}

// String 字符串类型结果
//  Author:  Kevin·CC
func (g *generate) String(val *string) *generate {
	if g.result.stringVal != nil && *g.result.stringVal != "" {
		*val = *g.result.stringVal
		g.result.stringVal = nil
		return g
	}
	if g.result.time != nil && !(*g.result.time).IsZero() {
		g.Format(*g.result.time)
		g.result.time = nil
		g.String(val)
	}

	if g.result.timestamp != nil && *g.result.timestamp != 0 {
		g.Format(time.Unix(*g.result.timestamp, 0))
		g.result.timestamp = nil
		g.String(val)
	}

	return g
}
