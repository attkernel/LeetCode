/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。



示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]


限制：

1 <= target <= 10^5

*/
package main

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	pFast, pSlow := 2, 1
	for pSlow < pFast {
		sum := (pFast + pSlow) * (pFast - pSlow + 1) / 2
		if sum == target {
			rowData := make([]int, 0)
			for i := pSlow; i <= pFast; i++ {
				rowData = append(rowData, i)
			}
			res = append(res, rowData)
			pFast++
		} else if sum > target {
			pSlow++
		} else {
			pFast++
		}
	}
	return res
}

func main() {}
