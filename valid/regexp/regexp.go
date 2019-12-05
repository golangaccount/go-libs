package regexp

import (
	"encoding/base64"
	"reflect"
	"regexp"

	"github.com/golangaccount/go-libs/valid"
)

func init() {
	valid.RegistValidFunc(REGEXP, validFunc, validParm)
	valid.RegistValidFunc(EMAIL, validFunc, nil)
	valid.RegistValidFunc(URL, validFunc, nil)
	valid.RegistValidFunc(MOBILPHONE, validFunc, nil)
}

func validParm(fieldinfo *valid.FieldInfo, ruleinfo *valid.RuleInfo) (bool, string) {
	var regstr = ruleinfo.Parm
	if fieldinfo.IsBase64 {
		bts, err := base64.StdEncoding.DecodeString(regstr)
		if err != nil {
			return false, "regexp参数值base64解码错误"
		}
		regstr = string(bts)
	}
	reg, err := regexp.Compile(regstr)
	if err != nil {
		return false, "regexp规则解析失败:" + err.Error()
	}
	ruleinfo.Tag[REGEXP] = reg
	return true, ""
}

func validFunc(value reflect.Value, field reflect.Value, fieldinfo *valid.FieldInfo, ruleinfo *valid.RuleInfo) bool {
	var reg *regexp.Regexp
	switch ruleinfo.Tp {
	case REGEXP:
		reg = ruleinfo.Tag[REGEXP].(*regexp.Regexp)
	case EMAIL:
		reg = EMAILREG
	case URL:
		reg = URLREG
	case MOBILPHONE:
		reg = MOBILPHONEREG
	default:
		panic("不支持的规则类型")
	}
	if field.Kind() == reflect.String {
		return reg.MatchString(field.String())
	}
	panic(ruleinfo.Tp + "只支持string类型")
}
