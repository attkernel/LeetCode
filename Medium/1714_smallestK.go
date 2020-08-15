/*
设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。

示例：

输入： arr = [1,3,5,7,2,4,6,8], k = 4
输出： [1,2,3,4]
提示：

0 <= len(arr) <= 100000
0 <= k <= min(100000, len(arr))

*/
package main

func smallestK(arr []int, k int) []int {
	mh := NewMinHeap(arr)
	arr1 := []int{}

	for k > 0 {
		arr1 = append(arr1, mh.Pop())
		k--
	}
	return arr1
}

//小顶堆的实现-------------------
type minHeap struct {
	h []int
}

func NewMinHeap(arr []int) *minHeap {
	mh := &minHeap{h: arr}
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		mh.downAdjust(i)
	}
	return mh
}
func (this *minHeap) downAdjust(parentIndex int) {
	tmp := this.h[parentIndex]
	childIndex := 2*parentIndex + 1
	Hlen := len(this.h)

	for childIndex < Hlen {
		if childIndex+1 < Hlen && this.h[childIndex+1] < this.h[childIndex] {
			childIndex++
		}
		if tmp < this.h[childIndex] {
			break
		}
		this.h[parentIndex] = this.h[childIndex]
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}
	this.h[parentIndex] = tmp
}
func (this *minHeap) upAdjust() {
	childIndex := len(this.h) - 1
	parentIndex := (childIndex - 1) / 2
	tmp := this.h[childIndex]

	for childIndex > 0 && tmp < this.h[parentIndex] {
		this.h[childIndex] = this.h[parentIndex]
		childIndex = parentIndex
		parentIndex = (childIndex - 1) / 2
	}
	this.h[childIndex] = tmp
}
func (this *minHeap) Insert(a int) {
	this.h = append(this.h, a)
	this.upAdjust()
}
func (this *minHeap) Pop() int {
	first := this.h[0]
	this.h[0] = this.h[len(this.h)-1]
	this.h = this.h[:len(this.h)-1]
	this.downAdjust(0)
	return first
}

func main() {}
