package main

import "fmt"

func reverseString(s []byte) {
	a, b := 0, len(s)-1
	for a < b {
		s[a], s[b] = s[b], s[a]
		a++
		b--
	}
}
func main() {
	fmt.Printf("%v\n", []byte{'h', 'e', 'l', 'l', 'o'})
	fmt.Printf("%v\n", []byte{'H', 'a', 'n', 'n', 'a', 'h'})
}
