/**
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

**/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	S := make([]int, 0)
	for head != nil {
		S = append(S, head.Val)
		head = head.Next
	}
	for i, j := 0, len(S)-1; i <= j; i, j = i+1, j-1 {
		if S[i] != S[j] {
			return false
		}
	}
	return true
}

func main() {}
