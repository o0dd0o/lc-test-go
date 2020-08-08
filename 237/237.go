package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	for node.Next.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}
	node.Val = node.Next.Val
	node.Next = nil
}

func printNode(a *ListNode) {
	println(a.Val)
	if a.Next != nil {
		printNode(a.Next)
	}

}
func main() {
	a := &ListNode{Val: 9}
	b := &ListNode{Val: 1, Next: a}
	c := &ListNode{Val: 5, Next: b}
	d := &ListNode{Val: 4, Next: c}
	deleteNode(b)
	printNode(d)
}
