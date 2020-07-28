/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	top, tm := &ListNode{}, &ListNode{}
	for head != nil {
		tm = head
		head = head.Next
		tm.Next = top.Next
		top.Next = tm
	}
	return top.Next
}

func main() {
	//1->2->3->4->5->NULL
	//1<-t 2->3->4->5->NULL
	//1<-2<-t 3->4->5->NULL
	a := &ListNode{Val: 5}
	b := &ListNode{Val: 4, Next: a}
	c := &ListNode{Val: 3, Next: b}
	d := &ListNode{Val: 2, Next: c}
	f := &ListNode{Val: 1, Next: d}
	f_ := reverseList(f)
	for f_ != nil {
		fmt.Printf("%d\n", f_.Val)
		f_ = f_.Next
	}
}
