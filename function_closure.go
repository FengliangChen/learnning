/*闭包,变量 x 的值是被保留*/

package main

import "fmt"

func main() {
	x := add()
	fmt.Print(x(1), "-")
	fmt.Print(x(1), "-")
	fmt.Print(x(1))
}

func add() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
