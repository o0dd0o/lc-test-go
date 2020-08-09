package main

func findDuplicate(nums []int) int {
	s, l := 0, 0
	for s, l = nums[s], nums[nums[l]]; s != l; s, l = nums[s], nums[nums[l]] {

	}

	for s = 0; s != l; s, l = nums[s], nums[l] {

	}

	return s
}
func main() {
	println(findDuplicate([]int{1, 3, 4, 2, 2}))
	println(findDuplicate([]int{3, 1, 3, 4, 2}))
}
