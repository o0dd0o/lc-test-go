package main

/**
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。



示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for k, _ := range dp {
		if k == 0 {
			dp[0] = -1
			continue
		}
		min := 0
		for _, c := range coins {
			if k < c {
				continue
			} else if k == c {
				min = 1
				break
			} else {
				if dp[k-c] < 0 {
					continue
				}
				if dp[k-c] > 0 && (dp[k-c]+1 < min || min == 0) {
					min = dp[k-c] + 1
				}
			}
		}

		if min == 0 {
			dp[k] = -1
		} else {
			dp[k] = min
		}

	}

	return dp[amount]
}
func main() {
	println(coinChange([]int{1, 2, 5}, 11)) //3
	println(coinChange([]int{2}, 3))        //-1
}
