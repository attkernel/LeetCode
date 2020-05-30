/**
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。

示例:

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
进阶:

如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。

**/

package main

import (
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	start, Len, sum := 0, math.MaxInt64, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= s {
			if i-start+1 < Len {
				Len = i - start + 1
			}
			sum -= nums[start]
			start++
		}
	}
	if Len == math.MaxInt64 {
		Len = 0
	}
	return Len
}

func main() {}
