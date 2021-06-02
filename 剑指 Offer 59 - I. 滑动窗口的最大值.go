package main

import (
	"fmt"
)

/**
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
 

提示：

你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	//fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 8))
	fmt.Println(maxSlidingWindow([]int{3, 1, 2}, 2))
}



/**
根据 最开始的窗口 初始化最大值指针，在窗口右移时，判断
 */
func maxSlidingWindow(nums []int, k int) []int {
	numLen := len(nums) - k
	var result = make([]int, numLen+1)

	// 查找窗口中最大值的下标
	maxIndex := func(ns []int, begin, end int) int {
		curMaxIndex := begin
		for i := begin; i <= end; i++ {
			if nums[curMaxIndex] <= nums[i] {
				curMaxIndex = i
			}
		}
		return curMaxIndex
	}

	curMaxIndex := maxIndex(nums, 0, k-1)

	for left := 0; left <= numLen; left++ {
		right := left + k - 1

		//如果滑动过程中，入数比当前最大下边大
		if nums[curMaxIndex] <= nums[right] {
			curMaxIndex = right
		}

		//如果滑动过程中，最大值下标已经不处于窗口内，那么重新查找最大值下标
		if curMaxIndex < left {
			curMaxIndex = maxIndex(nums, left, right)
		}

		result[left] = nums[curMaxIndex]
	}

	return result
}
