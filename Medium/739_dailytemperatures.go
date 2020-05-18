/**
根据每日 气温 列表，请重新生成一个列表，对应位置的输出是需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

**/
package main

import (
	"fmt"
)

func dailyTemperatures(T []int) []int {
	if len(T) == 0 {
		return nil
	}
	var result []int
	for i := 0; i < len(T); i++ {
		temp := T[i]
		steps := 0
		isFind := false
		for j := i + 1; j < len(T); j++ {
			steps++
			if T[j] > temp {
				result = append(result, steps)
				isFind = true
				break
			}
		}
		if !isFind {
			result = append(result, 0)
		}
	}
	return result
}
func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}
