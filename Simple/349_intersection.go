/**
给定两个数组，编写一个函数来计算它们的交集。

示例 1:

输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]
说明:

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。

**/

package main

func intersection(nums1 []int, nums2 []int) []int {
	Res := make([]int, 0)
	M := make(map[int]bool, 0)
	for idx, _ := range nums1 {
		M[nums1[idx]] = true
	}
	for idx, _ := range nums2 {
		if value, ok := M[nums2[idx]]; ok && value {
			Res = append(Res, nums2[idx])
			M[nums2[idx]] = false
		}
	}
	return Res
}

func main() {}
