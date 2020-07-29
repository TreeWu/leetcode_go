package main

import "fmt"

var result []int = make([]int, 8)

func main() {

}

func cal8queens(row int) {
	if row == 8 {
		return
	}

}

func printQueens(result []int) {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if result[row] == column {
				fmt.Println("Q ")
			} else {
				fmt.Println()
			}
		}
	}
}
