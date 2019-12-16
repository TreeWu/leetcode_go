package main

import "fmt"

func main() {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0}}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}

/**
在二维数组grid中，grid[i][j]代表位于某处的建筑物的高度。 我们被允许增加任何数量（不同建筑物的数量可能不同）的建筑物的高度。 高度 0 也被认为是建筑物。

最后，从新数组的所有四个方向（即顶部，底部，左侧和右侧）观看的“天际线”必须与原始数组的天际线相同。 城市的天际线是从远处观看时，由所有建筑物形成的矩形的外部轮廓。 请看下面的例子。

建筑物高度可以增加的最大总和是多少？

例子：
输入： grid = [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]
输出： 35
解释：
The grid is:
[ [3, 0, 8, 4],
  [2, 4, 5, 7],
  [9, 2, 6, 3],
  [0, 3, 1, 0] ]

从数组竖直方向（即顶部，底部）看“天际线”是：[9, 4, 8, 7]
从水平水平方向（即左侧，右侧）看“天际线”是：[8, 7, 9, 3]

在不影响天际线的情况下对建筑物进行增高后，新数组如下：

gridNew = [ [8, 4, 8, 7],
            [7, 4, 7, 7],
            [9, 4, 8, 7],
            [3, 3, 3, 3] ]
*/
/**
思路，求出行列最大值，然后遍历每个数，把这个数与其该行该列的最小值差值，加到结果里面去
*/
func maxIncreaseKeepingSkyline(grid [][]int) int {
	result := 0
	row := make([]int, len(grid))
	column := make([]int, len(grid))

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > row[i] {
				row[i] = grid[i][j]
			}
			if grid[j][i] > column[i] {
				column[i] = grid[j][i]
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			max := min(row[i], column[j])
			result += max - grid[i][j]
		}
	}

	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
