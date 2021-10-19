package optimization

import "time"

type Timer struct {
	Fn func()
}

func (t *Timer) Calc() float64 {
	now := time.Now()

	t.Fn()

	end := time.Now()

	return float64(end.UnixNano()-now.UnixNano()) / float64(time.Second)
}
