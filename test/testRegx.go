package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	a := "app::product::info"
	b := "app::product::.....*"
	b = strings.ReplaceAll(b, "*", ".*")
	fmt.Println(b)
	fmt.Println(regexp.Compile(b))

}
