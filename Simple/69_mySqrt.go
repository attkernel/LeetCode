/**
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。

**/

package main

import (
	"fmt"
)

func mySqrt(x int) int {
	start, end := 0, x
	for start <= end {
		mid := (start + end) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	if start*start > x {
		start -= 1
	}
	return start
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}
