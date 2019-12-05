package enum

import (
	"reflect"
	"testing"

	"github.com/golangaccount/go-libs/valid"
)

type Enum struct {
	String string `valid:"enum[123,456]"`
	Int    int    `valid:"enum[123,456]"`
	Int8   int8   `valid:"enum[-12,45]"`
	Uint   uint   `valid:"enum[123,456]"`
	Uint8  uint8  `valid:"enum[123,12]"`
}

func TestEnum(t *testing.T) {
	t.Log(valid.Valid(Enum{
		String: "123",
		Int:    123,
		Int8:   -12,
		Uint:   123,
		Uint8:  123,
	}))
}

func TestDynamic(t *testing.T) {
	tp := valid.Parse(reflect.TypeOf(Enum{}))
	valid.RegistTypeInfo(reflect.TypeOf(Enum{}), tp)
	t.Log(valid.Valid(Enum{
		String: "123",
		Int:    123,
		Int8:   -12,
		Uint:   123,
		Uint8:  123,
	}))
	for _, item := range tp.Fields {
		if item.Name == "String" {
			for _, rule := range item.Rules {
				if rule.Tp == ENUM {
					rule.Tag[ENUM] = []string{"789", "0"}
				}
			}
		}
	}
	t.Log(valid.Valid(Enum{
		String: "123",
		Int:    123,
		Int8:   -12,
		Uint:   123,
		Uint8:  123,
	}))
	t.Log(valid.Valid(Enum{
		String: "789",
		Int:    123,
		Int8:   -12,
		Uint:   123,
		Uint8:  123,
	}))
}
