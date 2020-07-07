package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var once sync.Once

func main() {
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		log.Fatal("The command is intended to work with pipes.")
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		once.Do(func() {
			s.Scan()
		})
		s.Scan()
		src := make([]int, 0)
		strs := strings.Split(s.Text(), " ")
		for idx, _ := range strs {
			v, err := strconv.Atoi(strs[idx])
			if err != nil {
				os.Exit(-1)
			}
			src = append(src, v)
		}
		fmt.Println(F(src[:]))
	}
	/*fmt.Println(F(5, []int{3, 1, 2, 3, 1}))
	fmt.Println(F(7, []int{1, 1, 1, 1, 1, 2, 2}))*/
}

func F(src []int) int {
	sum, m, count := 0, make(map[int]struct{}, 0), 0
	//m作为hash表记录已经遍历过的元素下标
	//数组求和与3取余 数组元素删除掉取余结果后 余下元素肯定能合并成被3整除的数
	for idx, _ := range src {
		sum += src[idx]
	}
	for idx, _ := range src {
		if src[idx]%3 == sum%3 && sum%3 != 0 {
			m[idx] = struct{}{}
			break
		}
	}
	//当m的长度等于src长度时退出循环代表所有元素都已经遍历了一次
	for len(m) < len(src) {
		//外层遍历
		for i := 0; i < len(src); i++ {
			//如果之前遍历过则跳过
			if _, ok := m[i]; ok {
				continue
			}
			//如果当前元素能被3整除则将下标记录到hash表，count++
			if src[i]%3 == 0 {
				m[i] = struct{}{}
				count++
			} else {
				flag := false
				//当外层元素不能被3整除时，开启内层遍历,如果内外层元素相加能被3整除,则将两个下标添加到hash表 count++
				for j := i + 1; j < len(src); j++ {
					if _, ok := m[j]; !ok && ((src[i]+src[j])%3 == 0) {
						m[i] = struct{}{}
						m[j] = struct{}{}
						count++
						flag = true
						break
					}
				}
				//当内外层元素相加不能被3整除时,说明两个元素和为2,则将和传给外层下标，内层下标加入hash表，进行下次遍历
				if !flag {
					for j := i + 1; j < len(src); j++ {
						if _, ok := m[j]; !ok {
							src[i] += src[j]
							m[j] = struct{}{}
							break
						}
					}
				}
			}
		}
	}
	return count
}
