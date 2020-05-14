/**
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

示例 1:

输入: "III"
输出: 3
示例 2:

输入: "IV"
输出: 4
示例 3:

输入: "IX"
输出: 9
示例 4:

输入: "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.

**/

package main

import (
	"fmt"
)

func romanToInt(s string) int {
	M := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	Bytes := []byte(s)
	Bound := len(s) - 1
	Res := 0
	for idx, _ := range Bytes {
		if Bytes[idx] == 'I' && idx+1 <= Bound && Bytes[idx+1] == 'V' {
			Res += 4
			Res -= 5
			continue
		}
		if Bytes[idx] == 'I' && idx+1 <= Bound && Bytes[idx+1] == 'X' {
			Res += 9
			Res -= 10
			continue
		}
		if Bytes[idx] == 'X' && idx+1 <= Bound && Bytes[idx+1] == 'L' {
			Res += 40
			Res -= 50
			continue
		}
		if Bytes[idx] == 'X' && idx+1 <= Bound && Bytes[idx+1] == 'C' {
			Res += 90
			Res -= 100
			continue
		}
		if Bytes[idx] == 'C' && idx+1 <= Bound && Bytes[idx+1] == 'D' {
			Res += 400
			Res -= 500
			continue
		}
		if Bytes[idx] == 'C' && idx+1 <= Bound && Bytes[idx+1] == 'M' {
			Res += 900
			Res -= 1000
			continue
		}
		Res += M[Bytes[idx]]
	}
	return Res
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
