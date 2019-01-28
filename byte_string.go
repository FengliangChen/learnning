package main

import (
	"fmt"
)

func main() {
	var b []byte
	var s string
	s = "hello"
	c := append(b, s...)
	fmt.Println(c)
	d := string(c)
	fmt.Println(d)
}
