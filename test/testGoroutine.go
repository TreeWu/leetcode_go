package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(1)
	var group sync.WaitGroup

	for i := 0; i < 5; i++ {
		group.Add(1)
		go func(i int) {
			fmt.Println(" i " + strconv.Itoa(i))
			group.Done()
		}(i)
	}
	group.Wait()
}

//funcR main() {
//	listener, err := net.Listen("tcp", "localhost:8000")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			log.Print(err) // e.g., connection aborted
//			continue
//		}
//		handleConn(conn) // handle one connection at a time
//	}
//}
//func handleConn(c net.Conn) {
//	defer c.Close()
//	for {
//		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
//		if err != nil {
//			fmt.Println(err)
//			return // e.g., client disconnected
//		}
//		time.Sleep(1 * time.Second)
//	}
//}

/*func main() {
	go spinner(200 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
*/
