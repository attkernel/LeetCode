/*
给出一个字符串 s（仅含有小写英文字母和括号）。

请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。

注意，您的结果中 不应 包含任何括号。



示例 1：

输入：s = "(abcd)"
输出："dcba"
示例 2：

输入：s = "(u(love)i)"
输出："iloveu"
示例 3：

输入：s = "(ed(et(oc))el)"
输出："leetcode"
示例 4：

输入：s = "a(bcdefghijkl(mno)p)q"
输出："apmnolkjihgfedcbq"


提示：

0 <= s.length <= 2000
s 中只有小写英文字母和括号
我们确保所有括号都是成对出现的

*/
package main

import "strings"

func reverseParentheses(s string) string {
	var stack []int
	bytes := []byte(s)
	for k, v := range s {
		if v == '(' {
			stack = append(stack, k)
		}
		if v == ')' {
			left := stack[len(stack)-1]
			reverseBytes1190(bytes, left, k)
			stack = stack[:len(stack)-1]
		}
	}
	var res strings.Builder
	for _, v := range bytes {
		if v != '(' && v != ')' {
			res.WriteByte(v)
		}
	}
	return res.String()
}

func reverseBytes1190(arr []byte, i int, k int) {
	for i < k {
		arr[i], arr[k] = arr[k], arr[i]
		i++
		k--
	}
}

func main() {}
