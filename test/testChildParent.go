package main

import "fmt"

type parent struct {
}

func (p *parent) Say() {
	fmt.Println("parent")
}

type child struct {
	parent
}

func (c *child) Say() {
	fmt.Println("child")
}

func main() {
	p := &parent{}
	p.Say() //parent
	c := &child{}
	c.parent.Say() //parent
	c.Say()        //child
}
