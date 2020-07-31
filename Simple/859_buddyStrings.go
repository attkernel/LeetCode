/*
给定两个由小写字母构成的字符串 A 和 B ，只要我们可以通过交换 A 中的两个字母得到与 B 相等的结果，就返回 true ；否则返回 false 。



示例 1：

输入： A = "ab", B = "ba"
输出： true
示例 2：

输入： A = "ab", B = "ab"
输出： false
示例 3:

输入： A = "aa", B = "aa"
输出： true
示例 4：

输入： A = "aaaaaaabc", B = "aaaaaaacb"
输出： true
示例 5：

输入： A = "", B = "aa"
输出： false


提示：

0 <= A.length <= 20000
0 <= B.length <= 20000
A 和 B 仅由小写字母构成。

*/
package main

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) || len(A) < 2 || len(B) < 2 {
		return false
	}
	tmp := []byte{}
	letter := [26]int{}
	for i := 0; i < len(A); i++ {
		letter[A[i]-'a']++
		if A[i] != B[i] {
			tmp = append(tmp, A[i])
			tmp = append(tmp, B[i])
		}
	}
	for j := 0; j < len(letter); j++ {
		if letter[j] >= 2 && A == B {
			return true
		}
	}
	if len(tmp) != 4 {
		return false
	}
	if tmp[0] != tmp[3] || tmp[1] != tmp[2] {
		return false
	}
	return true
}

func main() {}
