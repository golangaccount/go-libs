package length

import (
	"encoding/base64"
	"math"
	"reflect"
	"strconv"
	"strings"

	gstrings "github.com/golangaccount/go-libs/strings"
	"github.com/golangaccount/go-libs/valid"
)

func init() {
	valid.RegistValidFunc(LENGTH, length, validParm)
}

func length(value reflect.Value, field reflect.Value, fieldinfo *valid.FieldInfo, ruleinfo *valid.RuleInfo) bool {
	var start, end uint64 = ruleinfo.Tag[lengthStart].(uint64), ruleinfo.Tag[lengthEnd].(uint64)
	switch field.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
		l := uint64(field.Len())
		return l >= start && l <= end
	default:
		panic("length只支持string/array/slice/map/chan类型数据")
	}
}

func validParm(fieldinfo *valid.FieldInfo, ruleinfo *valid.RuleInfo) (bool, string) {
	var start, end uint64
	values := strings.SplitN(ruleinfo.Parm, "...", -1)
	if len(values) == 0 || len(values) > 2 {
		return false, "length的验证规则不许是一个具体的数字或区间，参数形式必须为：[xxxx]或[xxxx...xxxx]或[...xxx]或[xxxx...]"
	}
	if fieldinfo.IsBase64 {
		for i, item := range values {
			bts, err := base64.StdEncoding.DecodeString(item)
			if err != nil {
				return false, "length参数值base64解码错误"
			}
			values[i] = string(bts)
		}
	}
	if len(values) == 1 {
		i, err := strconv.ParseUint(values[0], 10, 32)
		if err != nil {
			return false, "length规则解析失败:" + err.Error()
		}
		start = i
		end = i
	} else {
		if gstrings.IsEmptyOrWhite(values[0]) {
			start = 0
		} else {
			i, err := strconv.ParseUint(values[0], 10, 32)
			if err != nil {
				return false, "length规则解析失败:" + err.Error()
			}
			start = i
		}
		if gstrings.IsEmptyOrWhite(values[1]) {
			end = math.MaxUint32
		} else {
			i, err := strconv.ParseUint(values[1], 10, 32)
			if err != nil {
				return false, "length规则解析失败:" + err.Error()
			}
			end = i
		}
	}

	switch fieldinfo.Field.Type.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
	default:
		return false, "length只支持string/array/slice/map/chan类型数据"
	}
	ruleinfo.Tag[lengthStart] = start
	ruleinfo.Tag[lengthEnd] = end
	return true, ""
}
