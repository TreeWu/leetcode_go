package main

import "fmt"

type A struct {
	name string
}

func (a *A)say()   {
	fmt.Printf(a.name)
}

type B A
type myint int

func main() {

	b:=&B{"a"}
	(*A)(b).say()
}

func add(a, b int) int {
	return a + b
}
