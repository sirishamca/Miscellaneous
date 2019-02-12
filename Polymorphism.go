package main

import "fmt"

type Parent interface {
	M1() string
}

type ChildClass struct {
	name string
}

func (t ChildClass) M1() string {
	return t.name
}

func Hello(p Parent) {
	fmt.Printf("Hi, my name is %s\n", p.M1())
}

func main() {
	Hello(ChildClass{name: "Sirisha"})
}
