package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {

	s = strings.TrimRight(s, " ")
	fmt.Println(len(s))
	l := strings.LastIndex(s, " ")
	fmt.Println(l)
	if l == -1 {
		return len(s)
	}
	return len(s) - l - 1

}

func main() {
	fmt.Println(lengthOfLastWord("a "))
}
