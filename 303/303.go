package main

type NumArray struct {
	arr []int
	sum []int
	cal bool
}

func Constructor(nums []int) NumArray {
	return NumArray{arr: nums, sum: make([]int, len(nums))}
}

func (this *NumArray) SumRange(i int, j int) int {
	if this.cal {
		if i == 0 {
			return this.sum[j]
		} else {
			return this.sum[j] - this.sum[i-1]
		}
	}
	for k, v := range this.arr {
		if k > 0 {
			this.sum[k] = this.sum[k-1] + v
		} else {
			this.sum[k] = v
		}
	}
	this.cal = true
	return this.SumRange(i, j)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	println(obj.SumRange(0, 2))
	println(obj.SumRange(2, 5))
	println(obj.SumRange(0, 5))
}
