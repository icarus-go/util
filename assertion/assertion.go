package assertion

import (
	"errors"
	"reflect"
)

// IsUnset
//  Author: Kevin·Cai
//  Description: 是否是初始值
//  Param value 值
//  Return bool 是否是初始化值
func IsUnset(value interface{}) bool {
	reflectValue := reflect.ValueOf(value)
	switch reflectValue.Kind() {
	case reflect.String:
		return reflectValue.Len() == 0
	case reflect.Bool:
		return !reflectValue.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflectValue.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return reflectValue.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return reflectValue.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return reflectValue.IsNil()
	}
	return reflect.DeepEqual(reflectValue.Interface(), reflect.Zero(reflectValue.Type()).Interface())
}

// ThrowUnsetValue 抛出未设置值得错误
//  Author: Kevin·Cai
//  Description: 是否未设置初始值
//  Param value 值
//  Param msg 消息内容
//  Return error 当为空时,返回的错误消息
func ThrowUnsetValue(value interface{}, msg string) error {
	if IsUnset(value) {
		if msg == "" {
			return errors.New("值未设置为初始值")
		}
		return errors.New(msg)
	}
	return nil
}
