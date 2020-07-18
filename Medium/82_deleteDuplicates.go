/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3

*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	m := make(map[int]int, 0)
	newHead := new(ListNode)
	newHead.Next = head
	p := newHead
	tmp := head
	for tmp != nil {
		m[tmp.Val]++
		tmp = tmp.Next
	}
	for head != nil {
		if v, ok := m[head.Val]; ok && v > 1 {
			p.Next = head.Next
			head = head.Next
			continue
		}
		head = head.Next
		p = p.Next
	}
	return newHead.Next
}

func main() {}
