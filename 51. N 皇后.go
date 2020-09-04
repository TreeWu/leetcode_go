package main

import "fmt"

/**
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。



上图为 8 皇后问题的一种解法。

给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

 

示例：

输入：4
输出：[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
解释: 4 皇后问题存在两个不同的解法。
 

提示：

皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-queens
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(solveNQueens(4))
}
func solveNQueens(n int) [][]string {
	var result [][]string

	posToString := func(pos []int) []string {
		length := len(pos)
		res := make([]string, length)
		for index := range res {
			s := make([]byte, length)
			for i := 0; i < length; i++ {
				if i == pos[index] {
					s[i] = 'Q'
					continue
				}
				s[i] = '.'
			}
			res[index] = string(s)
		}
		return res
	}

	// 记录每一行，哪个位置放了棋子 ,既以下标为 行，值为列，记录每一个皇后的位置
	rowPos := make([]int, n)
	for i := 0; i < n; i++ {
		rowPos[i] = -1
	}
	var dfs func(row int)
	dfs = func(row int) {
		for i := 0; i < n; i++ {
			//check
			ok := true
			for j := 0; j < row; j++ {
				if rowPos[j] == i || rowPos[j]-j == i-row || rowPos[j]+j == i+row { //同列，正斜 反斜上是否有值了
					ok = false
					break
				}
			}
			if ok {
				rowPos[row] = i
				if row == n-1 {
					result = append(result, posToString(rowPos))
				} else {
					dfs(row + 1)
				}
			}
		}
	}
	dfs(0)

	return result
}
