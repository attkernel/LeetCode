package main

import (
	"fmt"
)

func Fibonacci(n int) int {
	M := make(map[int]int)
	return F(n, M)
}

func F(n int, m map[int]int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	_, ok := m[n]
	if !ok {
		m[n] = F(n-1, m) + F(n-2, m)
	}
	return m[n]
}

func main() {
	fmt.Println(Fibonacci(1))
	fmt.Println(Fibonacci(2))
	fmt.Println(Fibonacci(3))
	fmt.Println(Fibonacci(4))
	fmt.Println(Fibonacci(5))
	fmt.Println(Fibonacci(6))
	fmt.Println(Fibonacci(7))
}
