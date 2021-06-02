package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Test interface {
	test1() int
	test2() int
}
type Student struct {
	Name string
	sex  int
}

func (s Student) test1() int {
	return 1
}
func (s *Student) test2() int {
	return 2
}
func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	ticker := time.NewTicker(time.Second)
loop:
	for {
		select {
		case _, ok := <-c:
			if !ok {
				break loop
			}
		case <-ticker.C:
			fmt.Println("Second")
		}
	}
	fmt.Println("end")

}

func testNil() {
	var w io.Writer

	w = os.Stdout
	if w == nil {
		fmt.Println("w == nil")
	} else {
		fmt.Println("w != nil")
	}

}
func f(v interface{}) {
	switch v.(type) {
	case Student, *Student:
		fmt.Println("值")
		// interface 类型通过 v.(具体类型)  完成转换
		fmt.Println(v.(*Student))
	}
}
