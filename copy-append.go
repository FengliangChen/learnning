package main

import (
	"fmt"
	//"io/ioutil"
	//"os"
)

func main() {
	a := make([]byte, 9)
	b := copy(a[:], "how are you")
	c := make([]byte, 10)
	d := append(c[:], 10, 10)
	fmt.Println("copy is ", a, "len is:", b)
	fmt.Println("append is ", d)
}
