package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func F(n, m int, times []int) int {
	//接水口数大于人数直接接
	if m > n {
		return 0
	}
	//接水口数等于人数，需要等待人数中接水时间最少的时间
	if m == n {
		min := math.MaxInt64
		for idx, _ := range times {
			if times[idx] < min {
				min = times[idx]
			}
		}
		return min
	}
	//接水人可看做队列M是接水人和接水时间的对应关系
	M := make(map[int]int, n)
	for idx, _ := range times {
		M[idx+1] = times[idx]
	}
	//初始看做每个节水口存在一个人
	nums, count := m+1, 0
	for nums <= n+1 {
		//循环遍历接水口,每遍历一次等待时间+1,M[i]==0表示当前节水口需要换人
		for i := 1; i <= m; i++ {
			M[i]--
			if M[i] == 0 {
				M[i] = M[nums]
				nums++
			}
		}
		count++
	}
	return count + 1
}

func main() {
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		log.Fatal("The command is intended to work with pipes.")
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		n, m, src := 0, 0, make([]int, 0)
		str := strings.Split(s.Text(), " ")
		if len(str) != 2 {
			os.Exit(-1)
		}
		n, err := strconv.Atoi(str[0])
		if err != nil {
			os.Exit(-1)
		}
		m, err = strconv.Atoi(str[1])
		if err != nil {
			os.Exit(-1)
		}
		s.Scan()
		strSrc := strings.Split(s.Text(), " ")
		for idx, _ := range strSrc {
			v, err := strconv.Atoi(strSrc[idx])
			if err != nil {
				os.Exit(-1)
			}
			src = append(src, v)
		}
		fmt.Println(F(n, m, src[:]))
	}
	//fmt.Println(F(4, 2, []int{1, 1, 1, 1}))
}
