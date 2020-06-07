/*
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

示例:

s = "abaccdeff"
返回 "b"

s = ""
返回 " "


限制：

0 <= s 的长度 <= 50000

*/
package main

func firstUniqChar(s string) byte {
	var S [26]int
	for idx := range s {
		S[s[idx]-'a']++
	}
	for idx := range s {
		if S[s[idx]-'a'] == 1 {
			return s[idx]
		}
	}
	return ' '
}

func main() {}
