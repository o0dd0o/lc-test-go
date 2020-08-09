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
func (this *NumArray) Update(i int, val int) {
	update := val - this.arr[i]
	this.arr[i] = val
	if this.cal {
		lens := len(this.sum)
		for ; i < lens; i++ {
			this.sum[i] += update
		}
	}

}
func main() {
	nums := []int{1, 3, 5}
	obj := Constructor(nums)
	println(obj.SumRange(0, 2))
	obj.Update(1, 2)
	println(obj.SumRange(0, 2))
}
