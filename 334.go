package main

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	min := INT_MIN
	allmin := nums[0]
	for k, v := range nums {
		if v < allmin {
			allmin = v
		}
		if k > 0 && v > allmin {
			if min == INT_MIN {
				min = v
			} else if v <= min {
				min = v
			} else {
				return true
			}
		}

	}
	return false
}
func main() {
	println(increasingTriplet([]int{1, 2, 3, 4, 5}))                //true
	println(increasingTriplet([]int{5, 4, 3, 2, 1}))                //false
	println(increasingTriplet([]int{1, 0, 0, -1, 0, 10}))           //false
	println(increasingTriplet([]int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2})) //false
	println(increasingTriplet([]int{5, 1, 5, 5, 2, 5, 4}))          //false
}
