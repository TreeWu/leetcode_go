package main

import "fmt"

type age struct {
	Age int
}

func main() {
	/*	for i := 0; i < 5; i++ {
		defer func(i int) { print(i) }(i)
	}*/


	beforeReturn := changeBeforeReturn(3)
	fmt.Println(beforeReturn)
	fmt.Println(trip(3))
	/*
		for i := 0; i < 5; i++ {
			defer print(i)
		}*/
}

func print(i int) {
	fmt.Println(i)
}

func changeBeforeReturn(in int) *age {
	result := &age{Age: in + 1}
	defer func() { result.Age = 10000 }()
	return result
}

func double(x int) int {
	return x + x
}
func trip(x int) (result int) {
	defer func() { result += 1 }()
	return double(x)
}
