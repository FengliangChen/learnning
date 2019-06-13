package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("in main()")
	go longWait()
	go shortWait()
	fmt.Println("start to sleep")
	time.Sleep(10 * 1e9)
	fmt.Println("end of main()")
}

func longWait() {
	fmt.Println("start the long wait")
	time.Sleep(5 * 1e9)
	fmt.Println("End of the long wait")
}
func shortWait() {
	fmt.Println("start the short wait")
	time.Sleep(2 * 1e9)
	fmt.Println("End of the short wait")
}
