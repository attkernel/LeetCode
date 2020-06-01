/**
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。



示例 1：

输入：n = 2
输出：1
示例 2：

输入：n = 5
输出：5


提示：

0 <= n <= 100

**/

package main

import (
	"fmt"
)

func fib(n int) int {
	M := make(map[int]int, 0)
	return recurs(n, M)
}

func recurs(n int, m map[int]int) int {
	if n < 2 {
		return n
	}
	if _, ok := m[n]; !ok {
		m[n] = recurs(n-1, m) + recurs(n-2, m)
	}
	return m[n] % 1000000007
}
func main() {
	fmt.Println(fib(95))
}
