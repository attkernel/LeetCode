/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。



示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

**/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	Res := make([][]int, 0)
	S := make([]*TreeNode, 0)
	S = append(S, root)
	for len(S) > 0 {
		Len := len(S)
		rowRes := make([]int, 0)
		for i := 0; i < Len; i++ {
			node := S[i]
			rowRes = append(rowRes, node.Val)
			if node.Left != nil {
				S = append(S, node.Left)
			}
			if node.Right != nil {
				S = append(S, node.Right)
			}
		}
		Res = append(Res, rowRes)
		S = S[Len:]
	}
	return Res
}

func main() {

}
