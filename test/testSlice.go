package main

import (
	"errors"
	"fmt"
)

func appendSlice(slice **[]int) {
	fmt.Println(slice)
	tmp := append(**slice, (**slice)...)
	t:=&tmp
	slice =&t
	fmt.Println(slice)
}

func main() {
	slice := &[]int{1}
	fmt.Println(slice)
	appendSlice(&slice)
	fmt.Println(slice)
	panic()
}
