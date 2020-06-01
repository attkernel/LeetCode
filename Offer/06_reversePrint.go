/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]


限制：

0 <= 链表长度 <= 10000

**/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	Res := make([]int, 0)
	for head != nil {
		Res = append(Res, head.Val)
		head = head.Next
	}
	for i, j := 0, len(Res)-1; i <= j; i, j = i+1, j-1 {
		Res[i], Res[j] = Res[j], Res[i]
	}
	return Res
}

func main() {}
