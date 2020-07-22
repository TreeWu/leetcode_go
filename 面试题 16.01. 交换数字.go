package main

import "fmt"

func main() {

	s := []int{2, 3}
	fmt.Println(swapNumbers(s))
}
func swapNumbers(numbers []int) []int {
	numbers[0] = numbers[0] ^ numbers[1]
	numbers[1] = numbers[0] ^ numbers[1]
	numbers[0] = numbers[0] ^ numbers[1]
	return numbers
}
