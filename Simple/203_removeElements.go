/**
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5

**/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	tmp := &ListNode{}
	tmp.Next = head
	res := tmp
	for head != nil {
		if head.Val == val {
			tmp.Next = head.Next
			head = head.Next
			continue
		}
		head = head.Next
		tmp = tmp.Next
	}
	return res.Next
}

func main() {}
