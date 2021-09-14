package main

import (
	"fmt"
	"strings"
)

/**
524. 通过删除字母匹配到字典里最长单词
给你一个字符串 s 和一个字符串数组 dictionary 作为字典，找出并返回字典中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。

如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。



示例 1：

输入：s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
输出："apple"
示例 2：

输入：s = "abpcplea", dictionary = ["a","b","c"]
输出："a"


提示：

1 <= s.length <= 1000
1 <= dictionary.length <= 1000
1 <= dictionary[i].length <= 1000
s 和 dictionary[i] 仅由小写英文字母组成

*/

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}))
	fmt.Println(findLongestWord("abce", []string{"abe", "abc"}))

}
func findLongestWord(s string, dictionary []string) string {

	ans := -1
	maxLen := -1
	for i, _ := range dictionary {
		s2 := dictionary[i]
		sIndex, dIndex := 0, 0
		for sIndex < len(s) && dIndex < len(s2) {
			if s[sIndex] == s2[dIndex] {
				sIndex++
				dIndex++
			} else {
				sIndex++
			}

		}
		// 匹配到了单词 需要比较单词长度
		if dIndex == len(s2) {
			if dIndex > maxLen {
				maxLen = dIndex
				ans = i
			} else if dIndex == maxLen {
				if strings.Compare(dictionary[ans], dictionary[i]) > 0 {
					ans = i
				}
			}
		}
	}
	if ans == -1 {
		return ""
	}

	return dictionary[ans]
}
