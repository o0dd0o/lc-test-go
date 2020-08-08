package main

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	divide := [3]int{2, 3, 5}
	for _, i := range divide {
		for num%i == 0 {
			num = num / i
		}
	}
	return num == 1
}

func main() {
	println(isUgly(14))
	println(isUgly(6))
	println(isUgly(8))
}
