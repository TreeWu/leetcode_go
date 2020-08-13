package main

import "fmt"

/**
给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。

重复出现的子串要计算它们出现的次数。

示例 1 :

输入: "00110011"
输出: 6
解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。

请注意，一些重复出现的子串要计算它们出现的次数。

另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
示例 2 :

输入: "10101"
输出: 4
解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
注意：

s.length 在1到50,000之间。
s 只包含“0”或“1”字符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-binary-substrings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(countBinarySubstrings2("10101"))
}

/**
满足条件的子字符串，满足对称条件，如 0011 满足条件，00111因为是 2 个0,3 个1所以不满足条件
另外 0011 的子串 01 也是满足条件的，注意观察发现，子串 满足数 为 0  1 连续出现相邻最小数

如 001110 可记为{2,3,1},相邻的两个数字表示不同的数字连续长度，相邻两两组合取最小值 为 2+1 =3

*/

func countBinarySubstrings(s string) int {

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	rs := []rune(s)

	ans := 0
	var arr []int

	var cur rune

	for i := range rs {
		if rs[i] == cur {
			arr[len(arr)-1] += 1
		} else {
			cur = rs[i]
			arr = append(arr, 1)
		}
	}
	for i := 1; i < len(arr); i++ {
		ans += min(arr[i-1], arr[i])
	}
	return ans
}

/**
进阶解法，因为统计数组每个数只需要和前面的数做比较，所以只需要记住上一个 字符 的连续数 和当前字符的连续数就行了
*/
func countBinarySubstrings2(s string) int {

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	rs := []rune(s)

	ans := 0
	last, pre := 0, 0

	var cur rune

	for i := range rs {
		if rs[i] == cur {
			pre++
		} else {
			ans += min(pre, last)
			cur = rs[i]

			last = pre
			pre = 1
		}
	}
	// 遍历到最后一个数时，因为没有转换条件，所以需要需要手动加上
	ans += min(pre, last)
	return ans
}
