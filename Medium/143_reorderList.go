/*
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	pFirst, pPointor := head, head
	for pPointor.Next.Next != nil {
		pPointor = pPointor.Next
	}
	pLast := pPointor.Next
	pPointor.Next = nil
	tmp := pFirst.Next
	pFirst.Next = pLast
	pLast.Next = tmp
	reorderList(tmp)
}

func main() {}
