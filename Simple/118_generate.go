/**
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

**/
package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	Res := make([][]int, 0, numRows)
	recurs(1, numRows, Res)
	return Res[:numRows]
}

func recurs(current, numRows int, res [][]int) {
	if len(res) == numRows {
		return
	}
	rowData := make([]int, 0, numRows)
	for i := 0; i < current; i++ {
		if i == 0 || i == current-1 {
			rowData = append(rowData, 1)
			continue
		}
		rowData = append(rowData, res[current-1-1][i-1]+res[current-1-1][i])
	}
	res = append(res, rowData)
	recurs(current+1, numRows, res)
}

func main() {
	fmt.Println(generate(5))
}
