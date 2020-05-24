package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	M := make(map[*Node]*Node)
	res := head
	for head != nil {
		if _, ok := M[head]; !ok {
			M[head] = &Node{
				Val: head.Val,
			}
		}
		next := head.Next
		if next == nil {
			M[head].Next = nil
		} else {
			if _, ok := M[next]; !ok {
				M[next] = &Node{
					Val: next.Val,
				}
			}
		}
		M[head].Next = M[next]

		random := head.Random
		if random == nil {
			M[head].Random = nil
		} else {
			if _, ok := M[random]; !ok {
				M[random] = &Node{
					Val: random.Val,
				}
			}
		}
		M[head].Random = M[random]
		head = head.Next
	}
	return M[res]
}

func main() {}
