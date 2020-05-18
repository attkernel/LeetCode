package main

import (
	"fmt"
	"unsafe"
)

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
	s := make([]byte, 0)

	for k, _ := range S {
		if len(s) == 0 {
			s = append(s, S[k])
		} else if S[k] == s[len(s)-1] {
			s = s[:len(s)-1]
		} else {
			s = append(s, S[k])
		}
	}
	return *(*string)(unsafe.Pointer(&s))
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("abbbabaaa"))
}
