/**
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。



示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."


限制：

0 <= s 的长度 <= 10000

**/
package main

import (
	"fmt"
)

func replaceSpace(s string) string {
	Res := make([]byte, 0)
	for idx, _ := range s {
		if s[idx] == ' ' {
			Res = append(Res, []byte("%20")...)
			continue
		}
		Res = append(Res, s[idx])
	}
	return string(Res)
}

func main() {
	fmt.Println(replaceSpace("We are happy."))
}
