package main

import (
	"fmt"
)

type MyLinkedList struct {
	val  int
	prev *MyLinkedList
	next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		val:  0,
		prev: nil,
		next: nil,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	pointer := this.next
	for pointer != nil {
		if index == 0 {
			return pointer.val
		}
		index--
		pointer = pointer.next
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := &MyLinkedList{
		val: val,
	}
	if this.next == nil {
		this.next = node
		node.prev = this
		return
	}
	node.next = this.next
	node.prev = this
	this.next.prev = node
	this.next = node
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := &MyLinkedList{
		val: val,
	}
	for this.next != nil {
		this = this.next
	}
	this.next = node
	node.prev = this
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	for this.next != nil {
		if index == 0 {
			this.AddAtHead(val)
			return
		}
		index--
		this = this.next
	}
	if this.next == nil && index == 0 {
		this.AddAtTail(val)
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	pointer := this.next
	for pointer != nil {
		if index == 0 {
			if pointer.next == nil {
				pointer.prev.next = pointer.next
				return
			}
			pointer.prev.next = pointer.next
			pointer.next.prev = pointer.prev
			return
		}
		index--
		pointer = pointer.next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func main() {
	Link := Constructor()
	Link.AddAtHead(7)
	Link.AddAtHead(2)
	Link.AddAtHead(1)
	Link.AddAtIndex(3, 0)
	Link.DeleteAtIndex(2)
	Link.AddAtHead(6)
	Link.AddAtTail(4)
	fmt.Println(Link.Get(4))
}
