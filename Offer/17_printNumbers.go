/**
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:

输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]


说明：

用返回一个整数列表来代替打印
n 为正整数

**/
package main

func printNumbers(n int) []int {
	Res := make([]int, 0)
	max := 0
	for n > 0 {
		max = max*10 + 9
		n--
	}
	for i := 1; i <= max; i++ {
		Res = append(Res, i)
	}
	return Res
}

func main() {}
