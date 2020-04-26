/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。



示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

**/
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	M := make(map[int]int, 0)
	for k, _ := range nums {
		diff := target - nums[k]
		value, ok := M[diff]
		if ok {
			return []int{value, k}
		}
		M[nums[k]] = k
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
