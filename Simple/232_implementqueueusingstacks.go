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

你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。

**/
package main

import (
	"fmt"
)

type MyStack struct {
	S []int
}

func NewStack() MyStack {
	return MyStack{
		S: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.S = append(this.S, x)
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}
	value := this.S[len(this.S)-1]
	this.S = this.S[:len(this.S)-1]
	return value
}

func (this *MyStack) Peek() int {
	if this.Empty() {
		return -1
	}
	return this.S[len(this.S)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.S) == 0
}

type MyQueue struct {
	In, Out *MyStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		In: &MyStack{
			S: make([]int, 0),
		},
		Out: &MyStack{
			S: make([]int, 0),
		},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.In.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Out.Empty() {
		for !this.In.Empty() {
			this.Out.Push(this.In.Pop())
		}
	}
	return this.Out.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Out.Empty() {
		for !this.In.Empty() {
			this.Out.Push(this.In.Pop())
		}
	}
	return this.Out.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.In.Empty() && this.Out.Empty()
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Empty())
}
