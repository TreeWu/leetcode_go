package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	count := 10

	wg := sync.WaitGroup{}
	wg.Add(count * 2)
	for i := 0; i < count; i++ {
		go func() {
			fmt.Printf("[%d]", i)
			wg.Done()
		}()
		time.Sleep(time.Second)
	}

	for i := 0; i < count; i++ {
		go func(i int) {
			fmt.Printf("--%d--", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
