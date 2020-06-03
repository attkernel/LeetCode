/**
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。



你可以假设数组是非空的，并且给定的数组总是存在多数元素。



示例 1:

输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2


限制：

1 <= 数组长度 <= 50000

**/

package main

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	votes, res := 0, nums[0]
	for idx, _ := range nums {
		if votes == 0 {
			res = nums[idx]
		}
		if nums[idx] == res {
			votes++
		} else {
			votes--
		}
	}
	return res
}

func main() {}
