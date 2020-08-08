package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	lens, smap := len(s), map[uint8]int{}
	for i := 0; i < lens; i++ {
		smap[s[i]] += 1
	}
	for i := 0; i < lens; i++ {
		smap[t[i]] -= 1
		if smap[t[i]] < 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	re := isAnagram(s, t)
	println(re)
}
