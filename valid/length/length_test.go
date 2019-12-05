package length

import (
	"testing"

	"github.com/golangaccount/go-libs/valid"
)

type Length struct {
	Len  string `valid:"length[5]"`
	Pass string `valid:"length[8...20]"`
}

func TestLength(t *testing.T) {
	t.Log(valid.Valid(Length{Len: "123", Pass: "456"}))
	t.Log(valid.Valid(Length{Len: "12345", Pass: "1234567890"}))
	t.Log(valid.Valid(Length{Len: "12345", Pass: "123456789012345678901"}))
}
