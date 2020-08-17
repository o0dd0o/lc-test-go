package main

import "fmt"

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	if num == 1 {
		return []int{0, 1}
	}
	re := []int{0, 1}
	for i := 2; i <= num; i++ {
		re = append(re, re[i>>1]+i&1)
	}
	return re
}
func main() {
	fmt.Printf("%v\n", countBits(2)) //[0,1,1]
	fmt.Printf("%v\n", countBits(5)) //[0,1,1,2,1,2]
}
