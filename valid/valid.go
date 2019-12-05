package valid

import (
	"errors"
	"reflect"

	gstrings "github.com/golangaccount/go-libs/strings"
)

type valid interface {
	Valid() (bool, error)
}

//Valid 验证
func Valid(src interface{}) (bool, error) {
	if src == nil {
		return false, nil
	}
	if v, ok := src.(valid); ok {
		return v.Valid()
	}
	value := reflect.ValueOf(src)
	if !value.IsValid() { //用来解决：var a interface{} 的情况
		return false, errors.New("不是有效的数据")
	}
	if value.Kind() == reflect.Interface || value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	typeInfo := getType(value.Type())
	for _, item := range typeInfo.Fields {
		field := value.FieldByName(item.Name)
		for _, subitem := range item.Rules {
			f, has := validFuncCache[subitem.Tp]
			if !has {
				return false, nil
			}
			//value reflect.Value, field reflect.StructField, info RuleInfo

			if !f.Func(value, field, item, subitem) {
				var err error
				if !gstrings.IsEmptyOrWhite(subitem.Prompt) {
					err = errors.New(subitem.Prompt)
				} else if !gstrings.IsEmptyOrWhite(item.Prompt) {
					err = errors.New(item.Prompt)
				} else {
					err = errors.New(item.Name + "验证失败")
				}
				return false, err
			}
		}
	}
	return true, nil
}
