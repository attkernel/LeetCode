/*
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

     3
    / \
   4   5
  / \
 1   2
给定的树 B：

   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：

输入：A = [1,2,3], B = [3,1]
输出：false
示例 2：

输入：A = [3,4,5,1,2], B = [4,1]
输出：true
限制：

0 <= 节点个数 <= 10000

*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	//isSubStructure(A.Left, B) || isSubStructure(A.Right, B)这种递归，相当于把所有的树都取到了
	//所以在所有的情况下，只要dfs能返回一个true,就能让所有的false变成true。因为使用的是并集
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B) || dfs(A, B)
}
func dfs(A *TreeNode, B *TreeNode) bool {
	//B已经找完
	if B == nil {
		return true
	}
	//A如果找完，说明没有叶子节点
	if A == nil {
		return false
	}
	return dfs(A.Left, B.Left) && dfs(A.Right, B.Right) && A.Val == B.Val
}

func main() {}
