/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3

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
	res, m := head, make(map[int]struct{}, 0)
	m[head.Val] = struct{}{}
	for head != nil && head.Next != nil {
		if _, ok := m[head.Next.Val]; !ok {
			m[head.Next.Val] = struct{}{}
			head = head.Next
		} else {
			head.Next = head.Next.Next
		}
	}
	return res
}

func main() {}
