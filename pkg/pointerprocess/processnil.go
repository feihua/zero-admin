package pointerprocess

import "reflect"

// DefaltData
/*
Author: xixigh <1746877224@qq.com>
Date: 2024/6/27 18:13
*/
func DefaltData(v interface{}) interface{} {
	// 使用反射检查传入的值是否是一个指针
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if !val.IsNil() {
			// 返回指针指向的值
			return val.Elem().Interface()
		}
		// 返回指针类型的零值
		return reflect.Zero(val.Type().Elem()).Interface()
	}
	// 如果不是指针类型，直接返回零值
	return reflect.Zero(val.Type()).Interface()
}
