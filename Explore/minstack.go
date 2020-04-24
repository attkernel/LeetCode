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
