package main

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	wordNums := make([]int, len(words))
	for i := range words {
		wordNums[i] = f(words[i])
	}
	sort.Ints(wordNums)
	l := len(words)
	nums := make([]int, len(queries))
	var index, freq int
	for i := range queries {
		freq = f(queries[i])
		index = sort.SearchInts(wordNums, freq+1)
		nums[i] = l - index
	}
	return nums
}

func f(s string) int {
	minc := 'z' + 1
	cnt := 0
	for _, c := range s {
		if minc > c {
			minc = c
			cnt = 1
		} else if minc == c {
			cnt++
		}
	}
	return cnt
}

func main() {}
