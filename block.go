package main

import (
    "fmt"
    // "time"
)

func f1(in chan int) {
	// time.Sleep(1e9)
    fmt.Println(<-in)
}

func main() {
    out := make(chan int)
    out <- 2 // deadlock due to there is no channel to receive this value at this moment. Switch then OK.
    go f1(out)
}