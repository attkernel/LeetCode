/*
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。



进阶：

如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。



示例：

输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7

*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2 []int
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	var head *ListNode = nil
	sum, index, len1, len2 := 0, 1, len(stack1), len(stack2)
	for index <= len1 || index <= len2 || sum > 0 {
		if index <= len1 {
			sum += stack1[len1-index]
		}
		if index <= len2 {
			sum += stack2[len2-index]
		}

		temp := ListNode{
			Val:  sum % 10,
			Next: head,
		}
		head = &temp
		sum /= 10
		index++
	}
	return head
}

func main() {}
