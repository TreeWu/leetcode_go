package test

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer func(i int) { print(i) }(i)
	}

	for i := 0; i < 5; i++ {
		defer func() { print(i) }()
	}

	for i := 0; i < 5; i++ {
		defer print(i)
	}
}

func print(i int) {
	fmt.Println(i)
}
