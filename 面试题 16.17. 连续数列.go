package main

import "fmt"

/**
面试题 16.17. 连续数列
给定一个整数数组，找出总和最大的连续数列，并返回总和。

示例：

输入： [-2,1,-3,4,-1,2,1,-5,4]
输出： 6
解释： 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶：

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/
func main() {

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	maxSubArray := func(nums []int) int {

		ans := nums[0]
		curSum := nums[0]

		for i := 1; i < len(nums); i++ {
			if curSum >= 0 {
				curSum = nums[i] + curSum
			} else {
				curSum = nums[i]
			}
			ans = max(ans, curSum)
		}

		return ans
	}

	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
