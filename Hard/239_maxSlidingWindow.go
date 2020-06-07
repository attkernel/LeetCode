/*
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：

你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

*/

package main

import (
	"math"
)

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	res, max, pSlow, pFast := make([]int, 0), math.MinInt64, 0, k-1
	for i := 0; i < k; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	res = append(res, max)
	for pFast < len(nums)-1 {
		pSlow++
		pFast++
		if nums[pFast] > max {
			max = nums[pFast]
		}
		if nums[pSlow-1] == max {
			max = math.MinInt64
			for i := pSlow; i <= pFast; i++ {
				if nums[i] > max {
					max = nums[i]
				}
			}
		}
		res = append(res, max)
	}
	return res
}

func main() {}
