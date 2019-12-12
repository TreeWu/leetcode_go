package main

import "fmt"

func judgeCircle(moves string) bool {
	row, column := 0, 0
	ss := []byte(moves)
	for _, s := range ss {
		if s == 'U' {
			row++
		} else if s == 'D' {
			row--
		} else if s == 'L' {
			column++
		} else if s == 'R' {
			column--
		}
	}
	return row == 0 && column == 0
}
func main() {
	fmt.Println(judgeCircle("LL"))
}
