package main

func main() {

}

/**
1269. 停在原地的方案数
有一个长度为 arrLen 的数组，开始有一个指针在索引 0 处。

每一步操作中，你可以将指针向左或向右移动 1 步，或者停在原地（指针不能被移动到数组范围外）。

给你两个整数 steps 和 arrLen ，请你计算并返回：在恰好执行 steps 次操作以后，指针仍然指向索引 0 处的方案数。

由于答案可能会很大，请返回方案数 模 10^9 + 7 后的结果。



示例 1：

输入：steps = 3, arrLen = 2
输出：4
解释：3 步后，总共有 4 种不同的方法可以停在索引 0 处。
向右，向左，不动
不动，向右，向左
向右，不动，向左
不动，不动，不动
示例  2：

输入：steps = 2, arrLen = 4
输出：2
解释：2 步后，总共有 2 种不同的方法可以停在索引 0 处。
向右，向左
不动，不动
示例 3：

输入：steps = 4, arrLen = 2
输出：8

1 <= steps <= 500
1 <= arrLen <= 10^6
*/
/**
这类问题应该使用 动态规划求解
假设 dp[step][arrlen] 表示在step步后，指针指向arrlen的解法数 ，那么答案就是 dp[step][0]
利用动态的思想 dp[i][j]  可以由 dp[i-1][] 的三种情况转化过来
1、可以原地不动，那么 d[i-1][j]
2、在右边  那么  d[i-1][j+1]
3、在左边  那么  d[i-1][j-1]

所以dp[i][j] = dp[i-1][j]+dp[i-1][j+1] + dp[i-1][j-1]

边界确定
如果 i==j  dp[i][j] =1
如果 i<j dp[i][j] =0

*/
func numWays(steps int, arrLen int) int {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	const mod = 1e9 + 7
	maxColumn := min(arrLen-1, steps)
	dp := make([][]int, steps+1)
	for i := range dp {
		dp[i] = make([]int, maxColumn+1)
	}
	dp[0][0] = 1
	for i := 1; i <= steps; i++ {
		for j := 0; j <= maxColumn; j++ {
			dp[i][j] = dp[i-1][j]
			if j-1 >= 0 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % mod
			}
			if j+1 <= maxColumn {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod
			}
		}
	}
	return dp[steps][0]

}
