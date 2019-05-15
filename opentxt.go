package main

import (
	//"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("./text_collection.txt")
	var readfile io.Reader = inputFile
	p := make([]byte, 10)
	for {
		_, err := readfile.Read(p)
		if err == io.EOF {
			return
		}
		fmt.Print(string(p))
	}
}
