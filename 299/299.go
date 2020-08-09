package main

import (
	"strconv"
)

func getHint(secret string, guess string) string {
	amap := map[int32]int{}
	for _, v := range secret {
		amap[v] += 1
	}
	a, b := 0, 0
	for k, v := range guess {
		if secret[k] == uint8(v) {
			a += 1
			amap[v] -= 1
		}
	}
	for k, v := range guess {
		if secret[k] != uint8(v) && amap[v] > 0 {
			b += 1
			amap[v] -= 1
		}
	}

	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}
func main() {
	println(getHint("1122", "1222"))
	println(getHint("1123", "0111"))
}
