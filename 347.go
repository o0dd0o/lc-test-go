package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	mapq := map[int]int{}
	mapp := []int{}
	for _, v := range nums {
		mapq[v] += 1
	}
	for k, _ := range mapq {
		mapp = append(mapp, k)
	}
	mapp = jishu(mapq, mapp)
	var re []int
	lens := len(mapq)
	for i := 0; i < k; i++ {
		re = append(re, mapp[lens-i-1])
	}
	return re
}

func maopao(arr map[int]int, mapq []int) {
	lens := len(mapq)
	for i := 0; i < lens-1; i++ {
		for j := 0; j < lens-1; j++ {
			if arr[mapq[j]] > arr[mapq[j+1]] {
				mapq[j], mapq[j+1] = mapq[j+1], mapq[j]
			}
		}
	}
}

func xuanze(arr map[int]int, mapq []int) {
	lens := len(mapq)
	for i := 0; i < lens-1; i++ {
		min := i
		for j := i + 1; j < lens; j++ {
			if arr[mapq[min]] > arr[mapq[j]] {
				min = j
			}
		}
		mapq[min], mapq[i] = mapq[i], mapq[min]
	}
}
func charu(arr map[int]int, mapq []int) {
	lens := len(mapq)
	for i := 1; i < lens; i++ {
		current := mapq[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[mapq[j]] > arr[current] {
				mapq[j+1] = mapq[j]
			} else {
				break
			}
		}
		mapq[j+1] = current
	}
}

func xier(arr map[int]int, mapq []int) {
	lens := len(mapq)
	gaps := []int{lens / 3}
	for gaps[len(gaps)-1] >= 3 {
		gaps = append(gaps, gaps[len(gaps)-1]/3)
	}
	if lens < 3 {
		gaps = []int{1}
	}
	for _, gap := range gaps {
		for i := gap; i < lens; i++ {
			current := mapq[i]
			j := i - gap
			for ; j >= 0; j -= gap {
				if arr[mapq[j]] > arr[current] {
					mapq[j+gap] = mapq[j]
				} else {
					break
				}
			}
			mapq[j+gap] = current
		}
	}
}

func guibin(arr map[int]int, mapq []int) []int {
	lens := len(mapq)
	if lens < 2 {
		return mapq
	}
	return merge(arr, guibin(arr, mapq[:lens/2]), guibin(arr, mapq[lens/2:]))
}
func merge(arr map[int]int, left []int, right []int) []int {
	re := []int{}
	leni := len(left)
	lenj := len(right)
	for i, j := 0, 0; i < leni || j < lenj; {
		if i == leni {
			re = append(re, right[j])
			j++
		} else if j == lenj {
			re = append(re, left[i])
			i++
		} else if arr[left[i]] > arr[right[j]] {
			re = append(re, right[j])
			j++
		} else {
			re = append(re, left[i])
			i++
		}
	}
	return re
}
func kuaisu(arr map[int]int, mapq []int, s, l int) {
	if s < l {
		current := mapq[s]
		s1, l1 := s, l
		for s < l {
			for ; s < l && arr[mapq[l]] > arr[current]; l-- {

			}
			if s < l {
				mapq[s] = mapq[l]
			} else {
				break
			}
			for ; s < l && arr[mapq[s]] <= arr[current]; s++ {

			}
			if s < l {
				mapq[l] = mapq[s]
			} else {
				break
			}
		}
		mapq[s] = current
		kuaisu(arr, mapq, s1, s1-1)
		kuaisu(arr, mapq, s1+1, l1)
	}
}
func dui(arr map[int]int, mapq []int) {
	lens := len(mapq)
	for i := lens / 2; i >= 0; i-- {
		shiftDown(arr, mapq, i, lens)
	}

	for i := lens - 1; i >= 0; i-- {
		mapq[i], mapq[0] = mapq[0], mapq[i]
		shiftDown(arr, mapq, 0, i)
	}

}
func shiftDown(arr map[int]int, mapq []int, index int, lens int) {
	tar := 2 * index
	if tar+1 < lens && arr[mapq[tar+1]] > arr[mapq[tar]] {
		tar++
	}
	if tar < lens && arr[mapq[tar]] > arr[mapq[index]] {
		mapq[tar], mapq[index] = mapq[index], mapq[tar]
		shiftDown(arr, mapq, tar, lens)
	}
}
func jishu(arr map[int]int, mapq []int) []int {
	re := []int{}
	lens := len(mapq)
	max, min := arr[mapq[0]], arr[mapq[0]]
	for i := 1; i < lens; i++ {
		if arr[mapq[i]] > max {
			max = arr[mapq[i]]
		}
		if min > arr[mapq[i]] {
			min = arr[mapq[i]]
		}
	}
	bu := make(map[int][]int, max-min+1)
	for i := 0; i < lens; i++ {
		bu[arr[mapq[i]]] = append(bu[arr[mapq[i]]], mapq[i])
	}

	for i := min; i <= max; i++ {
		for _, v := range bu[i] {
			re = append(re, v)
		}
	}

	return re
}

