package main

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	tmap := []string{}
	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, path string) {
		if node.Left != nil {
			dfs(node.Left, path+"->"+strconv.Itoa(node.Left.Val))
		}

		if node.Right != nil {
			dfs(node.Right, path+"->"+strconv.Itoa(node.Right.Val))
		}

		if node.Left == nil && node.Right == nil {
			tmap = append(tmap, path)
		}
	}
	dfs(root, strconv.Itoa(root.Val))

	return tmap
}

func main() {
	a := &TreeNode{Val: 5}
	b := &TreeNode{Val: 2, Right: a}
	c := &TreeNode{Val: 3}
	d := &TreeNode{Val: 1, Left: b, Right: c}
	n := binaryTreePaths(d)
	fmt.Printf("%v", n)

}
