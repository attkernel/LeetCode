/*
给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。

示例 1:
输入:

0 0 0
0 1 0
0 0 0
输出:

0 0 0
0 1 0
0 0 0
示例 2:
输入:

0 0 0
0 1 0
1 1 1
输出:

0 0 0
0 1 0
1 2 1
注意:

给定矩阵的元素个数不超过 10000。
给定矩阵中至少有一个元素是 0。
矩阵中的元素只在四个方向上相邻: 上、下、左、右。

*/
package main

var direction [][]int
var result [][]int

func updateMatrix(matrix [][]int) [][]int {
	result = makeEmptyBoard(len(matrix), len(matrix[0]))

	direction = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// used:=makeEmptyBoard(len(matrix),len(matrix[0]))
			curPoint := matrix[i][j]
			if curPoint == 0 {
				result[i][j] = 0
				continue
			}
			queue := make([][]int, 0)
			queue = append(queue, []int{i, j})
			times := 0
			for len(queue) > 0 {
				ifFound := false
				preLen := len(queue)
				for k := 0; k < preLen; k++ {
					curr := queue[k]
					if matrix[curr[0]][curr[1]] == 0 {
						result[i][j] = times
						ifFound = true
						break
					}

					for m := 0; m < len(direction); m++ {
						d := direction[m]
						x := curr[0] + d[0]
						y := curr[1] + d[1]
						if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) {
							queue = append(queue, []int{x, y})

						}
					}
				}
				if ifFound {
					break
				}
				queue = queue[preLen:]
				times++
			}
		}
	}
	return result
}

func makeEmptyBoard(width int, height int) [][]int {
	temp := make([][]int, width)
	for i := 0; i < width; i++ {
		temp[i] = make([]int, height)
	}
	return temp
}

func main() {}
