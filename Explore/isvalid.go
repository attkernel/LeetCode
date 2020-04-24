package main

import (
	"fmt"
)

var M = map[byte]int{
	'(': -1,
	')': 1,
	'{': -2,
	'}': 2,
	'[': -3,
	']': 3}

func isValid(s string) bool {
	Stack := make([]byte, 0)
	for _, v := range []byte(s) {
		if len(Stack) == 0 || M[Stack[len(Stack)-1]]+M[v] != 0 {
			Stack = append(Stack, v)
			continue
		}
		Stack = Stack[:len(Stack)-1]
	}
	return len(Stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
