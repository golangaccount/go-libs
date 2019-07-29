package gtype

import (
	"reflect"
)

var _bool bool
var _int int
var _int8 int8
var _int16 int16
var _int32 int32
var _int64 int64
var _uint uint
var _uint8 uint8
var _uint16 uint16
var _uint32 uint32
var _uint64 uint64
var _float32 float32
var _float64 float64
var _complex64 complex64
var _complex128 complex128
var _string string

func IsDefault(v interface{}) bool {
	kind := reflect.TypeOf(v)
	switch kind.Kind() {
	case reflect.Bool:
		return _bool == v.(bool)
	case reflect.Int:
		return _int == v.(int)
	case reflect.Int8:
		return _int8 == v.(int8)
	case reflect.Int16:
		return _int16 == v.(int16)
	case reflect.Int32:
		return _int32 == v.(int32)
	case reflect.Int64:
		return _int64 == v.(int64)
	case reflect.Uint:
		return _uint == v.(uint)
	case reflect.Uint8:
		return _uint8 == v.(uint8)
	case reflect.Uint16:
		return _uint16 == v.(uint16)
	case reflect.Uint32:
		return _uint32 == v.(uint32)
	case reflect.Uint64:
		return _uint64 == v.(uint64)
	case reflect.Float32:
		return _float32 == v.(float32)
	case reflect.Float64:
		return _float64 == v.(float64)
	case reflect.Complex64:
		return _complex64 == v.(complex64)
	case reflect.Complex128:
		return _complex128 == v.(complex128)
	case reflect.String:
		return _string == v.(string)
	default:
		return false
	}
}
