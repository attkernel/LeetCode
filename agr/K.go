package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//快排
func quickSort(src []int) {
	_quickSort(src, 0, len(src)-1)
}

func _quickSort(src []int, left, right int) {
	if left < right {
		partitionIndex := partition(src, left, right)
		_quickSort(src, left, partitionIndex-1)
		_quickSort(src, partitionIndex+1, right)
	}
}

func partition(src []int, left, right int) int {
	pivot := left
	index := left + 1
	for i := index; i <= right; i++ {
		if src[i] < src[pivot] {
			src[i], src[index] = src[index], src[i]
			index++
		}
	}
	src[pivot], src[index-1] = src[index-1], src[pivot]
	return index - 1
}

func F(customers, supports []int) int {
	//M作为hash表去重
	M, count := make(map[int]struct{}), 0
	//通过快排将客户表和技术支持表排升序
	quickSort(customers)
	quickSort(supports)
	//将客户需要能力值在技术支持能力值里查找最接近的值
	for i, _ := range customers {
		for j, _ := range supports {
			if _, ok := M[i]; !ok && supports[j] >= customers[i] {
				M[j] = struct{}{}
				count++
			}
		}
	}
	return count
}

func main() {
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		log.Fatal("The command is intended to work with pipes.")
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		strs := strings.Split(s.Text(), " ")
		if len(strs) == 1 {
			continue
		} else if len(strs) == 2 {
			cus, sup := make([]int, 0), make([]int, 0)
			s.Scan()
			strsCus := strings.Split(s.Text(), " ")
			for idx, _ := range strsCus {
				v, err := strconv.Atoi(strsCus[idx])
				if err != nil {
					continue
				}
				cus = append(cus, v)
			}
			s.Scan()
			strsSup := strings.Split(s.Text(), " ")
			for idx, _ := range strsSup {
				v, err := strconv.Atoi(strsSup[idx])
				if err != nil {
					continue
				}
				sup = append(sup, v)
			}
			fmt.Println(F(cus[:], sup[:]))
		}
	}
	/*fmt.Println(F([]int{5, 7, 9}, []int{6, 8, 10}))
	fmt.Println(F([]int{5, 7, 9}, []int{3, 3, 5, 3, 3}))*/
}
