/*
给定一个包含了一些 0 和 1 的非空二维数组 grid 。

一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)



示例 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。

示例 2:

[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。



注意: 给定的矩阵grid 的长度和宽度都不超过 50。

*/
package main

var (
	directionI = []int{0, 0, -1, 1}
	directionJ = []int{1, -1, 0, 0}
)

func maxAreaOfIsland(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	max := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				nums := BFS(grid, row, col, i, j)
				if nums > max {
					max = nums
				}
			}
		}
	}
	return max
}
func BFS(grid [][]int, row, col, i, j int) int {
	nums := 0
	queue := make([]int, 0)
	queue = append(queue, i, j)
	grid[i][j] = 0
	nums++
	for len(queue) > 0 {
		currentI, currentJ := queue[0], queue[1]
		queue = queue[2:]
		for k := 0; k < 4; k++ {
			newI, newJ := currentI+directionI[k], currentJ+directionJ[k]
			if newI >= 0 && newI < row && newJ >= 0 && newJ < col && grid[newI][newJ] == 1 {
				grid[newI][newJ] = 0
				nums++
				queue = append(queue, newI, newJ)
			}
		}
	}
	return nums
}

func main() {}
