package cstring

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
