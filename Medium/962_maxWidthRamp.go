/*
给定一个整数数组 A，坡是元组 (i, j)，其中  i < j 且 A[i] <= A[j]。这样的坡的宽度为 j - i。

找出 A 中的坡的最大宽度，如果不存在，返回 0 。



示例 1：

输入：[6,0,8,2,1,5]
输出：4
解释：
最大宽度的坡为 (i, j) = (1, 5): A[1] = 0 且 A[5] = 5.
示例 2：

输入：[9,8,1,0,1,9,4,0,4,1]
输出：7
解释：
最大宽度的坡为 (i, j) = (2, 9): A[2] = 1 且 A[9] = 1.


提示：

2 <= A.length <= 50000
0 <= A[i] <= 50000

*/
package main

func maxWidthRamp(A []int) int {
	stack := []int{}
	n := len(A)

	for i := 0; i < n; i++ {
		if len(stack) == 0 || A[stack[len(stack)-1]] > A[i] {
			stack = append(stack, i)
		}
	}

	res := 0
	i := n - 1
	for i > res {
		for len(stack) > 0 && A[stack[len(stack)-1]] <= A[i] {
			if i-stack[len(stack)-1] > res {
				res = i - stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
		i--
	}

	return res
}

func main() {}
