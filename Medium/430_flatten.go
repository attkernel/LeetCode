package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	current := root
	for current != nil {
		if current.Child != nil {
			tmp1 := current
			tmp2 := current.Next
			current.Child.Prev = current
			current.Next = current.Child
			current.Child = nil
			for tmp1.Next != nil {
				tmp1 = tmp1.Next
			}
			tmp1.Next = tmp2
			if tmp2 != nil {
				tmp2.Prev = tmp1
			}
		}
		current = current.Next
	}
	return root
}

func main() {}
