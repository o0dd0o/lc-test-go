package main

import "math"

func maxProfit(prices []int) int {
	f1, f2, f3 := 0, 0, 0
	for k, v := range prices {
		if k == 0 {
			f1 = -v
		} else {
			f1, f2, f3 = int(math.Max(float64(f1), float64(f2-v))),
				int(math.Max(float64(f2), float64(f3))),
				f1+v
		}
	}
	if f2 > f3 {
		return f2
	}
	return f3
}
func main() {
	println(maxProfit([]int{1, 2, 3, 0, 2})) //3
}
