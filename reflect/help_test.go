package reflect

import (
	"testing"
)

func TestIndirect(t *testing.T) {
	A := "123中国"
	t.Log(A[:1])
	t.Log(A[4:])
	// ts := &TS{}
	// v := reflect.ValueOf(&ts)
	// t.Log("original:", v.Kind().String())
	// v = Indirect(v)
	// t.Log("Indirect:", v.Kind().String())

}
