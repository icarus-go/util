package optimization

import (
	"fmt"
	"testing"
	"time"
)

func Test_timer_calc(t *testing.T) {
	timer := new(Timer)

	timer.Fn = func() {
		println("HelloWorld")
		time.Sleep(5*time.Second + 51*time.Millisecond)
	}

	calc := timer.Calc()

	fmt.Printf("共计花费: %.2f 秒", calc)
}
