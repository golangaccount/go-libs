package strings

import "testing"

func TestRandom(t *testing.T) {
	a := 'A'
	char := make([]rune, 26)
	char[0] = a
	for i := 0; i < 26; i++ {
		char[i] = rune(int(a) + i)
	}
	t.Log(string(char))
	t.Log(Random("1234567890", 100))
}
