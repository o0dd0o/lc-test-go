package main

func nthUglyNumber(n int) int {
	a2, a3, a5, re := 0, 0, 0, []int{}
	re = append(re, 1)
	for i := n; i > 1; i-- {
		if re[a2]*2 > re[a3]*3 {
			if re[a3]*3 > re[a5]*5 {
				re = append(re, re[a5]*5)
				a5++
			} else {
				re = append(re, re[a3]*3)
				if re[a3]*3 == re[a5]*5 {
					a5++
				}
				if re[a3]*3 == re[a2]*2 {
					a2++
				}
				a3++
			}
		} else {
			if re[a2]*2 > re[a5]*5 {
				re = append(re, re[a5]*5)
				a5++
			} else {
				re = append(re, re[a2]*2)
				if re[a2]*2 == re[a5]*5 {
					a5++
				}
				if re[a3]*3 == re[a2]*2 {
					a3++
				}
				a2++
			}
		}
	}
	return re[n-1]
}
func main() {
	println(nthUglyNumber(10))
}
