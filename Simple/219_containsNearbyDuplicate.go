/*
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。



示例 1:

输入: nums = [1,2,3,1], k = 3
输出: true
示例 2:

输入: nums = [1,0,1,1], k = 1
输出: true
示例 3:

输入: nums = [1,2,3,1,2,3], k = 2
输出: false

*/
package main

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 1; i <= k; i++ {
		if len(nums) >= i+1 {
			start, end := 0, i
			for end <= len(nums)-1 {
				if nums[start] == nums[end] {
					return true
				}
				start++
				end++
			}
		}
	}
	return false
}

func main() {}
