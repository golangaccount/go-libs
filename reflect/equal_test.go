package reflect

import (
	"fmt"
	"testing"
)

type EqualStruct struct {
	Id   string
	Name string
}

func (q *EqualStruct) Equal(e *EqualStruct) bool {
	if q.Id == e.Id {
		return true
	} else {
		return false
	}
}

type EqualStructNoFunc struct {
	Id   string
	Name int
}

type EqualStructArray []*EqualStruct

func (esa EqualStructArray) Equal(es EqualStructArray) bool {
	fmt.Println("---------------")
	return true
}

func TestEqual(t *testing.T) {
	// boolparm1 := true
	// boolparm2 := true
	// boolparm3 := false
	// t.Log(Equal(boolparm1, boolparm2), Equal(boolparm2, boolparm3))

	// int1 := 1
	// int2 := 1
	// int3 := 2
	// t.Log(Equal(int1, int2), Equal(int2, int3))

	// q1 := EqualStruct{"1", "zhang"}
	// q2 := EqualStruct{"1", "wang"}
	// q3 := &EqualStruct{"1", "chen"}
	// q44 := &EqualStruct{"1", "chen"}
	// t.Log(Equal(q1, q2), Equal(&q1, &q2), Equal(q2, &q3), Equal(q44, &q3))

	// q4 := EqualStructNoFunc{"1", 1}
	// q5 := EqualStructNoFunc{"1", 1}
	// q6 := EqualStructNoFunc{"1", 2}
	// t.Log(Equal(q4, q5), Equal(q5, q6))

	// a := [4]string{"a", "b", "c", "d"}
	// b := [4]string{"a", "b", "c", "d"}
	// c := [4]string{"a", "b", "c", "e"}
	// d := [3]string{"a", "b", "c"}
	// t.Log(Equal(a, b), Equal(b, c), Equal(b, d))
	eq1 := make(EqualStructArray, 1)
	eq1[0] = &EqualStruct{"1", "zhang"}
	eq2 := make(EqualStructArray, 1)
	eq2[0] = &EqualStruct{"2", "li"}
	t.Log(Equal(eq1, eq2))

}
