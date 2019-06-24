package main

import (
	"fmt"
	//"log"
	//"net/http"
	"time"
)

var meme chan int = make(chan int)

func run(n int) {
	time.Sleep(1e9)
	fmt.Println(n, "---------tested")
	meme <- n
}

func re(n int) {
	for i := 0; i < n; i++ {
		<-meme
	}
}

func main() {
	go run(1)
	go run(2)
	go run(3)
	go run(4)
	go run(5)
	go run(6)
	re(1)
}
