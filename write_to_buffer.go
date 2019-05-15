package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("./text_collection.txt")
	defer inputFile.Close()
	var writer bytes.Buffer
	p := make([]byte, 8)
	for {
		_, err := inputFile.Read(p)
		if err == io.EOF {
			break
		}
		w, e := writer.Write(p)
		if e != nil {
			fmt.Println(e)
		}
		if w != len(p) {
			fmt.Println("failed to write")
			os.Exit(1)
		}
	}
	fmt.Println(writer.String())
}
