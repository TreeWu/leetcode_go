package main

import "fmt"

func main() {
	fmt.Println(sumNums(3))

}

func sumNums(n int) int {

	var ans int

	var sum func(a int) bool

	sum = func(a int) bool {

		ans += a

		return a > 1 && sum(a-1)
	}
	sum(n)
	return ans
}
