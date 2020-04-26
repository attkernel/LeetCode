package main

import "fmt"

var (
	directionI = []int{0, 0, -1, 1}
	directionJ = []int{1, -1, 0, 0}
)

type Node struct {
	Row int
	Col int
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return -1
	}
	rows := len(grid)
	cols := len(grid[0])
	visited := make(map[Node]int8)
	nums := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if _, ok := visited[Node{i, j}]; !ok && grid[i][j] == '1' {
				visited[Node{i, j}] = 0
				DFS(grid, i, j, rows, cols, visited)
				nums++
			}
		}
	}
	return nums
}

func DFS(grid [][]byte, i, j, rows, cols int, v map[Node]int8) {
	for k := 0; k < 4; k++ {
		newI, newJ := i+directionI[k], j+directionJ[k]
		if _, ok := v[Node{newI, newJ}]; !ok && newI >= 0 && newI < rows && newJ >= 0 && newJ < cols && grid[i][j] == '1' {
			v[Node{newI, newJ}] = 0
			DFS(grid, newI, newJ, rows, cols, v)
		}
	}
}

func main() {
	num := numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}})
	fmt.Println(num)
}
