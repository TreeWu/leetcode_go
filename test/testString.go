package main

import "fmt"

func main()  {

	var arr *[]int =new([]int)

	ints := append((*arr), 0)
	fmt.Printf("%T,%v",arr,arr)

	fmt.Printf("%T,%v",ints,ints)
}
