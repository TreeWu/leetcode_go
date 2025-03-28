package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 100
	group := sync.WaitGroup{}
	group.Add(3)

	a, b, c := make(chan interface{}), make(chan interface{}), make(chan interface{})

	p := func(s string, waitC, notifyC chan interface{}) {
		defer group.Done()
		for i := 1; i <= count; i++ {
			<-waitC // 等待自己的信号
			fmt.Printf("%s%d ", s, i)
			if i < count || s != "C" { // 防止最后一次循环后还发送信号
				notifyC <- struct{}{}
			}
		}
	}
	go p("A", a, b)
	go p("B", b, c)
	go p("C", c, a)

	a <- new(interface{})

	group.Wait()
}
