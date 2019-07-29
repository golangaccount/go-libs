package reflect

import (
	"reflect"
)

const EQUALMETHOD = "Equal"

func Equal(v1, v2 interface{}) bool {
	return EqualValue(reflect.ValueOf(v1), reflect.ValueOf(v2))
}

func EqualValue(v1, v2 reflect.Value) bool {
	if v1.Kind() != v2.Kind() {
		return false
	}
	switch v1.Kind() {
	case reflect.Bool:
		return v1.Bool() == v2.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v1.Int() == v2.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v1.Uint() == v2.Uint()
	case reflect.Float32, reflect.Float64:
		return v1.Float() == v2.Float()
	case reflect.Complex64, reflect.Complex128:
		return v1.Complex() == v2.Complex()
	case reflect.Array, reflect.Slice: //集合类型，首先进行程度和集合类型检测，一致进行后续判断
		if v1.Len() != v2.Len() || v1.Type().Elem() != v2.Type().Elem() {
			return false
		}
		if checkEqualFunc(v1, v2) {
			return true
		}
		for i := 0; i < v1.Len(); i++ {
			if !EqualValue(v1.Index(i), v2.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Chan:
		panic(Err_Equal_UnsupportType)
	case reflect.Func:
		return v1.Type() == v2.Type()
	case reflect.Interface:
		v := v2.Convert(v1.Type())
		if v.IsNil() {
			return false
		}
		return true
	case reflect.Map:
		if v1.Len() != v2.Len() || v1.Type().Key() != v2.Type().Key() || v1.Type().Elem() != v2.Type().Elem() {
			return false
		}
		if checkEqualFunc(v1, v2) {
			return true
		}
		keys := v1.MapKeys()
		for _, key := range keys {
			v1item := v1.MapIndex(key)
			v2item := v2.MapIndex(key)
			if !v2item.IsValid() {
				return false
			}
			if !EqualValue(v1item, v2item) {
				return false
			}
		}
		return true
	case reflect.Ptr:
		if v1.Pointer() == v2.Pointer() {
			return true
		} else if v1.Elem().Type() == v2.Elem().Type() {
			if checkEqualFunc(v1, v2) {
				return true
			} else {
				return EqualValue(v1.Elem(), v2.Elem())
			}
		}
	case reflect.String:
		return v1.String() == v2.String()
	case reflect.Struct:
		if v1.Type() != v2.Type() {
			return false
		}
		if checkEqualFunc(v1, v2) {
			return true
		}
		numfile := v1.NumField()
		for i := 0; i < numfile; i++ {
			if !EqualValue(v1.Field(i), v2.Field(i)) {
				return false
			}
		}
		return true
	case reflect.UnsafePointer:
		panic(Err_Equal_UnsupportType)
	}
	return false
}

func checkEqualFunc(v1, v2 reflect.Value) (outputresult bool) {
	defer func() {
		if err := recover(); err != nil {
			outputresult = false
		}
	}()
	if f, has := v1.Type().MethodByName(EQUALMETHOD); has {
		inlen := f.Type.NumIn()
		outlen := f.Type.NumOut()
		if inlen != 2 || outlen != 1 {
			return false
		}
		if f.Type.In(1) != v1.Type() || f.Type.Out(0).Kind() != reflect.Bool {
			return false
		}
		result := f.Func.Call([]reflect.Value{v1, v2})
		return result[0].Bool()
	} else {
		return false
	}
}

func EqualValueMust(v1, v2 reflect.Value) bool {
	defer func() {}()
	return EqualValue(v1, v2)
}

func EqualMust(v1, v2 interface{}) bool {
	defer func() {}()
	return Equal(v1, v2)
}
