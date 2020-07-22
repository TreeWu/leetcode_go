package main

import (
	"fmt"
	"strings"
)

/**
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"、"-1E-16"及"12e+5.4"都不是。

 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

	fmt.Println(isNumber("12e+5.4"))
	fmt.Println(isNumber("-1E-16"))
	fmt.Println(isNumber("1.2.3"))
	fmt.Println(isNumber("1a3.14"))
	fmt.Println(isNumber("12e"))
	fmt.Println(isNumber("1 "))
	fmt.Println(isNumber("  "))

}
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	runes := []rune(s)
	hasE := false
	hasPoint := false
	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case '-', '+': // 符号不是在开头
			if i != 0 {
				return false
			}
		case 'e':
			if i == 0 || hasE || i == len(runes)-1 { //  e 在开头,两个 e ,e在结尾 不算
				return false
			}
			hasE = true
		case '.':
			if hasPoint { // 两个 点，
				return false
			}
			hasPoint = true
		default:
			if '0' > runes[i] || runes[i] > '9' { // 其他不在 1~9范围内的
				return false
			}
		}

	}
	return true
}
