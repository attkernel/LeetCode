/**
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

**/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	S := make([]*TreeNode, 0)
	Res := make([]int, 0)
	if root == nil {
		return nil
	}
	S = append(S, root)
	for len(S) > 0 {
		node := S[len(S)-1]
		S = S[:len(S)-1]
		Res = append(Res, node.Val)
		if node.Right != nil {
			S = append(S, node.Right)
		}
		if node.Left != nil {
			S = append(S, node.Left)
		}
	}
	return Res
}

func main() {}
