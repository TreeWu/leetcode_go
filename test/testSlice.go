package main

import (
	"fmt"
)

type Q struct {
	name string
}

func appendSlice(slice **[]int) {
	fmt.Println(slice)
	tmp := append(**slice, (**slice)...)
	t := &tmp
	slice = &t
	fmt.Println(slice)
}

func main() {
	qs := []Q{{name: "1"}, {name: "2"}, {name: "3"}}
	fmt.Println(qs)
	for i := range qs {
		q := qs[i]
		q.name = q.name + "f"
		qs[i] = q
	}
	fmt.Println(qs)

}
