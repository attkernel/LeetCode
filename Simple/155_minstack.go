/**
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。


示例:

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.


提示：

pop、top 和 getMin 操作总是在 非空栈 上调用。

**/
package main

import (
	"fmt"
)

type MinStack struct {
	Data []int
	Min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Data: make([]int, 0),
		Min:  make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.Data = append(this.Data, x)
	if len(this.Min) == 0 || this.Min[len(this.Min)-1] >= x {
		this.Min = append(this.Min, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.Data) == 0 || len(this.Min) == 0 {
		return
	}
	value := this.Data[len(this.Data)-1]
	this.Data = this.Data[:len(this.Data)-1]
	if value == this.Min[len(this.Min)-1] {
		this.Min = this.Min[:len(this.Min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.Data) == 0 {
		return -1
	}
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
