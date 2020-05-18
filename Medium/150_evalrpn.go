/**
根据逆波兰表示法，求表达式的值。

有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：

整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
示例 1：

输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: ((2 + 1) * 3) = 9
示例 2：

输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: (4 + (13 / 5)) = 6
示例 3：

输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22

**/
package main

import (
	"errors"
	"strconv"
)

type f func(first, second int) (int, error)

var M = map[string]f{
	"+": Add,
	"-": Sub,
	"*": Mul,
	"/": Div,
}

func Add(first, second int) (int, error) {
	return (first + second), nil
}

func Sub(first, second int) (int, error) {
	return (first - second), nil
}

func Mul(first, second int) (int, error) {
	return (first * second), nil
}

func Div(first, second int) (int, error) {
	if second == 0 {
		return -1, errors.New("error")
	}
	return (first / second), nil
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for k, _ := range tokens {
		if _, ok := M[tokens[k]]; !ok {
			numInt, err := strconv.Atoi(tokens[k])
			if err != nil {
				return -1
			}
			stack = append(stack, numInt)
			continue
		}
		if len(stack) < 2 {
			return -1
		}
		p1 := stack[len(stack)-1]
		p2 := stack[len(stack)-2]
		if result, err := M[tokens[k]](p2, p1); err == nil {
			stack[len(stack)-2] = result
			stack = stack[:len(stack)-1]
			continue
		}
		return -1
	}
	return stack[len(stack)-1]
}

func main() {
	println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
