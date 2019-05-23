package reflect

import "reflect"

/*
* 获取指针指向的实际的对象信息
 */
func Indirect(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		return Indirect(v.Elem())
	} else {
		return v
	}
}
