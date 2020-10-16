package main

import (
	"fmt"
)

func appendSlice(slice **[]int) {
	fmt.Println(slice)
	tmp := append(**slice, (**slice)...)
	t:=&tmp
	slice =&t
	fmt.Println(slice)
}

func main() {
	var arr [100]int

	a1:=arr[90:]
	a2:=arr[:10]

	fmt.Printf("%d , %d \n",len(a1),cap(a1))
	fmt.Printf("%d , %d \n",len(a2),cap(a2))


	a1=append(a1,1)
	fmt.Printf("%d , %d \n",len(a1),cap(a1))


	a2=append(a2,2)
	fmt.Printf("%d , %d \n",len(a2),cap(a2))
}
