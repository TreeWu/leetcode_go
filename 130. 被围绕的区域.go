package main

import "fmt"

/**
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X
解释:

被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surrounded-regions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	/*	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}*/
	board := [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
	}
	printBoard(board)
	fmt.Println()
	solve(board)
	printBoard(board)
}

/**
农村包围城市的策略
先从外圈遍历，如果遇到 将 O 标记为 A，然后往 上下左右 四个方向 遍历
*/
func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	row, column := len(board), len(board[0])
	var dsf func(board [][]byte, x, y int)
	dsf = func(board [][]byte, x, y int) {
		if x < 0 || x > column-1 || y < 0 || y > row-1 {
			return
		}
		if board[y][x] == 'O' {
			board[y][x] = 'A'
			dsf(board, x, y+1)
			dsf(board, x, y-1)
			dsf(board, x-1, y)
			dsf(board, x+1, y)
		}
	} 
	for i := 0; i < row; i++ {
		dsf(board, 0, i)
		dsf(board, column-1, i)
	}
	for i := 0; i < column; i++ {
		dsf(board, i, 0)
		dsf(board, i, row-1)
	}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}

}

func printBoard(board [][]byte) {
	for row := range board {
		for col := range board[row] {
			fmt.Printf("%s ", string(board[row][col]))
		}
		fmt.Println()
	}
}
