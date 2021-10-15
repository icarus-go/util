package date

import (
	"time"
)

type generate struct {
	location Location
	format   Format
	err      error
	result   interface{}
}

func New() *generate {
	i := new(generate)
	i.location = CST
	return i
}

func (g *generate) SetLocation(local Location) *generate {
	g.location = local
	return g
}

func (g *generate) GetLocation() *time.Location {
	if g.location == CST {
		return time.UTC
	}

	if g.location == UTC {
		return time.UTC
	}

	if g.location == Local {
		return time.Local
	}

	return time.Local
}

// SetFormat 设置日期格式
//  Author:  Kevin·CC
func (g *generate) SetFormat(format Format) *generate {
	g.format = format
	return g
}

func (g *generate) Morning() *generate {
	if g.format != Ymd && g.format != YmdSlash {
		g.format = Ymd
	}

	g.result = time.Now().Format(Ymd.Value()) + " 00:00:00"
	return g
}

func (g *generate) ParseInLocation(date string) *generate {
	val, err := time.ParseInLocation(g.format.Value(), date, g.GetLocation())
	if err != nil {
		g.err = err
		return g
	}
	g.result = val
	return g
}

func (g *generate) EndNight() *generate {
	if g.format != Ymd && g.format != YmdSlash {
		g.format = Ymd
	}

	g.result = time.Now().Format(Ymd.Value()) + " 23:59:59"
	return g
}

// Timestamp
//  Author:  Kevin·CC
func (g *generate) Timestamp() (int64, bool) {
	r, ok := g.result.(int64)
	return r, ok
}

func (g *generate) Time() (time.Time, bool) {
	t, ok := g.result.(time.Time)
	return t, ok
}

// String
//  Author:  Kevin·CC
func (g *generate) String() (string, bool) {
	r, ok := g.result.(string)
	return r, ok
}

func (g *generate) Problem() error {
	return g.err
}
