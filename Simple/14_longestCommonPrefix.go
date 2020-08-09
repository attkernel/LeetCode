/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

*/
package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	lenMin := len(strs[0])
	idxMin := 0
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < lenMin {
			idxMin = i
		}
	}
	for k, v := range strs[idxMin] {
		for _, value := range strs {
			if value[k] != byte(v) {
				return string(strs[idxMin][:k])
			}
		}
	}
	return strs[idxMin]
}

func main() {}
