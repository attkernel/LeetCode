/**
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。



示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
     2->1->3->4->5->NULL

限制：

0 <= 节点个数 <= 5000

**/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := head
	for head.Next != nil {
		tmp := head.Next
		head.Next = tmp.Next
		tmp.Next = newHead
		newHead = tmp
	}
	return newHead
}

func main() {}
