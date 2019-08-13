package main

import (
	"fmt"
)

type Student struct {
	Name string
	sex  int
}

type Myint int

type SecInt int

func main() {
	var i int
	var mi Myint
	var si SecInt
	i = 3
	mi = Myint(i)
	si = SecInt(i)
	mi = Myint(si)
	si = SecInt(mi)
	f(&Student{"TM", 2}) // 具体类型，通过隐形转换，成为interface
	var b  []byte = nil


	fmt.Println(len(b[0:]))
}

func f(v interface{}) {
	switch v.(type) {
	case Student, *Student:
		fmt.Println("值")
		// interface 类型通过 v.(具体类型)  完成转换
		fmt.Println(v.(*Student))
	}
}
