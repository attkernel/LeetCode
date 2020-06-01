/**
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3


限制：

2 <= n <= 100000

**/
package main

import (
	"fmt"
)

/*func findRepeatNumber(nums []int) int {
	M := make(map[int]bool, 0)
	for idx, _ := range nums {
		if _, ok := M[nums[idx]]; ok {
			return nums[idx]
		}
		M[nums[idx]] = true
	}
	return -1
}*/
func findRepeatNumber(nums []int) int {
	for k, v := range nums {
		if k == v {
			continue
		}
		if nums[v] == v {
			return v
		}
		nums[v], nums[k] = nums[k], nums[v]
	}
	return -1
}

func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(findRepeatNumber([]int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}))
}
