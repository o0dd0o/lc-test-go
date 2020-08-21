package main

/**
有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？

如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。

你允许：

装满任意一个水壶
清空任意一个水壶
从一个水壶向另外一个水壶倒水，直到装满或者倒空
*/
func canMeasureWater(x int, y int, z int) bool {
	return false
}
func main() {
	println(canMeasureWater(3, 5, 4)) //true
	println(canMeasureWater(2, 6, 5)) //false
}
