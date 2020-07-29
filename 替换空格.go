package main

import "fmt"

/**
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。



示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	replaceSpace := func(s string) string {

		var result []byte
		for i := 0; i < len(s); i++ {
			if s[i] == ' ' {
				result = append(result, []byte("%20")...)
			} else {
				result = append(result, s[i])
			}
		}
		return string(result)
	}

	fmt.Println(replaceSpace("fdas fa  df"))
}
