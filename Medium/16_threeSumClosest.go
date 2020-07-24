/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。



示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。


提示：

3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4

*/
package main

import "sort"

func ZVal(x int) int {
	if x > 0 {
		return x
	} else {
		return 0 - x
	}
}
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return -1
	}
	var val int = nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for k := 0; k < len(nums)-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		for i, j := k+1, len(nums)-1; i < j; {
			s := nums[i] + nums[j] + nums[k]
			if ZVal(target-val) > ZVal(target-s) {
				val = s
			}
			if s == target {
				return target
			} else if s < target {
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				i++
			} else if s > target {
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				j--
			}
		}
	}
	return val
}

func main() {}
