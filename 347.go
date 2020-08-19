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
	xier(mapq, mapp)
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
		gaps = append(gaps, gaps[0]/3)
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

func guibin(arr map[int]int, mapq []int) {
	lens := len(mapq)
}
func main() {
	fmt.Printf("%v\n", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Printf("%v\n", topKFrequent([]int{1}, 1))
}
