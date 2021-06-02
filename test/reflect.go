package main

import (
	"fmt"
	common "leetcode_go/base"
	"unsafe"
)

func main() {
	sizeof := unsafe.Sizeof(common.Tree{})
	fmt.Println(sizeof)
}
