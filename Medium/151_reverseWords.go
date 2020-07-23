/*
给定一个字符串，逐个翻转字符串中的每个单词。



示例 1：

输入: "the sky is blue"
输出: "blue is sky the"
示例 2：

输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。


说明：

无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。


进阶：

请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。

*/
package main

import "strings"

func reverseWords(s string) string {
	if s == "" {
		return ""
	}
	s = PreProcess(s)
	wordarry := strings.Split(s, " ")
	i := 0
	l := len(wordarry) - 1
	for i < l { //
		wordarry[i], wordarry[l] = wordarry[l], wordarry[i]
		i++
		l--
	}
	res := strings.Join(wordarry, " ")
	return res
}

func PreProcess(s string) string {
	l := len(s)
	var res []byte
	flag := 1
	for l != 0 && s[l-1] == ' ' {
		l--
	}
	for i := 0; i < l; i++ {
		if s[i] != ' ' {
			res = append(res, s[i])
			flag = 0
		}
		if s[i] == ' ' && flag == 0 {
			res = append(res, s[i])
			flag = 1
		}
	}
	return string(res)
}

func main() {}
