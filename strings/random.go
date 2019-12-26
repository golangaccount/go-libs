package strings

import (
	"math/rand"
	"time"
)

//Random 通过指定的字符串生成随机的字符串
//src 指定字符串
//length 生成的字符长度
func Random(src string, length int) string {
	runes := []rune(src)
	runeslen := len(runes)
	if runeslen < 1 {
		panic("")
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]rune, length)
	for i := 0; i < length; i++ {
		result[i] = runes[random.Intn(len(src))]
	}
	return string(result)
}
