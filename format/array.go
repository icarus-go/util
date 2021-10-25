package format

import (
	"fmt"
	"strings"
)

type _array struct{}

var Array = new(_array)

//ToString 将数组格式化为字符串
//@author: [SliverHorn](https://github.com/SliverHorn)
func (_array) ToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

// InterfacesToUint64s interface{}->uint64
//  Author:  Kevin·CC
func InterfacesToUint64s(slices []interface{}) []uint64 {
	res := make([]uint64, 0, len(slices))

	for _, slice := range slices {
		if val, ok := slice.(uint64); ok {
			res = append(res, val)
		}
	}
	return res
}

// InterfacesToUints []interface{} -> []int
//  Author:  Kevin·CC
func InterfacesToUints(slices []interface{}) []uint {
	res := make([]uint, 0, len(slices))

	for _, slice := range slices {
		if val, ok := slice.(uint); ok {
			res = append(res, val)
		}
	}
	return res
}

// InterfacesToStrings interface{} -> string
//  Author:  Kevin·CC
func InterfacesToStrings(slices []interface{}) []string {
	res := make([]string, 0, len(slices))

	for _, slice := range slices {
		if val, ok := slice.(string); ok {
			res = append(res, val)
		}
	}
	return res
}

// InterfacesToInt64s interface{} -> int64
//  Author:  Kevin·CC
func InterfacesToInt64s(slices []interface{}) []int64 {
	res := make([]int64, 0, len(slices))

	for _, slice := range slices {
		if val, ok := slice.(int64); ok {
			res = append(res, val)
		}
	}
	return res
}
