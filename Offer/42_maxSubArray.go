/*
输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。



示例1:

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。


提示：

1 <= arr.length <= 10^5
-100 <= arr[i] <= 100

*/
package main

import (
	"math"
)

func maxSubArray(nums []int) int {
	sumMax, preSum := math.MinInt64, math.MinInt64
	for idx := range nums {
		if preSum > 0 {
			preSum = preSum + nums[idx]
		} else {
			preSum = nums[idx]
		}
		if preSum > sumMax {

			sumMax = preSum
		}
	}
	return sumMax
}

func main() {}
