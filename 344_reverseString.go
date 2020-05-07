/**
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。



示例 1：

输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]
示例 2：

输入：["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]

**/
package main

import (
	"fmt"
	"unsafe"
)

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	/*var address = (**[5]int8)(unsafe.Pointer(&s))
	var body = **address
	var tmp int8
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		tmp = body[j]
		body[j] = body[i]
		body[i] = tmp
		//s[i], s[j] = s[j], s[i]
	}*/
}

func main() {
	//Slice1 := []byte{'h', 'e', 'l', 'l', 'o'}
	//Slice2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	//reverseString(Slice1)
	//reverseString(Slice2)
	//fmt.Printf("%s\n", Slice1)
	//fmt.Printf("%s\n", Slice2)
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	up := uintptr(unsafe.Pointer(&s[0]))
	/*v1 := (*int8)(unsafe.Pointer(up))
	up += unsafe.Sizeof(int8(0))
	v2 := (*int8)(unsafe.Pointer(up))
	fmt.Printf("%c\n", *v1)
	fmt.Printf("%c\n", *v2)*/
	for i := 1; i < len(s); i++ {
		up += unsafe.Sizeof(int8(0))
		fmt.Printf("%c", *(*int8)(unsafe.Pointer(up)))
	}
}
