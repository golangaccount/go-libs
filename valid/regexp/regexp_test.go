package regexp

import (
	"testing"

	"github.com/golangaccount/go-libs/valid"
)

type Regexp struct {
	Email string `valid:"email"`
	Url   string `valid:"url"`
	Phome string `valid:"mobilphone"`
	Reg   string `valid:"regexp[^\\S+$]"`
}

func TestRegexp(t *testing.T) {
	t.Log(valid.Valid(Regexp{Email: ""}))
	t.Log(valid.Valid(Regexp{Email: "123@11.com", Url: "http://www.baidu.com", Phome: "18212345678", Reg: "1234456"}))
}
