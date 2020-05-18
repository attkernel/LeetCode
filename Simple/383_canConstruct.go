/**
给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串 ransom 能不能由第二个字符串 magazines 里面的字符构成。如果可以构成，返回 true ；否则返回 false。

(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。杂志字符串中的每个字符只能在赎金信字符串中使用一次。)



注意：

你可以假设两个字符串均只含有小写字母。

canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true

**/
package main

import (
	"fmt"
)

/*func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	M := make(map[byte]int, len(magazine))
	for idx, _ := range magazine {
		M[magazine[idx]]++
	}
	for idx, _ := range ransomNote {
		if M[ransomNote[idx]] == 0 {
			return false
		}
		M[ransomNote[idx]]--
	}
	return true
}*/
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	if ransomNote == "" {
		return true
	}
	if magazine == "" {
		return false
	}
	Slice := [26]int{}
	for idx, _ := range magazine {
		Slice[int8(magazine[idx]-'a')]++
	}
	for idx, _ := range ransomNote {
		if Slice[int8(ransomNote[idx]-'a')] == 0 {
			return false
		}
		Slice[int8(ransomNote[idx]-'a')]--
	}
	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("", ""))
}
