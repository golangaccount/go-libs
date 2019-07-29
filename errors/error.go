package errors

import (
	"errors"
	"fmt"
	"strings"
)

func New(err string) error {
	return errors.New(err)
}

func Error(err error, str string) error {
	if err != nil {
		return err
	} else {
		return New(str)
	}
}

//携带错误类型信息
type typeError struct {
	tp   string
	info string
}

func (te *typeError) Error() string {
	return te.info
}

func (te *typeError) Type() string {
	return te.tp
}

func (te *typeError) String(format ...string) string {
	if len(format) == 0 {
		//便于输出错误的解析操作，将[或]使用\[或\]进行转义
		replace := strings.NewReplacer("[", "\\[", "]", "\\]")
		return fmt.Sprintf("[type:%s][error:%s]", replace.Replace(te.tp), replace.Replace(te.info))
	} else {
		return fmt.Sprintf(format[0], te.tp, te.info)
	}
}

func NewError(tp, info string) *typeError {
	return &typeError{
		tp:   tp,
		info: info,
	}
}
