package main

import (
	"fmt"
	"math"
)

/**
平面上有 n 个点，点的位置用整数坐标表示 points[i] = [xi, yi]。请你计算访问所有这些点需要的最小时间（以秒为单位）。

你可以按照下面的规则在平面上移动：

每一秒沿水平或者竖直方向移动一个单位长度，或者跨过对角线（可以看作在一秒内向水平和竖直方向各移动一个单位长度）。
必须按照数组中出现的顺序来访问这些点。
输入：points = [[1,1],[3,4],[-1,0]]
输出：7
解释：一条最佳的访问路径是： [1,1] -> [2,2] -> [3,3] -> [3,4] -> [2,3] -> [1,2] -> [0,1] -> [-1,0]
从 [1,1] 到 [3,4] 需要 3 秒
从 [3,4] 到 [-1,0] 需要 4 秒
一共需要 7 秒
示例 2：

输入：points = [[3,2],[-2,2]]
输出：5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-time-visiting-all-points
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
解法，顺序递归求两点之间的距离
因为只允许走直线或对角线，所以可以折叠为两点之间，X轴或Y轴的最大距离，即为两点之间最小距离。
*/
func main() {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	fmt.Println(minTimeToVisitAllPoints(points))
}
func minTimeToVisitAllPoints(points [][]int) int {
	result := 0
	p := [2]int{points[0][0], points[0][1]}
	i := len(points)
	for j := 1; j < i; j++ {
		q := [2]int{points[j][0], points[j][1]}
		result += minTimeToVisitPoints(p, q)
		p = q
	}
	return result
}

func minTimeToVisitPoints(a, b [2]int) int {
	return maxAbs(a[0]-b[0], a[1]-b[1])
}

func maxAbs(a, b int) int {
	return int(math.Max(math.Abs(float64(a)), math.Abs(float64(b))))
}
