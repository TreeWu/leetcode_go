package main

import (
	"fmt"
	"math"
)

/**
给定一个整数数组 nums，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释:连续子数组[4,-1,2,1] 的和最大，为6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/
/**
连续的子数组，需要遍历求和
我们假设前面的和是正的，
*/
func MaxSubArray(nums []int) int {
	ans := nums[0]
	sum := nums[0]
	for index := range nums {
		if sum > 0 {
			sum += nums[index]
		} else {
			sum = nums[index]
		}
		ans = int(math.Max(float64(ans), float64(sum)))
	}
	return ans
}

func main() {

	nums := []int{1, 2, 4, -5, 6}
	sum := MaxSubArray(nums)
	fmt.Println(sum)
}
