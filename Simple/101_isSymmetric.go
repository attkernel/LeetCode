/**
给定一个二叉树，检查它是否是镜像对称的。



例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3


但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3


进阶：

你可以运用递归和迭代两种方法解决这个问题吗？

**/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recursion(root.Left, root.Right)
}

func recursion(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	return left.Val == right.Val && recursion(left.Left, right.Right) && recursion(left.Right, right.Left)
}*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	S := make([]*TreeNode, 0)
	S = append(S, root.Left)
	S = append(S, root.Right)
	for len(S) > 0 {
		node1 := S[0]
		node2 := S[1]
		S = S[2:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		S = append(S, node1.Left)
		S = append(S, node2.Right)
		S = append(S, node1.Right)
		S = append(S, node2.Left)
	}
	return true
}

func main() {}
