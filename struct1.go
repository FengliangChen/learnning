package main

import "fmt"

type student struct {
	name string
	age  int
	sex  string
}

type class struct {
	monitor   *student
	committee *student
}

func main() {
	myclass := &class{
		monitor: &student{
			name: "simon",
			age:  18,
		},
		committee: &student{
			name: "mon",
			age:  10,
		},
	}
	fmt.Println(myclass.monitor.age)
}
