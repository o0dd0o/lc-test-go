package main

import (
	"container/list"
)

/**
使用队列实现栈的下列操作：

push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空
注意:

你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-stack-using-queues
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type MyStack struct {
	l1 *list.List
	l2 *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{l1: list.New(), l2: list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.l1.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */

func (this *MyStack) Pop() int {
	for this.l1.Front().Next() != nil {
		this.l2.PushBack(this.l1.Remove(this.l1.Front()))
	}
	head := this.l1.Front().Value.(int)
	this.l1.Remove(this.l1.Front())
	this.l1, this.l2 = this.l2, this.l1
	return head
}

/** Get the top element. */
func (this *MyStack) Top() int {
	for this.l1.Front().Next() != nil {
		this.l2.PushBack(this.l1.Remove(this.l1.Front()))
	}

	head := this.l1.Front().Value.(int)
	this.l2.PushBack(this.l1.Remove(this.l1.Front()))
	this.l1, this.l2 = this.l2, this.l1
	return head
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.l1.Front() == nil
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
func main() {
	obj := Constructor()
	obj.Push(3)
	obj.Push(4)
	param_2 := obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Empty()
	println(param_2, param_3, param_4)
}
