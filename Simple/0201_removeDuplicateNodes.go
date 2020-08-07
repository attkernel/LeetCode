/*
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:

 输入：[1, 2, 3, 3, 2, 1]
 输出：[1, 2, 3]
示例2:

 输入：[1, 1, 1, 1, 2]
 输出：[1, 2]
提示：

链表长度在[0, 20000]范围内。
链表元素在[0, 20000]范围内。
进阶：

如果不得使用临时缓冲区，该怎么解决？

*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	m := make(map[int]struct{}, 0)
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	if slow.Val == fast.Val && fast.Next == nil {
		slow.Next = nil
		return slow
	}
	m[slow.Val] = struct{}{}
	for fast != nil {
		if _, ok := m[fast.Val]; ok {
			fast = fast.Next
			slow.Next = fast
		} else {
			m[fast.Val] = struct{}{}
			fast = fast.Next
			slow = slow.Next
		}
	}
	return head
}

func main() {}
