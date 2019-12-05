package valid

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	gstrings "github.com/golangaccount/go-libs/strings"
)

//TypeInfo 类型信息
type TypeInfo struct {
	Tp     reflect.Type //类型信息
	Fields []*FieldInfo //类型中需要进行验证的字段信息
}

//FieldInfo 字段信息
type FieldInfo struct {
	Name     string              //字段名称
	Field    reflect.StructField //字段信息
	Rules    []*RuleInfo         //验证规则
	Prompt   string              //错误提示信息
	IsBase64 bool                //是否使用base64对参数进行编码，只要有对应的tag不管有不有值都默认使用base64编码(一般不需要，主要是用来处理一些特殊字符)
}

//RuleInfo 规则信息
type RuleInfo struct {
	Tp     string                 //验证规则类型
	Parm   string                 //验证规则参数
	Prompt string                 //验证规则提示信息
	Tag    map[string]interface{} //验证规则对应的package
}

//类型信息缓存，进行验证时，直接从缓存中获取信息，不再进行计算
var typeCache = struct {
	rw    sync.RWMutex
	cache map[reflect.Type]TypeInfo
}{sync.RWMutex{}, map[reflect.Type]TypeInfo{}}

//RegistTypeInfo 注册类型的验证信息
//tp：数据结构类型
//tpInfo:数据结构验证信息
//通过parse函数进行解析的会进行相关的验证，如果不使用该函数，自定义对应的数据信息，请按照各个验证规则实现相应的数据
func RegistTypeInfo(tp reflect.Type, tpInfo TypeInfo) {
	typeCache.rw.Lock()
	defer typeCache.rw.Unlock()
	if _, has := typeCache.cache[tp]; has {
		return
	}
	typeCache.cache[tp] = tpInfo
}

//Parse 计算struct的valid规则信息
func Parse(tp reflect.Type) TypeInfo {
	if tp.Kind() != reflect.Struct {
		panic("必须是struct结构类型")
	}
	v := TypeInfo{}
	v.Tp = tp
	v.Fields = make([]*FieldInfo, 0)
	length := tp.NumField()
	for i := 0; i < length; i++ {
		f := tp.Field(i)
		if !isExport(f.Name) {
			continue
		}
		rules := parse(f.Tag.Get(tagName))
		fieldinfo := &FieldInfo{Name: f.Name, Prompt: f.Tag.Get(tagPrompt), Field: f}
		if _, has := f.Tag.Lookup(tagBase64); has {
			fieldinfo.IsBase64 = true
		}
		if len(rules) > 0 {
			for _, item := range rules {
				if v, has := validFuncCache[item.Tp]; !has {
					panic("没有" + item.Tp + "对应的验证方法")
				} else if v.Valid != nil {
					if pass, info := v.Valid(fieldinfo, item); !pass {
						panic(fmt.Sprintf("%s类型%s字段%s验证规则设置的参数%s验证错误;%s", tp.Name(), f.Name, item.Tp, item.Parm, info))
					}
				}
			}
			fieldinfo.Rules = rules
			v.Fields = append(v.Fields, fieldinfo)
		}
	}
	return v
}

func getType(tp reflect.Type) TypeInfo {
	typeCache.rw.Lock()
	defer typeCache.rw.Unlock()
	if v, has := typeCache.cache[tp]; has {
		return v
	}
	v := Parse(tp)
	typeCache.cache[tp] = v
	return v
}

func isExport(name string) bool {
	ch, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(ch)
}

//validtype[parm];validtype[parm]
func parse(tag string) []*RuleInfo {
	if gstrings.IsEmptyOrWhite(tag) {
		return nil
	}
	runes := []rune(tag)
	length := len(runes)
	result := make([]*RuleInfo, 0)
	index := 0
	tp := 1           //1:tp 2:parm
	isclose := 0      //1表示打开 2表示关闭
	value := []rune{} //value
	rule := &RuleInfo{
		Tag: make(map[string]interface{}, 0),
	}
	for {
		if index == length {
			break
		}
		switch runes[index] {
		case ':': //使用:作为转义符 ::=: :[=[ :]=] :;=; :x=x
			index++
			value = append(value, runes[index])
		case '[':
			tp = 2
			isclose = 1
			rule.Tp = string(value)
			value = []rune{}
		case ']':
			tp = 0
			isclose = 2
			rule.Parm = string(value)
			value = []rune{}
		case ';':
			if isclose == 1 {
				panic("[]必须成对出现")
			}
			isclose = 0
			tp = 1
			result = append(result, rule)
			rule = &RuleInfo{
				Tag: make(map[string]interface{}, 0),
			}
		default:
			if tp != 0 {
				value = append(value, runes[index])
			}
		}
		index++
	}
	if tp == 1 {
		rule.Tp = string(value)
	} else if tp == 2 || isclose == 1 {
		panic("[]必须成对出现")
	}
	for _, item := range result {
		if gstrings.IsEmptyOrWhite(item.Tp) {
			panic("验证类型为空")
		}
		item.Tp = strings.TrimSpace(item.Tp)
	}
	if !gstrings.IsEmptyOrWhite(rule.Tp) {
		rule.Tp = strings.TrimSpace(rule.Tp)
		result = append(result, rule)
	}
	return result
}

//Func 验证函数
//为了减少数据复制，使用一份元数据，不要随意修改fieldinfo和info中的数据信息
type Func func(value reflect.Value, field reflect.Value, fieldinfo *FieldInfo, info *RuleInfo) bool

//FuncRuleParmValid rule parm验证函数
//进行相关的数据预处理
type FuncRuleParmValid func(fieldinfo *FieldInfo, info *RuleInfo) (bool, string)

type funcType struct {
	Func  Func
	Valid FuncRuleParmValid
}

var validFuncCache = map[string]funcType{}
var validFuncCacheLock sync.RWMutex

//RegistValidFunc 注册验证函数
//f 必须，fv允许为空
func RegistValidFunc(name string, f Func, fv FuncRuleParmValid) {
	if f == nil || gstrings.IsEmptyOrWhite(name) {
		return
	}
	name = strings.TrimSpace(name)
	validFuncCacheLock.Lock()
	defer validFuncCacheLock.Unlock()
	if _, has := validFuncCache[name]; has {
		return
	}
	validFuncCache[name] = funcType{
		Func:  f,
		Valid: fv,
	}
}
