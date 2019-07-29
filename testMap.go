package main

import "fmt"

/*type Stu struct {
	Name string
}

func main() {
	stus := make(map[string]*Stu)
	s1 := &Stu{"tom"}
	stus["tom"] = s1
	s2 := stus["tom"]
	s2.Name = "jar"

	fmt.Printf("s1: %p    s2: %p \n", s1, &s2)
	fmt.Printf("s1: %v   s2: %v", s1, s2)
}
*/

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		fmt.Println(i)
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

func main() {

	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}
