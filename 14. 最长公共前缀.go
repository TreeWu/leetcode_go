package main

import (
	"math"
)

/**
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。



示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。


提示：

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成
*/

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
