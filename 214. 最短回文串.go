package main

import "fmt"

/**
给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。

示例 1:

输入: "aacecaaa"
输出: "aaacecaaa"
示例 2:

输入: "abcd"
输出: "dcbabcd"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(checkPalindrome("aca", 0, 2))
	fmt.Println(Reverse("abc"))
	fmt.Println(shortestPalindrome("abcd"))
}

/**
先找到以字符开头的最长回文字符，然后将剩余的部分倒叙插入字符串开头
*/
func shortestPalindrome(s string) string {
	if len(s) <= 1 || checkPalindrome(s, 0, len(s)-1) {
		return s
	}
	curIndex := len(s) - 1
	result := s[1:] + s
	for curIndex > 0 {
		s2 := Reverse(s[curIndex:]) + s
		if checkPalindrome(s2, 0, len(s2)-1) {
			result = s2
			break
		}
		curIndex--
	}

	return result
	/*	longestIndex, curIndex := 0, 0
		for curIndex < len(s) {
			if checkPalindrome(s, 0, curIndex) {
				longestIndex = curIndex
			}
			curIndex++
		}
		reverse := Reverse(s[longestIndex+1:])
		return reverse + s*/
}

func checkPalindrome(s string, left, right int) bool {

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
