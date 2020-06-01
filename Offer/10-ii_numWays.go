/**
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：

输入：n = 2
输出：2
示例 2：

输入：n = 7
输出：21
提示：

0 <= n <= 100

**/
package main

func numWays(n int) int {
	M := make(map[int]int, 0)
	return recurs(n, M)
}

func recurs(n int, m map[int]int) int {
	if n == 0 {
		return 1
	}
	if n < 3 {
		return n
	}
	if _, ok := m[n]; !ok {
		m[n] = recurs(n-1, m) + recurs(n-2, m)
	}
	return m[n] % 1000000007
}

func main() {}
