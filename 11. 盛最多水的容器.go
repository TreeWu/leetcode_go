package main

import "fmt"

/*
*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0)。找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且n的值至少为 2。

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。

示例：

输入：[1,8,6,2,5,4,8,3,7]
输出：49
通过次数215,856

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water

双指针法，左右两边
*/
func main() {
	var ints = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(ints))
}

// 双指针法 ，从两边开始往中间收缩，面积 = 指针距离 * 最小的高度
// 每次遍历，都将最小高度往中间移，直到指针重合

func maxArea(height []int) int {
	var result = 0
	var l, r = 0, len(height) - 1

	for l < r {
		max := min(height[r], height[l]) * (r - l)
		if max > result {
			result = max
		}

		if height[r] > height[l] {
			l++
		} else {
			r--
		}
	}
	return result
}
