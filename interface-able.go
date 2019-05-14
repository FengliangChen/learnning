package main

import "fmt"

type S struct {
	name string
	age  int
}

func (p *S) printName() {
	fmt.Println(p.name)
}

func (p *S) printAge() {
	fmt.Println(p.age)
}

type P interface {
	printName()
	printAge()
}

func setble(n string, a int) P {
	return &S{name: n, age: a}
}

func main() {
	me := setble("simon", 28)
	me.printName()
	me.printAge()
	kate := setble("katle", 100)
	kate.printName()
	kate.printAge()
}
