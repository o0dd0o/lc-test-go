package main

import (
	"fmt"
)

func count(s string, t string) bool {
	sM, tM := s2i(s), s2i(t)
	lenS := len(sM)
	for i := 0; i < lenS; i++ {
		if sM[i] != tM[i] {
			return false
		}
	}
	return true
}

func s2i(s string) []int {
	lenS := len(s)
	a := make([]int, lenS)

	setS := make(map[uint8]int, lenS)

	for i := 0; i < lenS; i++ {
		if setS[s[i]] == 0 {
			setS[s[i]] = i
		}
	}

	for i := 0; i < lenS; i++ {
		a[i] = setS[s[i]]
	}

	return a
}

func main() {
	s := "add"
	t := "ege"

	d := count(s, t)
	fmt.Println(d)
}
