package main

import "fmt"

// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
//
// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
//
// 示例 1：
//
// 输入：s = ["h","e","l","l","o"]
// 输出：["o","l","l","e","h"]
// 示例 2：
//
// 输入：s = ["H","a","n","n","a","h"]
// 输出：["h","a","n","n","a","H"]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/reverse-string
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {
	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))
}

func reverseString(s []byte) {
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
