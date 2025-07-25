package main

import "strconv"

/*
*
给定两个字符串形式的非负整数num1 和num2，计算它们的和。

注意：

num1 和num2的长度都小于 5100.
num1 和num2 都只包含数字0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库，也不能直接将输入的字符串转换为整数形式。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

func addStrings(num1 string, num2 string) string {
	add := 0
	result := ""
	for i, j := len(num1)-1, len(num2)-1; j >= 0 || i >= 0 || add != 0; i, j = i-1, j-1 {
		var a1, a2 int
		if i >= 0 {
			a1 = int(num1[i] - '0')
		}
		if j >= 0 {
			a2 = int(num2[j] - '0')
		}
		sum := a1 + a2 + add
		result = strconv.Itoa(sum%10) + result
		add = sum / 10
	}
	return result
}
