package main

import (
	"fmt"
	"time"
)

func main() {
	print := func(i int) {
		for {
			fmt.Println(i)
		}
	}

	go print(1)
	go print(2)

	time.Sleep(time.Second)

	for  {
		fmt.Println("main")
	}
}
