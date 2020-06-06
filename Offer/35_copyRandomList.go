/*
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。



示例 1：



输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
示例 2：



输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]
示例 3：



输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]
示例 4：

输入：head = []
输出：[]
解释：给定的链表为空（空指针），因此返回 null。


提示：

-10000 <= Node.val <= 10000
Node.random 为空（null）或指向链表中的节点。
节点数目不超过 1000 。

*/
package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	M := make(map[*Node]*Node)
	M[head] = &Node{
		Val: head.Val,
	}
	res := head
	for head.Next != nil {
		if _, ok := M[head.Next]; ok {
			M[head].Next = M[head.Next]
		} else {
			node := &Node{
				Val: head.Next.Val,
			}
			M[head.Next] = node
			M[head].Next = M[head.Next]
		}
		if _, ok := M[head.Random]; ok {
			M[head].Random = M[head.Random]
		} else {
			if head.Random != nil {
				node := &Node{
					Val: head.Random.Val,
				}
				M[head.Random] = node
				M[head].Random = M[head.Random]
			}
		}
		head = head.Next
	}
	head.Next = nil
	M[head].Random = M[head.Random]
	return M[res]
}

func main() {}
