/**
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:

输入: "hello"
输出: "holle"
示例 2:

输入: "leetcode"
输出: "leotcede"
说明:
元音字母不包含字母"y"。

**/

package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	M := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	i, j, b := 0, len(s)-1, []byte(s)
	for j >= 0 && i <= len(s)-1 {
		if _, ok := M[b[i]]; !ok {
			i++
			continue
		}
		if _, ok := M[b[j]]; !ok {
			j--
			continue
		}
		if i >= j {
			return string(b)
		}
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return s
}

func main() {
	//fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels(" "))
}
