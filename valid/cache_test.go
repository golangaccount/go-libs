package valid

import (
	"reflect"
	"testing"
)

func TestRule(t *testing.T) {
	t.Log(parse("abc [12:] ]; mkq"))
}

type ts struct {
	Id string `valid:"abc [12:] ]; mkq"`
}

func TestParse(t *testing.T) {
	t.Log(Parse(reflect.TypeOf(ts{})))
}
