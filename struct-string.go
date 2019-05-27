package main

import (
	"fmt"
	//"io/ioutil"
	//"os"
)

type path [10]string

type pather struct {
	p1 *path
	p2 *path
}

func main() {
	test1 := path{"how are you"}
	test2 := path{"you"}
	all := pather{
		p1: &test1,
		p2: &test2,
	}
	fmt.Println(all.p2)
}
