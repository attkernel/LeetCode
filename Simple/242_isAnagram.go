/**
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

**/
package main

func isAnagram(s string, t string) bool {
	M := make(map[byte]int, 0)
	for idx, _ := range s {
		M[s[idx]]++
	}
	for idx, _ := range t {
		if _, ok := M[t[idx]]; !ok {
			return false
		}
		M[t[idx]]--
	}
	for _, v := range M {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {}
