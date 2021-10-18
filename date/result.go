package date

import "time"

type result struct {
	stringVal *string
	timestamp *int64
	time      *time.Time
}

// Timestamp 时间戳类型结果
//  Author:  Kevin·CC
func (g *Generate) Timestamp() int64 {
	if g.result.timestamp != nil {
		return *g.result.timestamp
	}

	if g.result.time != nil {
		return g.result.time.Unix()
	}

	if g.result.stringVal != nil {
		g.Parse(*g.result.stringVal)

		g.result.stringVal = nil

		return g.result.time.Unix()
	}

	return 0
}

// TimeValue 时间类型结果
//  Author:  Kevin·CC
func (g *Generate) TimeValue() time.Time {
	if g.result.time != nil {
		return *g.result.time
	}

	if g.result.timestamp != nil {
		return time.Unix(*g.result.timestamp, 0)
	}

	if g.result.stringVal != nil {
		g.Parse(*g.result.stringVal)

		g.result.stringVal = nil

		return *g.result.time
	}

	return Zero
}

// String 字符串类型结果
//  Author:  Kevin·CC
func (g *Generate) String() string {
	if g.result.stringVal != nil && *g.result.stringVal != "" {
		return *g.result.stringVal

	}

	if g.result.time != nil && !(*g.result.time).IsZero() {
		g.Format()

		g.result.time = nil

		return *g.result.stringVal
	}

	if g.result.timestamp != nil && *g.result.timestamp != 0 {
		unix := time.Unix(*g.result.timestamp, 0)

		g.result.time = &unix

		g.result.timestamp = nil

		g.Format()

		result := *g.result.stringVal

		g.result.stringVal = nil

		return result
	}

	return ""
}
