package main

/**
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]





示例 1:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。
示例 2:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。


说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉搜索树中。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parentmap := map[int]*TreeNode{}
	var dfs func(*TreeNode)
	dfs = func(a *TreeNode) {
		if a.Left != nil {
			parentmap[a.Left.Val] = a
			dfs(a.Left)
		}

		if a.Right != nil {
			parentmap[a.Right.Val] = a
			dfs(a.Right)
		}

	}
	if root == nil {
		return root
	}
	dfs(root)
	qmap := map[int]*TreeNode{}

	for t := q; t != nil; t = parentmap[t.Val] {
		qmap[t.Val] = t
	}
	t := p
	for ; qmap[t.Val] == nil; t = parentmap[t.Val] {

	}
	return t
}

func main() {
	a := &TreeNode{Val: 3}
	b := &TreeNode{Val: 5}
	c := &TreeNode{Val: 0}
	d := &TreeNode{Val: 7}
	f := &TreeNode{Val: 9}
	g := &TreeNode{Val: 4, Left: a, Right: b}
	h := &TreeNode{Val: 2, Left: c, Right: g}
	i := &TreeNode{Val: 8, Left: d, Right: f}
	j := &TreeNode{Val: 6, Left: h, Right: i}
	println(lowestCommonAncestor(j, h, i).Val, h.Val, i.Val)
	println(lowestCommonAncestor(j, h, g).Val, h.Val, g.Val)
}
