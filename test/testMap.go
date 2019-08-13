package main

import "fmt"

type Stu struct {
	Name string
}

func main() {
	stus := new(map[string]Stu)

	fmt.Println(stus)
	s1 := Stu{"tom"}
	if v, ok := (*stus)["tom"]; !ok {   //new出来的map 可以取值，不过一直都是false
		fmt.Println(v)
	}
	(*stus)["tom"] = s1 //new出来的map不能赋值

}

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

/*func main() {

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
}*/
