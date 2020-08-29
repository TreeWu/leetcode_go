package main

import "fmt"

/**

给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。

示例 1:

输入: [5,7]
输出: 4
示例 2:

输入: [0,1]
输出: 0
*/
func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
}

/**
暴力法 ：直接遍历所有数字与然后返回

二进制 ：每个数字都可以表示为二进制的模式
*/

func rangeBitwiseAnd(m int, n int) int {
	shift := 0
	for m < n {
		m, n = m>>1, n>>1
		shift++
	}
	return m << shift

}
