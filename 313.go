package main

func nthSuperUglyNumber(n int, primes []int) int {
	re := []int{}
	a := make([]int, len(primes))
	re = append(re, 1)
	for i := n; i > 1; i-- {
		min := 0
		for k, v := range primes {
			if min > re[a[k]]*v || min == 0 {
				min = re[a[k]] * v
			}
		}
		re = append(re, min)
		for k, v := range primes {
			if min == re[a[k]]*v {
				a[k]++
			}
		}
	}
	return re[n-1]
}
func main() {
	println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}
