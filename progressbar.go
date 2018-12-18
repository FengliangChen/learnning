package main

import (
	"fmt"
	"strconv"
	"time"
)

func addString(a int) {
	mark := "["
	for b := 0; b <= 100; b++ {
		if b < a {
			mark = mark + "#"
		} else if a == 100 {
			break
		} else {
			mark = mark + "."
		}
	}
	mark = mark + "]" + strconv.Itoa(a) + "%"
	fmt.Printf("%s\r", mark)
	return
}

func main() {
	for b := 0; b <= 100; b++ {
		addString(b)
		time.Sleep(1 * time.Second)
	}
}
