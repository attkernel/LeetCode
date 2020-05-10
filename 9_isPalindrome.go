/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？

*
*/
package main

import (
	"fmt"
)

/*func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	Slice := make([]int, 0)
	for {
		if x%10 == x {
			Slice = append(Slice, x)
			break
		}
		Slice = append(Slice, x%10)
		x = x / 10
	}
	for i, j := 0, (len(Slice) - 1); i <= j; i, j = i+1, j-1 {
		if Slice[i] != Slice[j] {
			return false
		}
	}
	return true
}*/
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reverseNum := 0
	for x > reverseNum {
		reverseNum = reverseNum*10 + x%10
		x = x / 10
	}
	return x == reverseNum || x == reverseNum/10
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
