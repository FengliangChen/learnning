package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("/Users/imac-6/Desktop", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("1 error")
			return err
		}
		if info.IsDir() {
			fmt.Println("yes Dir", info.Name())
		}
		return nil
	})

}