func main() {
	fmt.Printf("%v\n", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Printf("%v\n", topKFrequent([]int{1}, 1))
	fmt.Printf("%v\n", topKFrequent([]int{-5, -9, 67, -10, 4, -57, 47, 13, -67, -26, -57, 63, 38, -68, 62, -45, -37, 95, 49, -91, -53, -42, -33, 80, 78, -30, -36, 22, 9, -86, 79, -1, 44, -92, 30, -68, -94, 58, -51, -26, -38, 5, -74, 25, 71, -93, 52, -12, -86, 7, -86, 53, 78, -31, -5, -87, 88, -98, -39, 9, 99, 23, 96, -90, 51, -64, 35, -73, 9, 60, -78, 70, 99, 14, 70, 71, -78, 50, 7, 46, -89, 57, -1, 87, -87, -95, 48, 49, 79, -54, 31, 28, -27, 75, 31, -76, -12, 35, 40, -90, -60, -60, -7, 67, 73, -34, -42, -35, 61, 51, 18, 95, 16, 78, -81, -91, -30, 92, 57, -79, 5, 41, 29, 72, -62, -47, 80, 29, 1, -21, -36, 5, 82, 4, -12, -52, -56, -47, -68, 95, 85, -87, -7, -12, 98, 75, -64, -93, 11, 73, -81, -9, -12, -9, 51, 17, -94, 33, -9, 57, -35, 10, -17, 87, -18, -55, -52, 30, -62, 73, 35, -74, -47, -63, 77, -72, -55, 5, 73, 21, 14, 7, -65, -51, -55, -49, 98, -20, -22, -68, 34, -20, 92, 55, 47, -20, 6, -54, -12, 3, 75, 69, 60, 15, 88, 64, 2, -27, -50, 55, 73, 46, -15, -64, 93, -47, -75, -55, -75, 21, -57, 91, -12, -99, -68, -56, -14, -4, -77, -94, 55, 93, -31, 68, -12, -23, 59, -56, -86, 43, 83, -93, -78, -11, -7, 96, -3, -87, -37, 19, -78, 67, -29, 77, -28, 91, -73, -68, -22, 18, -7, 3, 15, 77, 99, 31, -48, -86, -45, -82, 52, -39, 8, -88, -83, -58, -77, 5, 87, -61, 50, 32, -66, -36, 60, -53, 52, 70, -36, -1, 83, -56, 33, 98, -80, 28, 1, -21, -50, -60, 44, 99, 18, 83, -29, 83, -36, -55, -6, 96, -60, 61, 75, 6, -57, 2, 82, 62, -27, 36, 60, 72, 92, 61, -65, 79, -57, -34, -23, -28, -55, 53, 36, -80, 5, -76, 64, -81, -32, -43, -1, 50, -16, -72, -74, 22, 88, 28, -79, -99, 85, -13, -34, -76, 85, 6, 21, -99, 10, -46, 79, 11, -70, 17, 47, -22, -62, 0, 6, 75, -19, 57, -25, -52, -83, 90, 21, 95, 52, 68, 47, -12, 76, -9, -65, 86, 90, 16, 74, 64, 26, 84, 64, -42, 97, -72, 53, -76, 11, 89, -62, 67, 100, 15, 53, 68, -16, 24, 11, -77, 20, 59, -95, -50, -20, 27, 45, 94, 13, -93, 86, 49, 12, 19, 17, -33, -52, -28, 71, 79, -19, -73, 40, -99, 83, 77, 19, -20, 98, 86, -5, -5, 73, 18, 100, 73, -45, 33, 3, 89, 32, -53, 73, 16, -3, -26, -80, 49, -78, 67, 31, 1, -85, -44, -91, -68, 75, -74, 95, 23, 89, 99, -84, 54, -93, 68, 0, -41, 66, -15, -27, -23, -9, -68, 37, 45, -69, 57, 80, 10, -64, 35, 26, 55, -67, 31, -76, 36, -99, 21}, 7))
	fmt.Printf("%v\n", topKFrequent([]int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3}, 3))
	//-12 -68 73 -9 -55 5 75
}
