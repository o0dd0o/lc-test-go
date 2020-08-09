package main

func lengthOfLIS(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}
	dp := []int{nums[0]}
	maxlen := 1
	for k := 1; k < lens; k++ {
		if nums[k] > dp[maxlen-1] {
			dp = append(dp, nums[k])
			maxlen++
		} else {
			s, l := 0, maxlen-1
			for s <= l {
				mid := s + (l-s)>>2
				if dp[mid] > nums[k] {
					l = mid - 1
				} else if dp[mid] < nums[k] {
					s = mid + 1
				} else {
					l -= 1
				}
			}
			dp[s] = nums[k]
		}
	}

	return maxlen
}

func main() {
	println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	println(lengthOfLIS([]int{2, 2}))
}
