package main

/**
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
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

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast, top, tm := head, head, &ListNode{}, head

	for fast != nil {
		if fast.Next == nil {
			slow = slow.Next
			break
		}
		fast = fast.Next.Next
		tm = slow
		slow = slow.Next
		tm.Next = top
		top = tm
	}

	for ; slow != nil; slow, top = slow.Next, top.Next {
		if slow.Val != top.Val {
			return false
		}
	}

	return true
}
func printNode(head *ListNode) {
	print(head.Val, " ")
	if head.Next != nil {
		printNode(head.Next)
	}
}
func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2, Next: a}
	c := &ListNode{Val: 1, Next: b}
	println(isPalindrome(c))
}
