/**
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

**/

package main

import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	Res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start, end := i+1, len(nums)-1
		for start < end {
			if start > i+1 && nums[start] == nums[start-1] {
				start++
				continue
			}
			sum := nums[i] + nums[start] + nums[end]
			if sum == 0 {
				Res = append(Res, []int{nums[i], nums[start], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return Res
}

func main() {
	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//fmt.Println(threeSum([]int{0, 0, 0, 0}))
	//fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Println(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
}
