package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Unix())
	fmt.Println(t.Second())
	fmt.Println(t.Nanosecond())
	fmt.Println(t.UnixNano())
}
