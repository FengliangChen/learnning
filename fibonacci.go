package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	result := 0
	for i := 0; i <= 45; i++ {
		result = fibonacci(i)
		fmt.Printf("bibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("takes %s to accomplish\n", duration)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

/*33 seconds in my old computer*/