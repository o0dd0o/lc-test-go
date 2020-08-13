package main

import "fmt"

/**
给定一个无序的数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。

示例 1:

输入: nums = [1, 5, 1, 1, 6, 4]
输出: 一个可能的答案是 [1, 4, 1, 5, 1, 6]
示例 2:

输入: nums = [1, 3, 2, 2, 3, 1]
输出: 一个可能的答案是 [2, 3, 1, 3, 1, 2]
说明:
你可以假设所有输入都会得到有效的结果。

进阶:
你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/wiggle-sort-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func wiggleSort(nums []int) {
	lens := len(nums)
	mid := float32(lens-1) / 2
	var fastSelect func(nums []int, int2 int, int3 int)
	fastSelect = func(nums []int, start int, end int) {
		t := nums[end]
		i, j := start, start
		for ; j <= end; j++ {
			if nums[j] <= t {
				if j != i || nums[i] != nums[j] {
					nums[j], nums[i] = nums[i], nums[j]
				}
				i++
			}
		}
		if start < end-1 {
			if float32(i-1) > mid {
				fastSelect(nums, start, i-2)
			} else if float32(i-1) < mid {
				fastSelect(nums, i, end)
			}
		}

	}
	fastSelect(nums, 0, lens-1)

	midv := 0.0
	if lens%2 == 0 {
		midv = float64(nums[int(mid+0.5)]+nums[int(mid-0.5)]) / 2
	} else {
		midv = float64(nums[int(mid)])
	}
	var mapf func(int2 int) int
	mapf = func(int2 int) int {
		if float32(int2) <= mid {
			return (int(mid) - int2) * 2
		}
		return (lens-int2)*2 - 1
	}
	i, j, k := 0, 0, lens-1
	for j <= k {
		now := float64(nums[mapf(j)])
		if now < midv {
			nums[mapf(j)], nums[mapf(i)] = nums[mapf(i)], nums[mapf(j)]
			i++
			j++
		} else if now > midv {
			nums[mapf(j)], nums[mapf(k)] = nums[mapf(k)], nums[mapf(j)]
			k--
		} else {
			j++
		}
	}
}
func main() {
	ints1 := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(ints1)
	fmt.Printf("%v\n", ints1)

	ints2 := []int{1, 3, 2, 2, 3, 1}
	wiggleSort(ints2)
	fmt.Printf("%v\n", ints2)

	ints3 := []int{1, 1, 2, 2, 2, 1}
	wiggleSort(ints3)
	fmt.Printf("%v\n", ints3)
	ints4 := []int{4, 5, 5, 6}
	wiggleSort(ints4)
	fmt.Printf("%v\n", ints4)
}
