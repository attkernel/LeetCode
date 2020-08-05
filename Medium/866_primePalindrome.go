/*
求出大于或等于 N 的最小回文素数。

回顾一下，如果一个数大于 1，且其因数只有 1 和它自身，那么这个数是素数。

例如，2，3，5，7，11 以及 13 是素数。

回顾一下，如果一个数从左往右读与从右往左读是一样的，那么这个数是回文数。

例如，12321 是回文数。



示例 1：

输入：6
输出：7
示例 2：

输入：8
输出：11
示例 3：

输入：13
输出：101


提示：

1 <= N <= 10^8
答案肯定存在，且小于 2 * 10^8。

*/

package main

import (
	"math"
	"strconv"
)

func primePalindrome(N int) int {
	if N >= 8 && N <= 11 {
		return 11
	}
	for i := 1; i < 1e5; i++ {
		s := strconv.Itoa(i)
		t := []byte(s)
		l := len(t)
		for j := 0; j < l>>1; j++ {
			t[j], t[l-j-1] = t[l-j-1], t[j]
		}
		k, _ := strconv.Atoi(s + string(t[1:]))
		if k >= N && isPrime(k) {
			return k
		}
	}
	return -1
}

func isPrime(n int) bool {
	if n < 2 || n&1 == 0 {
		return n == 2
	}
	s := int(math.Sqrt(float64(n)))
	for i := 3; i <= s; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {}
