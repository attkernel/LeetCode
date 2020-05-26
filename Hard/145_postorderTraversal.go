/**
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

**/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	S := make([]*TreeNode, 0)
	Res := make([]int, 0)
	S = append(S, root)
	for len(S) > 0 {
		node := S[len(S)-1]
		S = S[:len(S)-1]
		Res = append([]int{node.Val}, Res...)
		if node.Left != nil {
			S = append(S, node.Left)
		}
		if node.Right != nil {
			S = append(S, node.Right)
		}
	}
	return Res
}

func main() {}
