package main

func isBadVersion(version int) bool {
	return version >= 4
}
func firstBadVersion(n int) int {
	return n
}
func main() {
	println(firstBadVersion(5))
}
