package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Fibonacci(n int) int {
	M := make(map[int]int)
	return F(n, M)
}

//可以用dp实现,由于不会只能用从上到下的递归,为了减少重复计算,加入缓存
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
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		log.Fatal("The command is intended to work with pipes.")
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			os.Exit(-1)
		}
		fmt.Println(Fibonacci(v))
	}
	/*fmt.Println(Fibonacci(1))
	fmt.Println(Fibonacci(2))
	fmt.Println(Fibonacci(3))
	fmt.Println(Fibonacci(4))
	fmt.Println(Fibonacci(5))
	fmt.Println(Fibonacci(6))
	fmt.Println(Fibonacci(7))*/
}
