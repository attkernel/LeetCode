package main

import (
	"fmt"
)

type MyStack struct {
	S []byte
}

func NewStack(len int) MyStack {
	return MyStack{
		S: make([]byte, 0, len),
	}
}

func (this *MyStack) Empty() bool {
	return len(this.S) == 0
}

func (this *MyStack) Peek() byte {
	if this.Empty() {
		return ' '
	}
	return this.S[len(this.S)-1]
}

func (this *MyStack) Pop() byte {
	if this.Empty() {
		return ' '
	}
	value := this.S[len(this.S)-1]
	this.S = this.S[:len(this.S)-1]
	return value
}

func (this *MyStack) Push(value byte) {
	this.S = append(this.S, value)
}

func (this *MyStack) GetData() []byte {
	return this.S
}

/*func removeDuplicates(S string) string {
	var Tmpbyte byte
	Stack := NewStack(len(S))
	for k, _ := range S {
		if Stack.Empty() {
			Stack.Push(S[k])
			Tmpbyte = S[k]
			continue
		}
		if Stack.Peek() == S[k] {
			Tmpbyte = Stack.Pop()
			continue
		}
		if Tmpbyte == S[k] {
			continue
		}
		Stack.Push(S[k])
		Tmpbyte = S[k]
	}
	return string(Stack.GetData())
}*/
func removeDuplicates(S string) string {
	Stack := NewStack(len(S))
	for k, _ := range S {
		if Stack.Empty() {
			Stack.Push(S[k])
			continue
		}
		if Stack.Peek() == S[k] {
			Stack.Pop()
			continue
		}
		Stack.Push(S[k])
	}
	return string(Stack.GetData())
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("abbbabaaa"))
}
