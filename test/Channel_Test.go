package main

import (
	"fmt"
)

func main() {
	chani := make(chan int, 11)
	i := 0
	for i < 10 {
		chani <- i
		i++
	}

	for ci := range chani {
		fmt.Println(ci)
	}

	fmt.Println(len(chani))

}

/*func main() {
	chani := make(chan int, 10)

	get := func(index int) {
		for {
			select {
			case i := <-chani:
				fmt.Printf("%d -> %d \n", index, i)
			}
		}
	}
	put := func() {
		i := 0
		for {
			chani <- i
			i++
			time.Sleep(time.Second)
		}
	}

	go get(1)
	go get(2)
	go put()

	select {}

}*/

/*func testWaitGroup() {
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

	//ch := make(chan int)
	ch := make(chan int, 10)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}
*/
