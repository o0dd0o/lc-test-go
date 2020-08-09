package main

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	smap := strings.Fields(str)
	if len(smap) != len(pattern) {
		return false
	}
	dmap := map[string]int{}
	dmap2 := map[int32]int{}
	gmap := []int{}
	gmap2 := []int{}
	for i, v := range smap {
		if dmap[v] == 0 {
			dmap[v] = i + 1
		}
		gmap = append(gmap, dmap[v]-1)
	}
	for i, v := range pattern {
		if dmap2[v] == 0 {
			dmap2[v] = i + 1
		}
		gmap2 = append(gmap2, dmap2[v]-1)
	}

	for i, v := range gmap {
		if gmap2[i] != v {
			return false
		}
	}
	return true
}
func main() {
	pattern := "abba"
	str := "dog cat cat dog"
	println(wordPattern(pattern, str))
	println(wordPattern("abba", "dog dog dog dog"))
}
