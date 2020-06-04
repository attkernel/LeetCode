/**
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：

给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。



提示：

节点总数 <= 10000

**/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return recurs(root, 1)
}

func recurs(root *TreeNode, depth int) int {
	if root.Left == nil && root.Right == nil {
		return depth
	}
	left, right := 0, 0
	if root.Left != nil {
		left = recurs(root.Left, depth+1)
	}
	if root.Right != nil {
		right = recurs(root.Right, depth+1)
	}
	if left > right {
		return left
	}
	return right
}

func main() {}
