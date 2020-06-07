/*
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。



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
  [20,9],
  [15,7]
]


提示：

节点总数 <= 1000

*/

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
	res, s, row := make([][]int, 0), make([]*TreeNode, 0), 1
	s = append(s, root)
	for len(s) > 0 {
		l := len(s)
		rowData := make([]int, 0)
		for i := 0; i < l; i++ {
			node := s[i]
			if row%2 == 1 {
				rowData = append(rowData, node.Val)
			} else {
				rowData = append([]int{node.Val}, rowData[:]...)
			}
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
		}
		res = append(res, rowData)
		s = s[l:]
		row++
	}
	return res
}

func main() {}
