/*
有 n 个城市通过 m 个航班连接。每个航班都从城市 u 开始，以价格 w 抵达 v。

现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到从 src 到 dst 最多经过 k 站中转的最便宜的价格。 如果没有这样的路线，则输出 -1。



示例 1：

输入:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
输出: 200
解释:
城市航班图如下


从城市 0 到城市 2 在 1 站中转以内的最便宜价格是 200，如图中红色所示。
示例 2：

输入:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 0
输出: 500
解释:
城市航班图如下


从城市 0 到城市 2 在 0 站中转以内的最便宜价格是 500，如图中蓝色所示。


提示：

n 范围是 [1, 100]，城市标签从 0 到 n - 1.
航班数量范围是 [0, n * (n - 1) / 2].
每个航班的格式 (src, dst, price).
每个航班的价格范围是 [1, 10000].
k 范围是 [0, n - 1].
航班没有重复，且不存在环路

*/
package main

import (
	"fmt"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	if n < 1 {
		return -1
	}
	matrix := make([][]int, n)
	save := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		save[i] = make([]int, K+1)
	}

	for i := 0; i < len(flights); i++ {
		s, e, v := flights[i][0], flights[i][1], flights[i][2]
		matrix[s][e] = v
	}

	res := math.MaxInt32
	queue := make([]int, 0, 2*n)
	queue = append(queue, src)
	step := 0

	for len(queue) > 0 && step <= K {

		preLen := len(queue)
		if preLen > n {
			fmt.Println(len(queue))
		}
		for j := 0; j < preLen; j++ {
			s := queue[0]
			queue = queue[1:]
			for i := 0; i < n; i++ {
				if matrix[s][i] <= 0 {
					continue
				}
				if step == 0 {
					save[i][step] = 0 + matrix[s][i]
				} else {
					s := save[s][step-1] + matrix[s][i]
					if s < save[i][step] || save[i][step] == 0 {
						save[i][step] = s
					} else {
						continue
					}
				}
				if save[i][step] > res {
					continue
				}

				if i == dst {
					res = save[i][step]
				}
				queue = append(queue, i)
			}
		}
		step++
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func main() {}
