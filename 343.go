package main

import "math"

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n%3 == 0 {
		return int(math.Pow(3, float64(n/3)))
	}
	if n%3 == 1 {
		return int(math.Pow(3, float64((n-4)/3))) * 4
	}
	if n%3 == 2 {
		return int(math.Pow(3, float64((n-2)/3))) * 2
	}
	return 0
}
func main() {
	println(integerBreak(2))
	println(integerBreak(10))
}
