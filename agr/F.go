package main

import (
	"fmt"
)

func main() {
	fmt.Println(F(5, []int{3, 1, 2, 3, 1}))
	fmt.Println(F(7, []int{1, 1, 1, 1, 2, 2}))
}

func F(n int, src []int) int {
	sum, m, count, flag := 0, make(map[int]struct{}, 0), 0, false
	for idx, _ := range src {
		sum += src[idx]
	}
	for idx, _ := range src {
		if src[idx]%3 == sum%3 {
			src = append(src[:idx], src[idx+1:]...)
			break
		}
	}
	for len(m) < len(src) {
		for i := 0; i < len(src); i++ {
			if _, ok := m[i]; !ok && src[i]%3 == 0 {
				m[i] = struct{}{}
				count++
			} else {
				for j := i + 1; j < len(src); j++ {
					if _, ok := m[j]; !ok && ((src[i]+src[j])%3 == 0) {
						m[i] = struct{}{}
						m[j] = struct{}{}
						count++
						flag = true
						break
					}
				}
				if !flag {
					for j := i + 1; j < len(src); j++ {
						if _, ok := m[j]; !ok {
							src[i] += src[j]
							m[j] = struct{}{}
							flag = true
							break
						}
					}
				}
			}
		}
	}
	return count
}
