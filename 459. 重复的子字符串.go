package main

import (
	"fmt"
	"strings"
)

/**

给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例 1:

输入: "abab"

输出: True

解释: 可由子字符串 "ab" 重复两次构成。
示例 2:

输入: "aba"

输出: False
示例 3:

输入: "abcabcabcabc"

输出: True

解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
通过次数48,307提交次数93,813
*/
func main() {
	fmt.Println(repeatedSubstringPattern("abab"))
}

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {

		childLength := i + 1
		if length%childLength != 0 {
			continue
		}
		childS := s[:i+1]
		count := strings.Count(s, childS)
		if count == length/childLength {
			return true
		}
	}
	return false
}
