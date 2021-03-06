package main

/**
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

示例 1:

输入: 1
输出: true
解释: 2^0 = 1
示例 2:

输入: 16
输出: true
解释: 2^4 = 16
示例 3:

输入: 218
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/power-of-two
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isPowerOfTwo(n int) bool {
	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n = n >> 1
	}
	return n == 1
}

func main() {
	println(isPowerOfTwo(1))
	println(isPowerOfTwo(16))
	println(isPowerOfTwo(218))
}
