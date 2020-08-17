package main

import "fmt"

func reverseVowels(s string) string {
	a, b := 0, len(s)-1
	v := []byte(s)
	mapq := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	for a < b {
		if mapq[v[a]] && mapq[v[b]] {
			v[a], v[b] = v[b], v[a]
			a++
			b--
		} else if mapq[v[a]] {
			b--
		} else if mapq[v[b]] {
			a++
		} else {
			a++
			b--
		}
	}
	return string(v)
}
func main() {
	fmt.Printf("%v\n", reverseVowels("hello"))
	fmt.Printf("%v\n", reverseVowels("leetcode"))
}
