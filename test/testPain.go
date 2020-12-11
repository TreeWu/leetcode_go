package main

import "fmt"

type People struct {
	name string
}

func (p *People) toString() string {
	return "string:" + p.name
}

func main() {
	testDefer()
}

func testDefer() {
	p := People{name: "tom"}
	fmt.Printf("init %p \n", &p)
	defer func() { fmt.Println("defer1", &p) }()
	defer func(p1 People) { fmt.Println("defer2", &p1) }(p)
	defer func(p1 *People) { fmt.Println("defer3", &p1) }(&p)

	p.name = "tom2"
}

func testPanic() string {
	res := "init"
	fmt.Println("init", &res)
	defer func(res string) { fmt.Println("defer1", &res, res) }(res)
	res = "init2"

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("panic", &res)
			res = "panic"
		}
	}()
	defer func(res string) { fmt.Println("defer2", &res, res) }(res)
	panic("panic")

	return res
}
