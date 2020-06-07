/*
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]


提示：

节点总数 <= 1000

*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	s := make([]*TreeNode, 0)
	s = append(s, root)
	for len(s) > 0 {
		l := len(s)
		for i := 0; i < l; i++ {
			root := s[i]
			res = append(res, root.Val)
			if root.Left != nil {
				s = append(s, root.Left)
			}
			if root.Right != nil {
				s = append(s, root.Right)
			}
		}
		s = s[l:]
	}
	return res
}

func main() {}
