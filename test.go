package main

import (
	"fmt"
)

func count(s string, t string) int {
	return 3
}

func s2i(s string) []int {
	var a []int
	fmt.Printf("%c\n", s[0])
	return a
}

func main() {
	s := "add"
	t := "egg"
	s2i(s)
	d := count(s, t)
	fmt.Println(d)
}
