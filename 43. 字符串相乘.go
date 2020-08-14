package main

import (
	"fmt"
	"strconv"
)

/**
43. 字符串相乘
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/
func main() {
	addStrings := func(num1 string, num2 string) string {
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
	multiply := func(num1 string, num2 string) string {
		if num2 == "0" || num1 == "0" {
			return "0"
		}
		ans := "0"
		for index := len(num1) - 1; index >= 0; index-- {
			i := index
			u, _ := strconv.Atoi(string(num1[index]))
			sum1 := "0"
			for index2 := len(num2) - 1; index2 >= 0; index2-- {
				j := index2
				p, _ := strconv.Atoi(string(num2[index2]))
				i := u * p
				itoa := strconv.Itoa(i)
				for j < len(num2)-1 {
					itoa += "0"
					j++
				}
				sum1 = addStrings(sum1, itoa)
			}

			for i < len(num1)-1 {
				sum1 += "0"
				i++
			}
			ans = addStrings(ans, sum1)
		}

		return ans
	}

	fmt.Println(multiply("123", "456"))

}
