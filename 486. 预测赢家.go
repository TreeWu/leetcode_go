package main

import (
	"fmt"
)

/*
*
给定一个表示分数的非负整数数组。 玩家 1 从数组任意一端拿取一个分数，随后玩家 2 继续从剩余数组任意一端拿取分数，然后玩家 1 拿，…… 。每次一个玩家只能拿取一个分数，分数被拿取之后不再可取。直到没有剩余分数可取时游戏结束。最终获得分数总和最多的玩家获胜。

给定一个表示分数的数组，预测玩家1是否会成为赢家。你可以假设每个玩家的玩法都会使他的分数最大化。

示例 1：

输入：[1, 5, 2]
输出：False
解释：一开始，玩家1可以从1和2中进行选择。
如果他选择 2（或者 1 ），那么玩家 2 可以从 1（或者 2 ）和 5 中进行选择。如果玩家 2 选择了 5 ，那么玩家 1 则只剩下 1（或者 2 ）可选。
所以，玩家 1 的最终分数为 1 + 2 = 3，而玩家 2 为 5 。
因此，玩家 1 永远不会成为赢家，返回 False 。
示例 2：

输入：[1, 5, 233, 7]
输出：True
解释：玩家 1 一开始选择 1 。然后玩家 2 必须从 5 和 7 中进行选择。无论玩家 2 选择了哪个，玩家 1 都可以选择 233 。

	最终，玩家 1（234 分）比玩家 2（12 分）获得更多的分数，所以返回 True，表示玩家 1 可以成为赢家。

提示：

1 <= 给定的数组长度<= 20.
数组里所有分数都为非负数且不会大于 10000000 。
如果最终两个玩家的分数相等，那么玩家 1 仍为赢家。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/predict-the-winner
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(dpMax([]int{1, 5, 2}))
	fmt.Println(PredictTheWinner([]int{1, 5, 2}))
}

func PredictTheWinner(nums []int) bool {
	return maxSum(0, len(nums)-1, nums) >= 0
}

/*
*
动态规划：
定义一个数组 DP[i][j] ，其中 i 代表左下标 ，j 代表右下标，该数组的值表示 该轮选择的最优解
容易得出，当 i ==j 时，玩家没得选择，所以 dp[i][j] = nums[i]
另外  DP[i][j] = max(nums[i] - dp[i+1][j],nums[j]-dp[i][j-1])
*/
func dpMax(nums []int) bool {
	length := len(nums)
	dp := make([][]int, len(nums))
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = nums[i]
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][length-1] >= 0
}

/**
递归法：
能不能取胜，在于自己取值的总和是不是比对方要大，所以可以用差值计算
将大问题分解成小问题，当前回合能不能赢，取决于当前选择的值，是否大于 下一个回合对手的放回值
*/

func maxSum(start, end int, nums []int) int {
	if start == end {
		return nums[start]
	}
	left := nums[start] - maxSum(start+1, end, nums)
	right := nums[end] - maxSum(start, end-1, nums)
	return max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
