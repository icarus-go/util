package assertion

import (
	"pmo-test4.yz-intelligence.com/kit/data/resource"
	"testing"
	"time"
)

func Test_Assertion_ThrowUnsertValue(t *testing.T) {
	current := resource.NewDatetimeByTime(time.Now())
	if err := ThrowUnsetValue(current, "时间未设置"); err != nil {
		t.Fatal(err.Error())
		return
	} else {
		t.Log("datetime success")
	}
}
