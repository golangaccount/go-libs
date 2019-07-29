package collection

import (
	"errors"
	"reflect"

	greflect "github.com/golangaccount/go-libs/reflect"
)

//唯一化结合元素
func Unique(result interface{}, items ...interface{}) error {
	resultvalue := reflect.ValueOf(result)
	ispoint := resultvalue.Kind() == reflect.Ptr
	resultvalue = reflect.Indirect(resultvalue)

	if !ispoint || resultvalue.Kind() != reflect.Slice {
		return errors.New("输出结果必须为指向slice的指针")
	}
	if resultvalue.Len() != 0 {
		return errors.New("用于输出结果的集合必须为空")
	}

	elemtype := resultvalue.Type().Elem()
	for _, item := range items {
		itemvalue := reflect.Indirect(reflect.ValueOf(item))
		if (itemvalue.Kind() == reflect.Slice || itemvalue.Kind() == reflect.Array) && itemvalue.Type().Elem() == elemtype {
			for i := 0; i < itemvalue.Len(); i++ {
				mark := false
				for j := 0; j < resultvalue.Len(); j++ {
					if greflect.EqualValue(itemvalue.Index(i), resultvalue.Index(j)) {
						//resultvalue
						mark = true
						break
					}
				}
				if !mark {
					resultvalue.Set(reflect.Append(resultvalue, itemvalue.Index(i)))
				}
			}
		} else {
			return errors.New("参数错误，必须是相同类型的slice或array")
		}
	}
	return nil
}

//交集
func Intersection() {

}

//并集
func Union() {

}

//补集
func Complement() {

}
