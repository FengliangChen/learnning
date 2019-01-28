//结构体工厂

package main

import "fmt"

type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd, name}
}

func main() {
	f := NewFile(10, "123.txt")
	fmt.Println(f.name)
}
