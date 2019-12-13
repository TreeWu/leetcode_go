package main

import "fmt"

func subtractProductAndSum(n int) int {
	begin := true
	sum, mul := 0, 0
	for n != 0 {
		tmp := n % 10
		n = n / 10
		sum = sum + tmp
		if begin {
			mul = 1
			begin = false
		}
		mul = mul * tmp
	}
	return mul - sum

}
func main() {
	fmt.Println(subtractProductAndSum(234))
}
