/**
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

**/
package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isValid(s[i]) && !isValid(s[j]) {
			i++
			j--
		} else if !isValid(s[i]) {
			i++
		} else if !isValid(s[j]) {
			j--
		} else if !equ(s[i], s[j]) {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func isValid(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func equ(b1, b2 byte) bool {
	if b1 >= 'a' && b1 <= 'z' {
		if b2 >= 'a' && b2 <= 'z' {
			return b1 == b2
		}
		if b2 >= 'A' && b2 <= 'Z' {
			return (int8(b1) - 32) == int8(b2)
		}
	} else if b1 >= 'A' && b1 <= 'Z' {
		if b2 >= 'A' && b2 <= 'Z' {
			return b1 == b2
		}
		if b2 >= 'a' && b2 <= 'z' {
			return (int8(b2) - 32) == int8(b1)
		}
	}
	return b1 == b2
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("0P"))
}
