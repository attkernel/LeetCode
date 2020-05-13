/**
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。


示例：

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3


提示：

所有val值都在 [1, 1000] 之内。
操作次数将在  [1, 1000] 之内。
请不要使用内置的 LinkedList 库。

**/
package main

import (
	"fmt"
)

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		Val:  0,
		Next: nil,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	Pointor := this.Next
	for Pointor != nil {
		if index == 0 {
			return Pointor.Val
		}
		Pointor = Pointor.Next
		index--
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	Node := &MyLinkedList{
		Val: val,
	}
	Node.Next = this.Next
	this.Next = Node
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	Pointor := this.Next
	Node := &MyLinkedList{
		Val:  val,
		Next: nil,
	}
	for Pointor.Next != nil {
		Pointor = Pointor.Next
	}
	Pointor.Next = Node
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {

	Pointor := this
	Node := &MyLinkedList{
		Val: val,
	}
	for Pointor != nil {
		if index == 0 {
			Node.Next = Pointor.Next
			Pointor.Next = Node
			return
		}
		Pointor = Pointor.Next
		index--
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	Pointor := this
	for Pointor.Next != nil {
		if index == 0 {
			Pointor.Next = Pointor.Next.Next
			return
		}
		index--
		Pointor = Pointor.Next
	}
}

func main() {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2)    //链表变为1-> 2-> 3
	fmt.Println(linkedList.Get(1)) //返回2
	fmt.Println(linkedList.Get(0))
	fmt.Println(linkedList.Get(2))
	linkedList.DeleteAtIndex(1)    //现在链表是1-> 3
	fmt.Println(linkedList.Get(1)) //返回3
}
