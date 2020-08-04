package main

import "container/list"

/**
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
备注:
这个问题是受到 Max Howell 的 原问题 启发的 ：

谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/invert-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	l := list.New()
	l.PushBack(root)
	for l.Front() != nil {
		t := l.Front()
		l.Remove(t)
		p := t.Value.(*TreeNode)
		if p.Left != nil {
			l.PushBack(p.Left)
		}
		if p.Right != nil {
			l.PushBack(p.Right)
		}
		p.Left, p.Right = p.Right, p.Left
	}
	return root
}

func printTree(root *TreeNode) {
	if root.Left != nil {
		printTree(root.Left)
	}
	print(root.Val)
	if root.Right != nil {
		printTree(root.Right)
	}

}
func main() {
	a := &TreeNode{Val: 1}
	b := &TreeNode{Val: 3}
	c := &TreeNode{Val: 2, Left: a, Right: b}
	d := &TreeNode{Val: 6}
	e := &TreeNode{Val: 9}
	f := &TreeNode{Val: 7, Left: d, Right: e}
	h := &TreeNode{Val: 4, Left: c, Right: f}
	printTree(h)
	println()
	printTree(invertTree(h))
}
