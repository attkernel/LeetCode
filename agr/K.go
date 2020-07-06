package main

import (
	"fmt"
)

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
	M, count := make(map[int]struct{}), 0
	quickSort(customers)
	quickSort(supports)
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
	fmt.Println(F([]int{5, 7, 9}, []int{6, 8, 10}))
	fmt.Println(F([]int{5, 7, 9}, []int{3, 3, 5, 3, 3}))
}
