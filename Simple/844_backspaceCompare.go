package main

import (
	"fmt"
	"strings"
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

func OutPutStr(s string) string {
	Stack := NewStack(len(s))
	for k, _ := range s {
		if s[k] != '#' {
			Stack.Push(s[k])
			continue
		}
		Stack.Pop()
	}
	return string(Stack.GetData())
}

func backspaceCompare(S string, T string) bool {
	return strings.Compare(OutPutStr(S), OutPutStr(T)) == 0
}

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a##c", "#a#c"))
	fmt.Println(backspaceCompare("a#c", "b"))
}
