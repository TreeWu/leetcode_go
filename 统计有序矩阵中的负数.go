package main

import "fmt"

func main() {
	//grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	grid := [][]int{{5,1,0},{-5,-5,-5}}
	fmt.Println(countNegatives(grid))
}

func countNegatives(grid [][]int) int {
	column := len(grid[0])
	row := len(grid)
	i, j := column-1, 0
	sum := 0
	for i >= 0 && j < row {
		if grid[j][i] < 0 {
			sum += row - j
			i--
		} else {
			j++
		}
	}
	return sum
}
