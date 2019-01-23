/*6.7将函数作为参数*/

package main

import "fmt"

func main() {
	callback(2, Add)

}

func Add(a, b int) {
	fmt.Printf("the sum of %d and %d are %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 999)
}
