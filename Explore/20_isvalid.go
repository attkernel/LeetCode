/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

**/
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
