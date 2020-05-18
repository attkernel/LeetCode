package main

import (
	"fmt"
)

type RecentCounter struct {
	S []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		S: make([]int, 0, 10000),
	}
}

/*func (this *RecentCounter) Ping(t int) int {
	this.S = append(this.S, t)
	for k, _ := range this.S {
		if this.S[k] >= (t-3000) && this.S[k] <= t {
			return len(this.S) - k
		}
	}
	return 0
}*/

/*func (this *RecentCounter) Ping(t int) int {
	this.S = append(this.S, t)
	for this.S[0] < t-3000 {
		this.S = this.S[1:]
	}
	return len(this.S)
}*/
func (this *RecentCounter) Ping(t int) int {
	this.S = append(this.S, t)
	i, j := 0, len(this.S)
	for i < j {
		mid := i + (j-i)/2
		if this.S[mid] >= t-3000 {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return len(this.S) - i
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))
	fmt.Println(obj.Ping(100))
	fmt.Println(obj.Ping(3001))
	fmt.Println(obj.Ping(3002))
}
