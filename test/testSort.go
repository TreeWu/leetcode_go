package main

import (
	"fmt"
	"sort"
)

func main(){
	str()
}

func str(){
	strs:=[]string{"adb","abe"}
	sort.Strings(strs)
	fmt.Println(strs)
}
