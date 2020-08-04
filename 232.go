package main

/**
使用栈实现队列的下列操作：

push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。


示例:

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false


说明:

你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type MyQueue struct {
	s1 []int
	s2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{s1: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1 = append(this.s1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.s2 = []int{}
	for len(this.s1) > 0 {
		this.s2 = append(this.s2, this.s1[len(this.s1)-1])
		this.s1 = this.s1[:len(this.s1)-1]
	}
	v := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	for len(this.s2) > 0 {
		this.s1 = append(this.s1, this.s2[len(this.s2)-1])
		this.s2 = this.s2[:len(this.s2)-1]
	}
	return v
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.s2 = []int{}
	for len(this.s1) > 0 {
		this.s2 = append(this.s2, this.s1[len(this.s1)-1])
		this.s1 = this.s1[:len(this.s1)-1]
	}
	v := this.s2[len(this.s2)-1]
	for len(this.s2) > 0 {
		this.s1 = append(this.s1, this.s2[len(this.s2)-1])
		this.s2 = this.s2[:len(this.s2)-1]
	}
	return v
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
func main() {
	obj := Constructor()
	obj.Push(3)
	obj.Push(4)
	param_2 := obj.Peek()
	param_3 := obj.Pop()
	param_4 := obj.Empty()
	println(param_2, param_3, param_4)
}
