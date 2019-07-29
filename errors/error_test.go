package errors

import (
	"testing"
)

func TestNewError(t *testing.T) {
	t.Log(NewError("[tp]", "[info]"))
	t.Log(NewError("[tp]", "[info]").String())
}
