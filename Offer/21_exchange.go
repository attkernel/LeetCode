/**
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。



示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。


提示：

1 <= nums.length <= 50000
1 <= nums[i] <= 10000

**/
package main

import (
	"fmt"
)

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	start, end := 0, len(nums)-1
	for {
		for nums[start]%2 == 1 && start < end {
			start++
		}
		for nums[end]%2 == 0 && end > start {
			end--
		}
		if start >= end {
			return nums
		}
		nums[start], nums[end] = nums[end], nums[start]
	}
}

func main() {
	fmt.Println(exchange([]int{1, 3, 5}))
}
