package main

import (
	"fmt"
	"math"
)

/**
530. 二叉搜索树的最小绝对差
给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。



示例：

输入：

   1
    \
     3
    /
   2

输出：
1

解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
*/
func main() {

	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	/**
	中序遍历，
	 */
	getMinimumDifference := func(root *TreeNode) int {
		ans, pre := math.MaxInt32, -1
		var dsf func(root *TreeNode)
		dsf = func(root *TreeNode) {
			if root == nil {
				return
			}
			dsf(root.Left)
			if pre != -1 && ans > root.Val-pre {
				ans = root.Val - pre
			}
			pre = root.Val
			dsf(root.Right)
		}
		dsf(root)
		return ans
	}

	difference := getMinimumDifference(&TreeNode{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}})
	fmt.Println(difference)
}
