package main

import (
	"fmt"
	"sync"
)

func main()  {

	ch:=make(chan int ,100)

	group := sync.WaitGroup{}
	group.Add(2)
	go func() {
		i:=100
		for i>0{
			ch <-i
			i--
		}
		close(ch)
		group.Done()
	}()

	go func() {
		for ch!=nil  {
			select {
			case i,ok:=<-ch:
				if ok {
					fmt.Println(i)
				}else {
				}
			}
		}
		group.Done()
	}()

	group.Wait()
}