package cstring

import (
	"html/template"
	"strings"
)

type injection struct{}

// Injection 防止SQL注入
var Injection = new(injection)

// Filter
//  Description: 过滤单个字符串值
//  Author: Kevin·CC
//  Param: value 值
//  return string 过滤后的值
func (injection) Filter(value string) string {
	value = strings.TrimSpace(value) // 过滤空格

	if value == "" {
		return ""
	} // 如果是空值

	value = template.HTMLEscapeString(value) // HTML转义

	return strings.ReplaceAll(value, "'", "''")
}

// Filters 批量过滤
//  Description: 字符串组批量过滤
//  Author: Kevin·CC
//  Param: values 字符串组值
//  return []string 过滤后的值
func (i injection) Filters(values ...string) []string {
	for index := 0; index < len(values); index++ {
		values[index] = i.Filter(values[index])
	}
	return values
}
