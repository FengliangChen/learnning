package main

import (
	"fmt"
	//	"net/http"
	"io/ioutil"
	"regexp"
)

func main() {
	path := "/Volumes/datavolumn_bmkserver_Pub/201906/0605"
	pat := "[BCU][0-9]{4}[0-9A-Z]{2}_[A-Z]{3}"
	re := regexp.MustCompile(pat)
	fileInfo, fileErro := ioutil.ReadDir(path)
	if fileErro != nil {
		fmt.Println("fileErro")
	}
	for _, file := range fileInfo {
		if file.IsDir() {
			if re.Match([]byte(file.Name())) {
				fmt.Println(file.Name())
			}
		}
	}
}
