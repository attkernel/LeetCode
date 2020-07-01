/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(getHeight(root.Left)-getHeight(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(getHeight(root.Left), getHeight(root.Right))
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {}
