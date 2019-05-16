package main

import (
	//"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("how are you!\n")
	proverbs.WriteString("this is the test \n")

	file, err := os.Create("./yest.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	if _, err := io.Copy(file, proverbs); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("file created")
}
