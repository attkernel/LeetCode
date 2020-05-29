/**
给定一个排序好的数组，两个整数 k 和 x，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。如果有两个数与 x 的差值一样，优先选择数值较小的那个数。

示例 1:

输入: [1,2,3,4,5], k=4, x=3
输出: [1,2,3,4]


示例 2:

输入: [1,2,3,4,5], k=4, x=-1
输出: [1,2,3,4]


说明:

k 的值为正数，且总是小于给定排序数组的长度。
数组不为空，且长度不超过 104
数组里的每个元素与 x 的绝对值不超过 104


更新(2017/9/19):
这个参数 arr 已经被改变为一个整数数组（而不是整数列表）。 请重新加载代码定义以获取最新更改。

**/
package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	var (
		mid  int
		prev int
		next int
		tmp  int
	)
	if k < 0 {
		return arr[:k]
	}
	if k > len(arr)-1 {
		return arr[len(arr)-k:]
	}
	tmp = k
	start, end := 0, len(arr)-1
	for start < end {
		mid := (start + end) / 2
		if arr[mid] == x {
			break
		} else if arr[mid] < x {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if arr[mid] == x {
		prev, next = mid, mid
	} else {
		prev, next = start-1, start
	}
	for {
		if prev < 0 {
			return arr[:k]
		}
		if x-arr[prev] <= arr[next]-x {
			prev--
			tmp--
		} else {
			next++
			tmp--
		}
		if prev < 0 {
			return arr[:k]
		}
		if next > len(arr)-1 {
			return arr[len(arr)-k:]
		}
		if next-prev-1 == k {
			return arr[prev+1 : next]
		}
	}
}

func main() {
	/*fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1))
	fmt.Println(findClosestElements([]int{1}, 1, 1))*/
	//fmt.Println(findClosestElements([]int{1, 1, 1, 10, 10, 10}, 1, 9))
	fmt.Println(findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5))
}
