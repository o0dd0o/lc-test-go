package main

import (
	"math"
)

func numSquares(n int) int {
	sqrt := int(math.Sqrt(float64(n)))
	qmap := []int{}
	for i := 1; i <= sqrt; i++ {
		qmap = append(qmap, i*i)
	}
	dmap := []int{n}
	h := 0
	for len(dmap) > 0 {
		h += 1
		newdmap := []int{}
		for _, i := range dmap {
			for _, v := range qmap {
				if v > i {
					break
				} else if v == i {
					return h
				} else {
					newdmap = append(newdmap, i-v)
				}
			}
		}
		dmap = newdmap
	}
	return h
}

func main() {
	println(numSquares(12))
	println(numSquares(13))
}
