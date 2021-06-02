package main

import (
	"fmt"
	"sync"
)

func main() {

	group := sync.WaitGroup{}
	group.Add(300)

	a2b, b2c, c2a := make(chan interface{}), make(chan interface{}), make(chan interface{})

	a := func() {
		i := 1
		for {
			select {
			case <-c2a:
				fmt.Println("A", i)
				i++
				group.Done()
				a2b <- new(interface{})
			}
		}
	}
	b := func() {
		i := 1
		for {
			select {
			case <-a2b:
				fmt.Println("B", i)
				i++
				group.Done()
				b2c <- new(interface{})
			}
		}
	}
	c := func() {
		i := 1
		for {
			select {
			case <-b2c:
				fmt.Println("C", i)
				i++
				group.Done()
				c2a <- new(interface{})
			}
		}
	}
	go a()
	go b()
	go c()
	c2a <- new(interface{})

	group.Wait()
}
