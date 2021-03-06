/**
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

**/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := postorder[len(postorder)-1]
	for k, _ := range inorder {
		if inorder[k] == root {
			return &TreeNode{
				Val:   root,
				Left:  buildTree(inorder[:k], postorder[:k]),
				Right: buildTree(inorder[k+1:], postorder[k:len(postorder)-1]),
			}
		}
	}
	return nil
}

func main() {}
