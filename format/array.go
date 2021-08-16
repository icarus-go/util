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
