package cstring

import (
	"encoding/json"
	"testing"
)

func TestMaskAsterisk(t *testing.T) {
	md := struct {
		Phone MaskAsterisk `json:"phone"`
		Name  MaskAsterisk `json:"name"`
		Test  MaskAsterisk `json:"test"`
	}{
		Phone: "11111111111",
		Name:  "xjm",
		Test:  "sdf32",
	}
	bytes, err := json.Marshal(md)
	if err != nil {
		t.Errorf("【SUCCESS】 %s", err.Error())
	}
	t.Errorf("【SUCCESS】 %s", string(bytes))
}
