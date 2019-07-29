package reflect

import (
	"reflect"
	"regexp"
)

func Get(s interface{}, field string) (result interface{}, err error) {
	v := reflect.ValueOf(s)
	v = Indirect(v)
	return nil, nil
}

var regexp_key = regexp.MustCompile(`^".+"$`)
var regexp_index = regexp.MustCompile(`"^[0-9]+$"`)
var regepx_field = regexp.MustCompile(`^[A-Z][a-z0-9-]{0,}$`)

func get(value reflect.Value, field string) (result reflect.Value, err error) {
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
	case reflect.Map:
	case reflect.Struct:
	}
	return reflect.Value{}, nil
}
