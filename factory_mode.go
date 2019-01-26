/*6.10 工厂函数*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	addBMP := MakeAddSuffix(".bmp")
	addJPG := MakeAddSuffix(".jpg")
	file1 := addBMP("123456789")
	file2 := addJPG("123456789")
	fmt.Println(file1)
	fmt.Println(file2)
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
