/*
统计一个数字在排序数组中出现的次数。



示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0


限制：

0 <= 数组长度 <= 50000

*/
package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start, end, res := 0, len(nums)-1, 0
	for start < end {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		for i := start; i < len(nums); i++ {
			if nums[i] == target {
				res++
			}
		}
	}
	return res
}

func main() {}
