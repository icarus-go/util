package format

import (
	"encoding/json"
	"testing"
)

func TestArray_Format(t *testing.T) {
	marshal, _ := json.Marshal([]string{"1", "2", "3"})

	s := struct {
		Temp string `json:"temp"`
	}{}

	print(string(marshal))

	s.Temp = string(marshal)

	marshal, _ = json.Marshal(s.Temp)

	print(string(marshal))
}
