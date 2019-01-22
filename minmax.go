package main

import (
	"fmt"
)

func main() {
	var min, max int
	min, max = minmax(10, 100)
	fmt.Printf("the min is %d, the max is %d\n", min, max)
}

func minmax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return
}
