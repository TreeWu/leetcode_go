package main

/**
剑指 Offer 44. 数字序列中某一位的数字
数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

请写一个函数，求任意第n位对应的数字。



示例 1：

输入：n = 3
输出：3
示例 2：

输入：n = 11
输出：0


限制：

0 <= n < 2^31
注意：本题与主站 400 题相同：https://leetcode-cn.com/problems/nth-digit/
*/
func main() {

}
func findNthDigit(n int) int {
	numLen := func(num int) int {
		length := 0
		if num > 10 {
			num = num / 10
			length++
		}
		length++
		return length
	}

	length := 0
	num := 0
	for length < n {
		length += numLen(num)
		num++
	}
	i := length - n


	num = num /10

	return num %10

}
