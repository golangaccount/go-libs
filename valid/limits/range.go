package limits

import (
	"encoding/base64"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	gstrings "github.com/golangaccount/go-libs/strings"
	"github.com/golangaccount/go-libs/valid"
)

//support type:
//int/uint/float/time.Time

func init() {
	valid.RegistValidFunc(RANGE, limit, validParm)
}

func validParm(fieldinfo *valid.FieldInfo, ruleinfo *valid.RuleInfo) (bool, string) {
	var start, end interface{}
	var err error
	values := strings.SplitN(ruleinfo.Parm, "...", -1)
	if len(values) != 2 {
		return false, "range的参数必须是xxx...xxx 或 ...xxx 或 xxx... 形式"
	}
	if fieldinfo.IsBase64 {
		for i, item := range values {
			bts, err := base64.StdEncoding.DecodeString(item)
			if err != nil {
				return false, "range参数值base64解码错误"
			}
			values[i] = string(bts)
		}
	}

	switch fieldinfo.Field.Type.Kind() {
	case reflect.Int:
		start, end, err = int64Range(values[0], values[1], 0)
	case reflect.Int8:
		start, end, err = int64Range(values[0], values[1], 8)
	case reflect.Int16:
		start, end, err = int64Range(values[0], values[1], 16)
	case reflect.Int32:
		start, end, err = int64Range(values[0], values[1], 32)
	case reflect.Int64:
		start, end, err = int64Range(values[0], values[1], 64)
	case reflect.Uint:
		start, end, err = uint64Range(values[0], values[1], 0)
	case reflect.Uint8:
		start, end, err = uint64Range(values[0], values[1], 8)
	case reflect.Uint16:
		start, end, err = uint64Range(values[0], values[1], 16)
	case reflect.Uint32:
		start, end, err = uint64Range(values[0], values[1], 32)
	case reflect.Uint64:
		start, end, err = uint64Range(values[0], values[1], 64)
	case reflect.Float32:
		start, end, err = floatRange(values[0], values[1], 32)
	case reflect.Float64:
		start, end, err = floatRange(values[0], values[1], 64)
	case reflect.Struct:
		if fieldinfo.Field.Type.PkgPath()+fieldinfo.Field.Type.Name() == "timeTime" {
			start, end, err = timeRange(values[0], values[1])
		} else {
			return false, "range只支持float和int或uint或time.Time形式数据"
		}
	default:
		return false, "range只支持float和int或uint形式数据"
	}
	if err != nil {
		return false, "range类型解析失败" + err.Error()
	}
	ruleinfo.Tag[rangeStart] = start
	ruleinfo.Tag[rangeEnd] = end
	return true, ""
}

func int64Range(start, end string, tp int) (interface{}, interface{}, error) {
	var s, e int64
	var err error
	if gstrings.IsEmptyOrWhite(start) {
		s = int64RangeSD(tp)
	} else {
		if s, err = strconv.ParseInt(start, 10, tp); err != nil {
			return s, e, err
		}
	}

	if gstrings.IsEmptyOrWhite(end) {
		e = int64RangeED(tp)
	} else {
		if e, err = strconv.ParseInt(end, 10, tp); err != nil {
			return s, e, err
		}
	}
	if s > e {
		s, e = e, s
	}
	return s, e, nil
}

func int64RangeSD(tp int) int64 {
	if tp == 0 {
		tp = strconv.IntSize
	}
	switch tp {
	case 8:
		return math.MinInt8
	case 16:
		return math.MinInt16
	case 32:
		return math.MinInt32
	case 64:
		return math.MinInt64
	}
	panic("数据类型错误")
}
func int64RangeED(tp int) int64 {
	if tp == 0 {
		tp = strconv.IntSize
	}
	switch tp {
	case 8:
		return math.MaxInt8
	case 16:
		return math.MaxInt16
	case 32:
		return math.MaxInt32
	case 64:
		return math.MaxInt64
	}
	panic("数据类型错误")
}

