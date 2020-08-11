package main

func maxProduct(words []string) int {
	var difWord func(w1 string, w2 string) bool
	difWord = func(w1 string, w2 string) bool {
		qmap := map[int32]bool{}
		for _, v := range w1 {
			qmap[v] = true
		}
		for _, v := range w2 {
			if qmap[v] == true {
				return false
			}
		}
		return true
	}

	lens, max := len(words), 0
	for i := 0; i < lens; i++ {
		for j := i + 1; j < lens; j++ {
			if difWord(words[i], words[j]) && len(words[i])*len(words[j]) > max {
				max = len(words[i]) * len(words[j])
			}
		}
	}
	return max
}
func main() {
	println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
	println(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
	println(maxProduct([]string{"a", "aa", "aaa", "aaaa"}))
}
