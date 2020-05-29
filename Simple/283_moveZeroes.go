/**
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

**/
package main

func moveZeroes(nums []int) {
	idxZ, idxNZ := 0, 0
	for ; idxNZ < len(nums); idxNZ++ {
		if nums[idxNZ] != 0 {
			if nums[idxZ] == 0 {
				nums[idxZ], nums[idxNZ] = nums[idxNZ], nums[idxZ]
			}
			idxZ++
		}
	}
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{1})
	moveZeroes([]int{0, 0})
	moveZeroes([]int{1, 0})
}
