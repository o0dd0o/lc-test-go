package main

func isBadVersion(version int) bool {
	return version >= 3
}
func firstBadVersion(n int) int {
	s, l := 1, n

	for s <= l {
		mid := s + int((l-s)/2)
		if isBadVersion(mid) {
			l = mid - 1
		} else {
			s = mid + 1
		}
	}

	return s
}
func main() {
	println(firstBadVersion(5))
}