func uint64Range(start, end string, tp int) (interface{}, interface{}, error) {
	var s, e uint64
	var err error
	if gstrings.IsEmptyOrWhite(start) {
		s = uint64RangeSD(tp)
	} else {
		if s, err = strconv.ParseUint(start, 10, tp); err != nil {
			return s, e, err
		}
	}

	if gstrings.IsEmptyOrWhite(end) {
		e = uint64RangeED(tp)
	} else {
		if e, err = strconv.ParseUint(end, 10, tp); err != nil {
			return s, e, err
		}
	}
	if s > e {
		s, e = e, s
	}
	return s, e, nil
}
func uint64RangeSD(tp int) uint64 {
	return 0
}
func uint64RangeED(tp int) uint64 {
	if tp == 0 {
		tp = strconv.IntSize
	}
	switch tp {
	case 8:
		return math.MaxUint8
	case 16:
		return math.MaxUint16
	case 32:
		return math.MaxUint32
	case 64:
		return math.MaxUint64
	}
	panic("数据类型错误")
}

func floatRange(start, end string, tp int) (interface{}, interface{}, error) {
	var s, e float64
	var err error
	s, err = strconv.ParseFloat(start, tp)
	if err != nil {
		return s, e, err
	}
	e, err = strconv.ParseFloat(end, tp)
	if err != nil {
		return s, e, err
	}
	if s > e {
		s, e = e, s
	}
	return s, e, nil
}

func timeRange(start, end string) (interface{}, interface{}, error) {
	var s, e time.Time
	var err error
	if gstrings.IsEmptyOrWhite(start) {
		s, _ = time.Parse("2006-01-02", "1970-01-01")
	} else {
		s, err = time.Parse("2006-01-02 15:04:05", start)
	}
	if err != nil {
		return s, e, err
	}
	if gstrings.IsEmptyOrWhite(end) {
		e = time.Now().Add(time.Hour * 24 * 365 * 10)
	} else {
		e, err = time.Parse("2006-01-02 15:04:05", start)
	}
	if err != nil {
		return s, e, err
	}
	if s.After(e) {
		s, e = e, s
	}
	return s, e, nil
}

func limit(value reflect.Value, field reflect.Value, fieldinfo *valid.FieldInfo, ruleinfo *valid.RuleInfo) bool {
	switch fieldinfo.Field.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return limitint(field.Int(), ruleinfo)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return limituint(field.Uint(), ruleinfo)
	case reflect.Float32, reflect.Float64:
		return limitfloat(field.Float(), ruleinfo)
	case reflect.Struct:
		if fieldinfo.Field.Type.PkgPath()+fieldinfo.Field.Type.Name() == "timeTime" {
			return limittime(field.Interface().(time.Time), ruleinfo)
		} else {
			return false
		}
	default:
		panic("range只支持float和int或uint形式数据")
	}
}

func limitint(v int64, info *valid.RuleInfo) bool {
	var start, end int64
	start = info.Tag[rangeStart].(int64)
	end = info.Tag[rangeEnd].(int64)
	if v >= start && v <= end {
		return true
	}
	return false
}

func limituint(v uint64, info *valid.RuleInfo) bool {
	var start, end uint64
	start = info.Tag[rangeStart].(uint64)
	end = info.Tag[rangeEnd].(uint64)
	if v >= start && v <= end {
		return true
	}
	return false
}

func limitfloat(v float64, info *valid.RuleInfo) bool {
	var start, end float64
	start = info.Tag[rangeStart].(float64)
	end = info.Tag[rangeEnd].(float64)
	if v >= start && v <= end {
		return true
	}
	return false
}

func limittime(v time.Time, info *valid.RuleInfo) bool {
	var start, end time.Time
	start = info.Tag[rangeStart].(time.Time)
	end = info.Tag[rangeEnd].(time.Time)
	if v.Before(end) && v.After(start) {
		return true
	}
	return false
}
