/**
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。



例如:
给定二叉树: [3,9,20,null,null,15,7],

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


提示：

节点总数 <= 1000

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
	res := make([][]int, 0)
	s := make([]*TreeNode, 0)
	s = append(s, root)
	for len(s) > 0 {
		rowData := make([]int, 0)
		Len := len(s)
		for i := 0; i < Len; i++ {
			value := s[i]
			rowData = append(rowData, value.Val)
			if value.Left != nil {
				s = append(s, value.Left)
			}
			if value.Right != nil {
				s = append(s, value.Right)
			}
		}
		s = s[Len:]
		res = append(res, rowData)
	}
	return res
}

func main() {}
