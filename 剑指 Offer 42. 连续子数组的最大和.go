package main

import "fmt"

/**
剑指 Offer 42. 连续子数组的最大和
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。



示例1:

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。


提示：

1 <= arr.length <= 10^5
-100 <= arr[i] <= 100

*/

func main() {
fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

/**
每一个数字在和前面的结果相加时都判断一下，如果前面的和是 负数，那就不需要相加了
*/
func maxSubArray(nums []int) int {

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	ans := nums[0]
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if sums[i-1] > 0 {
			sums[i] = sums[i-1]
		}
		sums[i] += nums[i]
		ans = max(ans, sums[i])
	}
	return ans

}
