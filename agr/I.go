/*
Fibonacci
Description

斐波那契数列（Fibonacci sequence），又称黄金分割数列、因数学家列昂纳多·斐波那契（Leonardoda Fibonacci）以兔子繁殖为例子而引入，故又称为“兔子数列”，指的是这样一个数列：1、1、2、3、5、8、13、21、34、……在数学上，斐波那契数列以如下被以递推的方法定义

F(1)=1F(1)=1
F(2)=1F(2)=1
F(n)=F(n-1)+F(n-2)F(n)=F(n−1)+F(n−2)
其中 N*n≥3,n∈N∗
在现代物理、准晶体结构、化学等领域，斐波纳契数列都有直接的应用。

请编程实现 F(n)F(n) 函数


Input
一列正整数n， 由换行符分隔开；


Output
斐波那契函数F(n)F(n)的计算结果，由换行符分隔开；


Sample Input 1

1
2
3
4
5
6
Sample Output 1

1
1
2
3
5
8

Hint

输入中没有注明总数据的行数，需要自行加一个判断；
*/
package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			os.Exit(-1)
		}
		fmt.Println(Fibonacci(v))
	}
}
func Fibonacci(n int) *big.Int {
	M := make(map[int]*big.Int)
	return F(n, M)
}

//可以用dp实现,由于不会只能用从上到下的递归,为了减少重复计算,加入缓存
func F(n int, m map[int]*big.Int) *big.Int {
	if n == 1 {
		return big.NewInt(1)
	}
	if n == 0 {
		return big.NewInt(0)
	}
	_, ok := m[n]
	if !ok {
		tmp := big.NewInt(0)
		m[n] = tmp.Add(F(n-1, m), (F(n-2, m)))
	}
	return m[n]
}
