package main

import "fmt"

//9*9乘法表
func main() {

	j := 1
	for j <= 9 {
		i := 1
		for i <= j {
			fmt.Printf("%d ", i*j)
			i++
		}
		fmt.Println()
		j++
	}
}
