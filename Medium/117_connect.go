/**
给定一个二叉树

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。



进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。


示例：



输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。


提示：

树中的节点数小于 6000
-100 <= node.val <= 100

**/
package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	S := make([]*Node, 0)
	S = append(S, root)
	for len(S) > 0 {
		Len := len(S)
		for i := 0; i < Len; i++ {
			if S[i].Left != nil {
				S = append(S, S[i].Left)
			}
			if S[i].Right != nil {
				S = append(S, S[i].Right)
			}
			if i == Len-1 {
				S[i].Next = nil
				continue
			}
			S[i].Next = S[i+1]
		}
		S = S[Len:]
	}
	return root
}

func main() {}
