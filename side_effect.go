/*
6.2.4 改变外部变量（outside variable）
*/

package main

import (
	"fmt"
)

func main() {
	n := 0
	reply := &n
	Multiply(10, 10, reply)
	fmt.Println(n)
}

func Multiply(a, b int, reply *int) {
	*reply = a + b
}
