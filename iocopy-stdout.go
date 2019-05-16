package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	s := []byte("hello how are you \n")
	b := new(bytes.Buffer)
	b.Write(s)
	io.Copy(os.Stdout, b)
}
