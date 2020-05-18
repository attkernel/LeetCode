/**
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例:

s = "3[a]2[bc]", 返回 "aaabcbc".
s = "3[a2[c]]", 返回 "accaccacc".
s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".

**/
package main

import (
	"fmt"
	"strconv"
)

type Mystack struct {
	S []byte
}

func NewStack() Mystack {
	return Mystack{
		S: make([]byte, 0),
	}
}

func (this *Mystack) Empty() bool {
	return len(this.S) == 0
}

func (this *Mystack) Peek() byte {
	if this.Empty() {
		return 0
	}
	return this.S[len(this.S)-1]
}

func (this *Mystack) Pop() byte {
	if this.Empty() {
		return 0
	}
	value := this.S[len(this.S)-1]
	this.S = this.S[:len(this.S)-1]
	return value
}

func (this *Mystack) Push(b ...byte) {
	this.S = append(this.S, b...)
}

func (this *Mystack) GetValue() []byte {
	return this.S
}

func Reverse(b []byte) {
	for left, right := 0, len(b)-1; left < right; left, right = left+1, right-1 {
		b[left], b[right] = b[right], b[left]
	}
}

func IsDigital(b byte) bool {
	return b >= '0' && b <= '9'
}
func decodeString(s string) string {
	var (
		NumSlice = make([]byte, 0)
		ResSlice = make([]byte, 0)
		Stack    = NewStack()
	)
	for k, _ := range s {
		if s[k] == ']' {
			for Stack.Peek() != '[' {
				ResSlice = append(ResSlice, Stack.Pop())
			}
			Reverse(ResSlice)
			Stack.Pop()
			for IsDigital(Stack.Peek()) {
				NumSlice = append(NumSlice, Stack.Pop())
			}
			Reverse(NumSlice)
			num, _ := strconv.Atoi(string(NumSlice))
			for i := 0; i < num; i++ {
				Stack.Push(ResSlice...)
			}
			NumSlice = NumSlice[0:0]
			ResSlice = ResSlice[0:0]
			continue
		}
		Stack.Push(s[k])
	}
	return string(Stack.GetValue())
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("100[leetcode]"))
}
