package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir("/Users/imac-6/Desktop/All/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
		k := f.Name()
		fmt.Println(string(k[0]))
	}
}
