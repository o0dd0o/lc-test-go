package main

func hIndex(citations []int) int {
	lens := len(citations)
	cou := make([]int, lens+1)

	for _, i := range citations {
		if i > lens {
			i = lens
		}
		cou[i] += 1
	}
	all := 0
	for i := lens; i >= 0; i-- {
		all += cou[i]
		if all >= i {
			return i
		}
	}
	return 0
}
func main() {
	println(hIndex([]int{3, 0, 6, 1, 5}))
}
