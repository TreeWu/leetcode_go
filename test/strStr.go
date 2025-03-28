package main

import "fmt"

/**
实现strStr()函数。

给定一个haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回 -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
*/

/*暴力遍历破解
 */
func strStr(haystack string, needle string) int {
	bh, bn := []byte(haystack), []byte(needle)

	i := 0
	for i <= len(bh)-len(bn) { // 保证查询长度长于目标长度，否则无效
		j := i
		for k := 0; k < len(bn); k++ {
			if bh[j] == bn[k] {
				j++
			} else {
				break
			}
		}
		if j-i == len(bn) { // 如果遍历次数 等于 目标长度，说明已经找到
			return i
		}
		i++
	}
	return -1
}

func main() {
	str := strStr("", "")
	fmt.Println(str)
}
