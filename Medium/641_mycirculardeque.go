package main

import "fmt"

type MyCircularDeque struct {
	S []int
	F int
	E int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		S: make([]int, k+1, k+1),
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.F = (this.F + len(this.S) - 1) % len(this.S)
	this.S[this.F] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.S[this.E] = value
	this.E = (this.E + 1) % len(this.S)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.F = (this.F + 1) % len(this.S)
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.E = (this.E + len(this.S) - 1) % len(this.S)
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.S[this.F]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.S[(this.E+len(this.S)-1)%len(this.S)]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.F == this.E
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return (this.E+1)%len(this.S) == this.F
}

func main() {
	circularDeque := Constructor(3)
	fmt.Println(circularDeque.InsertFront(9))
	fmt.Println(circularDeque.GetRear())
	fmt.Println(circularDeque.InsertFront(9))
	fmt.Println(circularDeque.GetRear())
	fmt.Println(circularDeque.InsertLast(5))
	fmt.Println(circularDeque.GetFront())
	fmt.Println(circularDeque.GetRear())
	fmt.Println(circularDeque.GetFront())
	fmt.Println(circularDeque.InsertLast(8))
	fmt.Println(circularDeque.DeleteLast())
	fmt.Println(circularDeque.GetFront())
}

//["MyCircularDeque","insertFront","getRear","insertFront","getRear","insertLast","getFront","getRear","getFront","insertLast","deleteLast","getFront"]
//[[3],[9],[],[9],[],[5],[],[],[],[8],[],[]]
