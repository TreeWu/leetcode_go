package main

import (
	"math"
	"leetcode_go/base"
)

func main() {
	name := common.Name
}

func longestCommonPrefix(strs []string) string {

	minLen := math.MaxInt64
	for i := range strs {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}
	for i := 0; i < minLen; i++ {

	}

	return ""

}
