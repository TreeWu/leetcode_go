package main

import "fmt"

/*
*
给定一个正整数n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 ×3 ×4 = 36。
说明: 你可以假设n不小于 2 且不大于 58。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(integerBreak3(100))
}

func integerBreak(n int) int {

	if n < 4 {
		return n - 1
	}
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = max(max(2*(i-2), 2*dp[i-2]), max(3*(i-3), 3*dp[i-3]))
	}
	return dp[n]

}

/*
*
动态规划
假设 n 可以拆分为 x + （n-x） ,那么分别求  x ,（n-x） 拆分整数后最大乘积就行
因为 0 ，1 的明显是没收益或者零收益的，所以  x 从 2 开始遍历
*/
func integerBreak2(n int) int {
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(i*(i-j), i*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}

/*
*
动态规划 递归
*/
func integerBreak3(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	curMax := 0
	for i := 2; i < n; i++ {
		curMax = max(curMax, max(i*(n-i), i*integerBreak3(n-i)))
	}
	return curMax
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
