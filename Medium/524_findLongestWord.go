/*
给定一个字符串和一个字符串字典，找到字典里面最长的字符串，该字符串可以通过删除给定字符串的某些字符来得到。如果答案不止一个，返回长度最长且字典顺序最小的字符串。如果答案不存在，则返回空字符串。

示例 1:

输入:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

输出:
"apple"
示例 2:

输入:
s = "abpcplea", d = ["a","b","c"]

输出:
"a"
说明:

所有输入的字符串只包含小写字母。
字典的大小不会超过 1000。
所有输入的字符串长度不会超过 1000。

*/
package main

import (
	"fmt"
)

func findLongestWord(s string, d []string) string {
	res, l := "", 0
	for _, v := range d {
		if canGet(s, v) {
			if res == "" {
				res = v
				l = len(v)
				continue
			}
			if len(v) > l {
				l = len(v)
				res = v
			} else if len(v) == l {
				if v < res {
					res = v
				}
			}
		}
	}
	return res
}
func canGet(s string, sub string) bool {
	if len(sub) <= 0 {
		return true
	}
	p1 := 0
	p2 := 0
	for p1 < len(s) {
		if s[p1] == sub[p2] {
			p1++
			p2++
		} else {
			p1++
		}
		if p2 == len(sub) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}))
}
