package main

import (
	"fmt"
	"strconv"
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
		qs[i].name = qs[i].name + "F"
	}
	fmt.Println(qs)

	qs = nil

	q := Q{}

	for i := 0; i < 3; i++ {
		q.name = strconv.Itoa(i)
		qs = append(qs, q)
	}
	fmt.Println(qs)

	p := q
	p.name = "ddd"

	fmt.Println(p, q)

}
