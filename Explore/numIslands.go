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
