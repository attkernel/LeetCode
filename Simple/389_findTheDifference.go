/*
给定两个字符串 s 和 t，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。



示例:

输入：
s = "abcd"
t = "abcde"

输出：
e

解释：
'e' 是那个被添加的字母。

*/
package main

func findTheDifference(s string, t string) byte {
	sums, sumt := 0, 0
	for _, v := range s {
		sums += int(v)
	}
	for _, v := range t {
		sumt += int(v)
	}
	return byte(sumt - sums)
}

func main() {}
