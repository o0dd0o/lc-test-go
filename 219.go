package main

import (
	"fmt"
)

/**
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。



示例 1:

输入: nums = [1,2,3,1], k = 3
输出: true
示例 2:

输入: nums = [1,0,1,1], k = 1
输出: true
示例 3:

输入: nums = [1,2,3,1,2,3], k = 2
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	numsLast := make(map[int]int)
	numsLen := len(nums)
	min := 0
	for i := 0; i < numsLen; i++ {
		v := nums[i]
		if numsLast[v] > 0 && (i-numsLast[v]+1 < min || min == 0) {
			min = i - numsLast[v] + 1
		}
		numsLast[v] = i + 1
	}

	return min <= k && min > 0
}
func main() {
	a := []int{1, 2, 3, 1}
	b := []int{1, 0, 1, 1}
	c := []int{1, 2, 3, 1, 2, 3}
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	f := []int{99, 99}
	e := []int{2, 2}

	fmt.Printf("%v\n", containsNearbyDuplicate(a, 3))
	fmt.Printf("%v\n", containsNearbyDuplicate(b, 1))
	fmt.Printf("%v\n", containsNearbyDuplicate(c, 2))
	fmt.Printf("%v\n", containsNearbyDuplicate(d, 15))
	fmt.Printf("%v\n", containsNearbyDuplicate(f, 2))
	fmt.Printf("%v\n", containsNearbyDuplicate(e, 3))
}
