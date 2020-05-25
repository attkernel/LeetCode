/**
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

**/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	S := make([]*TreeNode, 0)
	Res := make([]int, 0)
	for len(S) > 0 || root != nil {
		for root != nil {
			S = append(S, root)
			root = root.Left
		}
		node := S[len(S)-1]
		S = S[:len(S)-1]
		Res = append(Res, node.Val)
		root = node.Right
	}
	return Res
}

func main() {}
