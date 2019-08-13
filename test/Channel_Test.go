package main

import (
	"fmt"
	"sync"
	"time"
)

func testWaitGroup() {
	ch := make(chan int, 100)

	group := sync.WaitGroup{}
	group.Add(2)
	go func() {
		i := 100
		for i > 0 {
			ch <- i
			i--
		}
		close(ch)
		group.Done()
	}()

	go func() {
		for ch != nil {
			select {
			case i, ok := <-ch:
				if ok {
					fmt.Println(i)
				} else {
				}
			}
		}
		group.Done()
	}()

	group.Wait()
}

func testNesChannel() {
	c := new(chan int)

	i := 10
	for i > 0 {
		*c <- i
		i--
	}

}
func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}
func main() {
	panic()
	//ch := make(chan int)
	ch := make(chan int, 10)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}
