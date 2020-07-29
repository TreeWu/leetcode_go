package main

import "fmt"

func main() {
	c := &Child{}
	c.Say()
	c.SayReal(c)

	var name Name = &Child{}
	fmt.Println(name.Name())

}

type Name interface {
	Name() string
}

type Parent struct {
}

func (self *Parent) Say() {
	fmt.Println(self.Name())
}

func (salf *Parent) SayReal(child Name) {
	fmt.Println(child.Name())
}

func (self *Parent) Name() string {
	return "I'm Parent"
}

type Child struct {
	Parent
}

func (self *Child) Name() string {
	return "I'm Child"
}
