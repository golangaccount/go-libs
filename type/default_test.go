package gtype

import (
	"testing"
)

func TestDefault(t *testing.T) {
	t.Log(IsDefault(int(0)))
	t.Log(IsDefault(int8(1)))
	t.Log(IsDefault(""))
	t.Log(IsDefault(0 + 0i))
	t.Log(IsDefault(-1i * -1i))
}
