/**
请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。执行push、pop和min操作的时间复杂度必须为O(1)。


示例：

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

**/
package main

import (
	"fmt"
)

type MinStack struct {
	S   []int
	Min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		S:   make([]int, 0),
		Min: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if this.Empty() {
		this.S = append(this.S, x)
		this.Min = append(this.Min, x)
		return
	}
	this.S = append(this.S, x)
	if this.Min[len(this.Min)-1] >= x {
		this.Min = append(this.Min, x)
	}
}

func (this *MinStack) Pop() {
	if this.S[len(this.S)-1] == this.Min[len(this.Min)-1] {
		this.S = this.S[:len(this.S)-1]
		this.Min = this.Min[:len(this.Min)-1]
		return
	}
	this.S = this.S[:len(this.S)-1]
}

func (this *MinStack) Top() int {
	return this.S[len(this.S)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

func (this *MinStack) Empty() bool {
	return len(this.S) == 0
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}
