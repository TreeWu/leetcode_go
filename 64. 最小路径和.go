package main

import "fmt"

/**
64. 最小路径和
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/
func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}))
}

/**
此题使用DP求解

dp[x,y] = min (dp[x,y+1],dp[x+1,y]) + arr[x,y]

*/
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}


	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if j == n-1 {
				if i == m-1 {
					grid[i][j] = grid[i][j]
				} else {
					grid[i][j] = grid[i+1][j] + grid[i][j]
				}

			} else if i == m-1 {
				if j == n-1 {
					grid[i][j] = grid[i][j]
				} else {
					grid[i][j] = grid[i][j+1] + grid[i][j]
				}
			} else {
				grid[i][j] = min(grid[i][j+1], grid[i+1][j]) + grid[i][j]
			}
		}
	}
	return grid[0][0]
}
