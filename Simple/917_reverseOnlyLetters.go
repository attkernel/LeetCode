/*
给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。



示例 1：

输入："ab-cd"
输出："dc-ba"
示例 2：

输入："a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：

输入："Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"


提示：

S.length <= 100
33 <= S[i].ASCIIcode <= 122
S 中不包含 \ or "

*/
package main

func reverseOnlyLetters(s string) string {

	i := 0
	j := len(s) - 1
	S := []rune(s)
	for {
		if i > j {
			break
		}

		if S[i] >= 'a' && S[i] <= 'z' || S[i] >= 'A' && S[i] <= 'Z' {
			if S[j] >= 'a' && S[j] <= 'z' || S[j] >= 'A' && S[j] <= 'Z' {
				S[i], S[j] = S[j], S[i]
				i++
				j--
			} else {
				j--
			}
		} else {
			i++
		}
	}
	return string(S)
}

func main() {}
