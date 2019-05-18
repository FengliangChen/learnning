package main

import (
	"fmt"
)

func PrintMe(num ...int) {
	if len(num) == 1 {
		fmt.Println(num[0])
	} else {
		for _, n := range num {
			fmt.Println(n)
		}
	}
}

func main() {
	PrintMe(1)
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	PrintMe(s...)
}

/* 传递多个参数可以选择：
struct 传递
空接 + for-range, switch 断言
*/

