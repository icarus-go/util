package format

import (
	"reflect"
)

type _struct struct{}

var Struct = new(_struct)

//ToMap 利用反射将结构体转化为map
//@author: [SliverHorn](https://github.com/SliverHorn)
func (_struct) ToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)

	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
