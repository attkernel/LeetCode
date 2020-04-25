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
