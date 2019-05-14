package main

import (
	"fmt"
	"os"
)

func main() {
	fileinfo, _ := os.Stat("zip.go")
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.IsDir())
	fmt.Println(fileinfo.Size())
}
