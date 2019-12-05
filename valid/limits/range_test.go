package limits

import (
	"testing"

	"github.com/golangaccount/go-libs/valid"
)

type Range struct {
	Int   int     `valid:"range[...1000]"`
	Uint  uint    `valid:"range[5...10]"`
	Float float32 `valid:"range[-1.2...5.6]"`
}

func TestRange(t *testing.T) {
	t.Log(valid.Valid(Range{Int: 500, Uint: 11, Float: -10.5}))
	t.Log(valid.Valid(Range{Int: 500, Uint: 10, Float: -0.5}))
}
