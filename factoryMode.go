package main

import (
	"fmt"
	//"os"
)

type options struct {
	tag     int
	class   string
	content string
}

type option func(*options)

func Class(class string) option {
	return func(arg *options) {
		arg.class = class
	}
}

func main() {
	classA := Class("DDD")
	a := &options{}
	classA(a)
	fmt.Println(a)

}
