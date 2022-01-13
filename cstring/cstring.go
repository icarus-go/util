package cstring

import "strings"

//IsBlank
//  Author: Kevin·CC
//  Description: 是否是空字符串,如果是返回 `DefaultValue`
//  Param judgment
//  Param defaultValue
//  Return string
func IsBlank(judgment string, defaultValue string) string {
	if judgment == "" {
		return defaultValue
	}
	return judgment
}

// IndexOf
//  Author: Kevin·Cai
//  Description: (content) 是否包含 (include)
//  Param content 内容
//  Param include 要包含的内容
//  Return bool 是否相等
func IndexOf(content string, include string) bool {
	return strings.Index(content, include) > -1
}

// IndexOfs
//  Author: Kevin·Cai
//  Description: 是否包含这一批(includes)中的任意一个
//  Param content 内容
//  Param includes 要包含的集合
//  Return bool 是否包含
func IndexOfs(content string, includes ...string) bool {
	for _, include := range includes {
		if IndexOf(content, include) {
			return true
		}
	}
	return false
}
