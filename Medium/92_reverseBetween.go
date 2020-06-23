/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummnyhead := &ListNode{-1, head}
	i, j := m, n-m
	dummy := dummnyhead
	for i > 1 {
		dummy = dummy.Next
		i--
	}
	cur := dummy.Next.Next
	pre := dummy.Next
	for j > 0 {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		j--
	}
	h := dummy.Next
	dummy.Next = pre
	h.Next = cur
	return dummnyhead.Next
}

func main() {}
