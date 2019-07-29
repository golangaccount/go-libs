package collection

import (
	"reflect"

	greflect "github.com/golangaccount/go-libs/reflect"
)

func Index(coll interface{}, item interface{}) int {
	value := reflect.Indirect(reflect.ValueOf(coll))
	itemvalue := reflect.ValueOf(item)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		panic(Err_NotSliceOrArray)
	}
	l := value.Len()
	for i := 0; i < l; i++ {
		if greflect.EqualValueMust(value.Index(i), itemvalue) {
			return i
		}
	}
	return -1
}

func IndexField(coll interface{}, field string, item interface{}) int {
	return -1
}

func IndexFunc(coll interface{}, function func(interface{}, interface{}) bool) int {
	return -1
}
