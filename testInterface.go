package main

import (
	"fmt"
)

type Student struct {
	Name string
	sex  int
}

type Myint int

func main() {

	var i int = 5
	var mi Myint = 6
	var _ int = int(mi)
	var _ Myint = Myint(i)
	f(Student{"TM", 2}) // 具体类型，通过隐形转换，成为interface
}

func f(v interface{}) {
	switch v.(type) {
	case Student, *Student:
		fmt.Println("值")
		// interface 类型通过 v.(具体类型)  完成转换

		fmt.Println(v.(Student))
	}
}
