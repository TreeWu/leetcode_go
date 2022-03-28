package main

import (
	"math"
)

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
