package main

func countNumbersWithUniqueDigits(n int) int {
	if n < 1 {
		return 1
	}
	if n > 10 {
		n = 10
	}
	mapq := []int{1, 9}
	for i := 2; i <= 10; i++ {
		mapq = append(mapq, mapq[i-1]*(11-i))
	}
	re := 0
	for i := 0; i <= n; i++ {
		re += mapq[i]
	}
	return re
}
func main() {
	println(countNumbersWithUniqueDigits(2)) //91
	println(countNumbersWithUniqueDigits(3))
	println(countNumbersWithUniqueDigits(4))
}
