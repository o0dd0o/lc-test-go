package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		l, r := []int{0, 0}, []int{0, 0}
		if node.Left != nil {
			l = dfs(node.Left)
		}
		if node.Right != nil {
			r = dfs(node.Right)
		}
		no := int(math.Max(float64(l[0]), float64(l[1]))) + int(math.Max(float64(r[0]), float64(r[1])))
		return []int{l[1] + r[1] + node.Val, no}
	}
	re := dfs(root)
	return int(math.Max(float64(re[0]), float64(re[1])))
}
func main() {
	a := &TreeNode{Val: 3}
	b := &TreeNode{Val: 1}
	c := &TreeNode{Val: 2, Left: a}
	d := &TreeNode{Val: 3, Right: b}
	e := &TreeNode{Val: 3, Left: c, Right: d}
	println(rob(e))
}
