/*
返回字符串 text 中按字典序排列最小的子序列，该子序列包含 text 中所有不同字符一次。



示例 1：

输入："cdadabcc"
输出："adbc"
示例 2：

输入："abcd"
输出："abcd"
示例 3：

输入："ecbacba"
输出："eacb"
示例 4：

输入："leetcode"
输出："letcod"


提示：

1 <= text.length <= 1000
text 由小写英文字母组成

*/
package main

func smallestSubsequence(s string) string {
	var result string
	lastIndex := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		lastIndex[s[i]] = i
	}
	used := make(map[uint8]bool)
	for i := 0; i < len(s); i++ {
		if used[s[i]] {
			continue
		}
		for len(result) != 0 && result[len(result)-1] > s[i] && lastIndex[result[len(result)-1]] > i {
			used[result[len(result)-1]] = false
			result = result[:len(result)-1]
		}
		used[s[i]] = true
		result += string(s[i])
	}
	return result
}

func main() {}
