package main

import (
	"fmt"
	"io"
	"os"
)

type Student struct {
	Name string
	sex  int
}

type Myint int

type SecInt int


func main() {
	testNil()
	/*	var i int
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

	*/
}
func testNil()  {
	var w io.Writer

	w= os.Stdout
	if w==nil {
		fmt.Println("w == nil")
	}else {
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
