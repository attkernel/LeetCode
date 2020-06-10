/*
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。



示例:

输入: [1,2,3,4,5]
输出: [120,60,40,30,24]


提示：

所有元素乘积之和不会溢出 32 位整数
a.length <= 100000

*/

package main

func constructArr(a []int) []int {
	res, tmp := make([]int, len(a)), 1
	for idx := range a {
		res[idx] = 1
	}
	for i := 1; i < len(a); i++ {
		res[i] = res[i-1] * a[i-1]
	}
	for i := len(a) - 2; i >= 0; i-- {
		res[i] = res[i] * a[i+1] * tmp
		tmp *= a[i+1]
	}
	return res
}

func main() {}
