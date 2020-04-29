package main

import (
	"fmt"
)

type MyStack struct {
	S []int
}

func NewStack(len int) MyStack {
	return MyStack{
		S: make([]int, 0, len),
	}
}

func (this *MyStack) Empty() bool {
	return len(this.S) == 0
}

func (this *MyStack) Peek() int {
	if this.Empty() {
		return -1
	}
	return this.S[len(this.S)-1]
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}
	value := this.S[len(this.S)-1]
	this.S = this.S[:len(this.S)-1]
	return value
}

func (this *MyStack) Push(value int) {
	this.S = append(this.S, value)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	Result := make([]int, 0, len(nums1))
	Stack := NewStack(len(nums2))
	MaxMap := make(map[int]int, len(nums2))
	for k, _ := range nums2 {
		MaxMap[nums2[k]] = -1
		if Stack.Empty() {
			Stack.Push(nums2[k])
		} else {
			for nums2[k] > Stack.Peek() && !Stack.Empty() {
				MaxMap[Stack.Pop()] = nums2[k]
			}
			Stack.Push(nums2[k])
		}
	}
	for k, _ := range nums1 {
		Result = append(Result, MaxMap[nums1[k]])
	}
	return Result
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
	fmt.Println(nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}))
}
