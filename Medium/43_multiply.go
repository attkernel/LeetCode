/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

*/
package main

import "strings"

func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	pos := 0
	res := make([]byte, l1+l2+1) // 可能向前进位，多保留一位
	for i := range res {         // 初始化res
		res[i] = '0'
	}
	for i := l1 - 1; i >= 0; i-- { // 需要用一个数乘以另一个数的每一位，因此两层遍历
		for j := l2 - 1; j >= 0; j-- {
			temp := (num1[i] - '0') * (num2[j] - '0')
			pos = i + j + 2
			res[pos-1] += (temp + res[pos] - '0') / 10 // 进位
			res[pos] = (temp+res[pos]-'0')%10 + '0'
		}
	}
	ret := strings.TrimLeft(string(res), "0") // 移除左侧多余的0
	if ret == "" {
		return "0"
	}
	return ret
}

func main() {}
