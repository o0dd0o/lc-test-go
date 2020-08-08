package main

func hIndex(citations []int) int {
	lens := len(citations)
	s, l := 0, lens-1
	for s <= l {
		mid := s + int((l-s)/2)
		if lens-mid < citations[mid] {
			l = mid - 1
		} else if lens-mid > citations[mid] {
			s = mid + 1
		} else {
			return lens - mid
		}

	}
	return lens - s
}
func main() {
	println(hIndex([]int{0, 1, 3, 5, 6}))
	println(hIndex([]int{0}))
	println(hIndex([]int{1}))
	println(hIndex([]int{0, 2}))
}
