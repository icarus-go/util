package cstring

import (
	"encoding/json"
	"fmt"
	"html/template"
	"testing"
)

func TestMaskAsterisk(t *testing.T) {
	md := struct {
		Phone MaskAsterisk `json:"phone"`
		Name  MaskAsterisk `json:"name"`
		Test  MaskAsterisk `json:"test"`
	}{
		Phone: "440000200001011234",
		Name:  "xjm",
		Test:  "sdf32",
	}
	bytes, err := json.Marshal(md)
	if err != nil {
		t.Errorf("【SUCCESS】 %s", err.Error())
	}
	t.Errorf("【SUCCESS】 %s", string(bytes))
}

// Test_SQL_Injection
//  Description: SQL注入测试以及解决方案
//  Author: Kevin·CC
//  Param: t 测试实例
func Test_SQL_Injection(t *testing.T) {
	sql := "1' or 1=1"

	escapeString := template.HTMLEscapeString(sql)

	print(escapeString)
}

func Test_SQL_Injection_Filter(t *testing.T) {
	println(Injection.Filter("admin`!@@#$%@$#'"))
}

func Test_SQL_Injection_Filters(t *testing.T) {
	fmt.Printf("%v", Injection.Filters("admin\\`!@@#$%@$#'", "李四'"))
}
