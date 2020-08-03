/*
给你一个仅包含小写字母的字符串，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。



示例 1:

输入: "bcabc"
输出: "abc"
示例 2:

输入: "cbacdcbc"
输出: "acdb"

*/
package main

func removeDuplicateLetters(s string) string {
	var result string
	lastIndex := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		lastIndex[s[i]] = i
	}
	used := make(map[uint8]bool)
	for i := 0; i < len(s); i++ {
		if used[s[i]] {
			continue
		}
		for len(result) != 0 && result[len(result)-1] > s[i] && lastIndex[result[len(result)-1]] > i {
			used[result[len(result)-1]] = false
			result = result[:len(result)-1]
		}
		used[s[i]] = true
		result += string(s[i])
	}
	return result
}

func main() {}
