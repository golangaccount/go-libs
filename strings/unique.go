package strings

//Unique 唯一化字符串数组
func Unique(src []string) []string {
	dic := map[string]bool{}
	for _, item := range src {
		dic[item] = true
	}
	result := make([]string, 0, len(dic))
	for k := range dic {
		result = append(result, k)
	}
	return result
}
