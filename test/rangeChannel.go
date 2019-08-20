package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 10)
	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			i++
			close(c)
		}
	}()
	for i := range c {
		fmt.Println(i)
	}
}
