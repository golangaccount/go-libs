package strings

import (
	"sort"
	"strings"
	"testing"
)

var ids []string

func init() {
	ids = make([]string, 3000)
	for i := range ids {
		ids[i] = Random("1234567890", 2)
	}
	ids = append(ids, "", "", "")
}
func TestUnique(t *testing.T) {
	v := Unique(ids)
	sort.Strings(v)
	t.Log(strings.Join(v, ","))
}

func BenchmarkUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unique(ids)
	}
}
