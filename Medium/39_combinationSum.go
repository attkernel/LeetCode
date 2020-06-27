/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/

package main

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	seq := make([]int, 0)
	var combination func(int, int)

	combination = func(t, prev int) {
		for _, candidate := range candidates {
			if candidate > t || candidate < prev {
				continue
			} else if candidate == t {
				tmp := make([]int, len(seq))
				copy(tmp, seq)
				tmp = append(tmp, candidate)
				res = append(res, tmp)
			} else {
				seq = append(seq, candidate)
				combination(t-candidate, candidate)
				seq = seq[:len(seq)-1]
			}
		}
	}
	combination(target, -1)

	return res
}

func main() {}
