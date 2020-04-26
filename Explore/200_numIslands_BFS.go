/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

示例 1:

输入:
11110
11010
11000
00000
输出: 1

示例 2:

输入:
11000
11000
00100
00011
输出: 3
解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。

**/
package main

import (
	"fmt"
)

var (
	directionI = []int{0, 0, -1, 1}
	directionJ = []int{1, -1, 0, 0}
)

func numIslands(grid [][]byte) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	step := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				step = step + 1
				BFS(grid, row, col, i, j)
			}
		}
	}
	return step
}

func BFS(grid [][]byte, row, col, i, j int) {
	queue := make([]int, 0)
	queue = append(queue, i, j)
	grid[i][j] = '0'
	for len(queue) > 0 {
		currentI, currentJ := queue[0], queue[1]
		queue = queue[2:]
		for k := 0; k < 4; k++ {
			newI, newJ := currentI+directionI[k], currentJ+directionJ[k]
			if newI >= 0 && newI < row && newJ >= 0 && newJ < col && grid[newI][newJ] == '1' {
				grid[newI][newJ] = '0'
				queue = append(queue, newI, newJ)
			}
		}
	}
}

func main() {
	num := numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}})
	fmt.Println(num)
}
