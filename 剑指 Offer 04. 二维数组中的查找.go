package main

import "fmt"

/**
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
示例:
现有矩阵 matrix 如下：
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。
给定 target = 20，返回 false。

限制：
0 <= n <= 1000
0 <= m <= 1000
*/
func main() {
	var matrix [][]int
	matrix = append(matrix, []int{-5})

	fmt.Println(findNumberIn2DArray(matrix, -5))

}

/**
因为 是二维有序数组，所以左边一定不比右边大，上面一定不比下面大
定位一个数时，我们需要和当前数比较，然后确定下一步应该的方向
以上面的数组为列，查找 5，如果从 左上角 1 开始找，比较后有两个方向，一个是向右 4，一个是向下2，但是我们不能同时往两个方向找，同理 从右下角找也是不行
假设我们从右上角开始找， 15 比 5 大，而且因为二维数组是从上往下递增的，所以向下的可能行排除，只能往左边找
因此我们可以从 左下 或者 右上角开始搜索
*/
func findNumberIn2DArray(matrix [][]int, target int) bool {

	if matrix == nil || len(matrix) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])

	// 查找的开始左边
	curCol, curRow := col-1, 0

	// 查找的结束条件
	for curCol >= 0 && curRow < row {
		if matrix[curRow][curCol] == target {
			return true
		}
		if matrix[curRow][curCol] > target {
			curCol--
		} else {
			curRow++
		}
	}

	return false

}
