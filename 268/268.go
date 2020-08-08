package main

func missingNumber(nums []int) int {
	sentry := 0
	lens := len(nums)
	nums = append(nums, 0)
	for i := 0; i < lens; i++ {
		if nums[i] != i {
			sentry = nums[i]
			for sentry != nums[sentry] {
				tm := nums[sentry]
				nums[sentry] = sentry
				sentry = tm
			}
		}
	}
	for k, i := range nums {
		if k != i {
			return k
		}
	}
	return 0
}
func main() {
	println(missingNumber([]int{0}))
	println(missingNumber([]int{3, 0, 1}))
	println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
