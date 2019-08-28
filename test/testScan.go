package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	test := "123\n345\n567"

	buff := bytes.NewBufferString(test)
	s1 := bufio.NewScanner(buff)
	s2 := bufio.NewScanner(buff)
	if s1.Scan() {
		fmt.Println(s1.Text())
	}

	if s2.Scan() {
		fmt.Println(s2.Text())
	}
}
