package main

import "fmt"

type Obj1 struct {
	S string
}
type Obj2 struct {
}

func (o Obj1) object() {}
func (o Obj2) object() {}

type Obj3 struct {
	Obj1
	Obj2
}

func NewObjP() *Obj1 {
	return &Obj1{}
}

func NewObj(s string) Obj1 {
	return Obj1{s}
}

func main() {
	p1 := NewObjP()
	p2 := NewObjP()
	fmt.Println(p1 == p2)

	o1:=NewObj("a")
	o2:=NewObj("a")
	fmt.Println(o1 == o2)
}
