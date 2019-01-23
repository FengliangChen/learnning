/*6.9 应用闭包：将函数作为返回值*/

package main

import "fmt"

func main() {
	p2 := Add2()
	fmt.Printf("the number is 2 and the after is %v \n", p2(2))
	TwoAdder := Adder(2)
	fmt.Printf("the Adder result is %v", TwoAdder(9))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
