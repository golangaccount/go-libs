package collection

import (
	"encoding/json"
	"testing"
)

type TS struct {
	Id   string
	Name string
}

func (ts *TS) Equal(t *TS) bool {
	if ts.Id == t.Id {
		return true
	} else {
		return false
	}
}

func TestUnique(t *testing.T) {
	result := make([]string, 0)
	t.Log(Unique(&result, []string{"1", "2", "1"}, [4]string{"0", "1", "2", "3"}))
	t.Log(result)
	resultint := make([]int, 0)
	t.Log(Unique(&resultint, []int{1, 2, 2, 3}, [4]int{0, 1, 2, 3}))
	t.Log(resultint)
	results := make([]*TS, 0)
	t.Log(Unique(&results, []*TS{
		&TS{"1", "a"}, &TS{"2", "b"}, &TS{"1", "b"},
	}, [2]*TS{
		&TS{"1", "a"}, &TS{"3", "b"},
	}))
	bts, _ := json.Marshal(results)
	t.Log(string(bts))
}
