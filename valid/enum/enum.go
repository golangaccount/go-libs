package enum

import (
	"encoding/base64"
	"reflect"
	"strconv"
	"strings"

	gstrings "github.com/golangaccount/go-libs/strings"
	"github.com/golangaccount/go-libs/valid"
)

func init() {
	valid.RegistValidFunc(ENUM, enum, validParm)
}

func validParm(fieldinfo *valid.FieldInfo, ruleinfo *valid.RuleInfo) (bool, string) {
	var arr interface{}
	var err error
	enums := strings.SplitN(ruleinfo.Parm, ",", -1)
	if fieldinfo.IsBase64 {
		for i, item := range enums {
			bts, err := base64.StdEncoding.DecodeString(item)
			if err != nil {
				return false, "enum值base64解码错误"
			}
			enums[i] = string(bts)
		}
	}
	enums = gstrings.Unique(gstrings.RemoveEmpty(trim(enums)))
	if enums = gstrings.RemoveEmpty(enums); len(enums) == 0 {
		return false, "必须要有枚举值"
	}

	switch fieldinfo.Field.Type.Kind() {
	case reflect.String:
		arr, err = enums, nil
	case reflect.Int:
		arr, err = int64Arr(enums, 0)
	case reflect.Int8:
		arr, err = int64Arr(enums, 8)
	case reflect.Int16:
		arr, err = int64Arr(enums, 16)
	case reflect.Int32:
		arr, err = int64Arr(enums, 32)
	case reflect.Int64:
		arr, err = int64Arr(enums, 64)
	case reflect.Uint:
		arr, err = uint64Arr(enums, 0)
	case reflect.Uint8:
		arr, err = uint64Arr(enums, 8)
	case reflect.Uint16:
		arr, err = uint64Arr(enums, 16)
	case reflect.Uint32:
		arr, err = uint64Arr(enums, 32)
	case reflect.Uint64:
		arr, err = uint64Arr(enums, 64)
	default:
		return false, "enum只支持stirng和int或uint形式数据"
	}
	if err != nil {
		return false, "enum类型解析失败" + err.Error()
	}
	ruleinfo.Tag[ENUM] = arr
	return true, ""
}

func trim(src []string) []string {
	for i, item := range src {
		if strings.HasPrefix(item, "\"") && strings.HasSuffix(item, "\"") {
			src[i] = strings.Trim(item, "\"")
		}
	}
	return src
}

func int64Arr(value []string, tp int) ([]int64, error) { //0 8 16 32 64
	result := make([]int64, len(value))
	for i, item := range value {
		v, err := strconv.ParseInt(item, 10, tp)
		if err != nil {
			return nil, err
		}
		result[i] = v
	}
	return result, nil
}
func uint64Arr(value []string, tp int) ([]uint64, error) {
	result := make([]uint64, len(value))
	for i, item := range value {
		v, err := strconv.ParseUint(item, 10, tp)
		if err != nil {
			return nil, err
		}
		result[i] = v
	}
	return result, nil
}

func enum(value reflect.Value, field reflect.Value, fieldinfo *valid.FieldInfo, ruleinfo *valid.RuleInfo) bool {
	switch field.Type().Kind() {
	case reflect.String:
		return enumString(field.String(), ruleinfo)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return enumInt(field.Int(), ruleinfo)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return enumUint(field.Uint(), ruleinfo)
	default:
		panic("不支持的数据类型")
	}
}

func enumString(value string, info *valid.RuleInfo) bool {
	enums := info.Tag[ENUM].([]string)
	for _, item := range enums {
		if item == value {
			return true
		}
	}
	return false
}

func enumInt(value int64, info *valid.RuleInfo) bool {
	enums := info.Tag[ENUM].([]int64)
	for _, item := range enums {
		if item == value {
			return true
		}
	}
	return false
}

func enumUint(value uint64, info *valid.RuleInfo) bool {
	enums := info.Tag[ENUM].([]uint64)
	for _, item := range enums {
		if item == value {
			return true
		}
	}
	return false
}
