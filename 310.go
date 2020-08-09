package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	if n == 2 {
		return []int{0, 1}
	}
	qmap := map[int][]int{}
	for k, v := range edges {
		qmap[k] = v
	}
	var re []int
	c := map[int]int{}
	for _, v := range qmap {
		c[v[0]] += 1
		c[v[1]] += 1
	}
	for len(c) > 2 {
		re = []int{}
		for k, v := range c {
			if v == 1 {
				re = append(re, k)
				delete(c, k)
			}
		}
		if len(re) == 0 {
			break
		}
		for k, v := range qmap {
			for _, k1 := range re {
				if v[0] == k1 {
					delete(qmap, k)
					c[v[1]] -= 1
					break
				}
				if v[1] == k1 {
					delete(qmap, k)
					c[v[0]] -= 1
					break
				}
			}
		}
	}
	ret := []int{}
	for k, _ := range c {
		ret = append(ret, k)
	}
	return ret
}
func main() {
	edges := [][]int{
		{1, 0}, {1, 2}, {1, 3},
	}
	fmt.Printf("%v\n", findMinHeightTrees(4, edges))
	edges = [][]int{
		{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4},
	}
	fmt.Printf("%v\n", findMinHeightTrees(6, edges))
}
