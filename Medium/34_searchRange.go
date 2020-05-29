/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

**/
package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	Res := [2]int{-1, -1}
	if len(nums) == 0 {
		return Res[:]
	}
	for start < end {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		Res[0] = start
		for nums[start] == target {
			start++
			if start > len(nums)-1 {
				break
			}
		}
		Res[1] = start - 1
	}
	return Res[:]
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}
