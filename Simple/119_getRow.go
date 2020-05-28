/**
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 3
输出: [1,3,3,1]
进阶：

你可以优化你的算法到 O(k) 空间复杂度吗？

**/
package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	Res := make([][]int, 0, rowIndex+1)
	recurs(1, rowIndex+1, Res)
	tmp := Res[:rowIndex+1]
	return tmp[len(tmp)-1]
}

func recurs(current, k int, res [][]int) {
	if len(res) == k {
		return
	}
	rowData := make([]int, 0, k)
	for i := 0; i < current; i++ {
		if i == 0 || i == current-1 {
			rowData = append(rowData, 1)
			continue
		}
		rowData = append(rowData, res[current-2][i-1]+res[current-2][i])
	}
	res = append(res, rowData)
	recurs(current+1, k, res)
}

func main() {
	fmt.Println(getRow(3))
}
