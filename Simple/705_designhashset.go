package main

/**
不使用任何内建的哈希表库设计一个哈希集合

具体地说，你的设计应该包含以下的功能

add(value)：向哈希集合中插入一个值。
contains(value) ：返回哈希集合中是否存在这个值。
remove(value)：将给定值从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。

示例:

MyHashSet hashSet = new MyHashSet();
hashSet.add(1);
hashSet.add(2);
hashSet.contains(1);    // 返回 true
hashSet.contains(3);    // 返回 false (未找到)
hashSet.add(2);
hashSet.contains(2);    // 返回 true
hashSet.remove(2);
hashSet.contains(2);    // 返回  false (已经被删除)

注意：

所有的值都在 [0, 1000000]的范围内。
操作的总数目在[1, 10000]范围内。
不要使用内建的哈希集合库。

**/
type Bucket struct {
	B []int
}

type MyHashSet struct {
	Set []*Bucket
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	set := MyHashSet{
		Set: make([]*Bucket, 769),
	}
	for k, _ := range set.Set {
		set.Set[k] = new(Bucket)
	}
	return set
}

func (b *Bucket) Contains(key int) bool {
	for k, _ := range b.B {
		if b.B[k] == key {
			return true
		}
	}
	return false
}

func (b *Bucket) Add(key int) {
	if b.Contains(key) {
		return
	}
	b.B = append(b.B, key)
}

func (b *Bucket) Remove(key int) {
	for k, _ := range b.B {
		if b.B[k] == key {
			b.B = append(b.B[:k], b.B[k+1:]...)
			return
		}
	}
}

func (this *MyHashSet) calKey(key int) int {
	return key % 769
}

func (this *MyHashSet) Add(key int) {
	this.Set[this.calKey(key)].Add(key)
}

func (this *MyHashSet) Remove(key int) {
	this.Set[this.calKey(key)].Remove(key)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.Set[this.calKey(key)].Contains(key)
}

func main() {}
