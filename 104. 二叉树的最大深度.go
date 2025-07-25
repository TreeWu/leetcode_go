package main

import (
	"fmt"
)

/*
*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明:叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

	  3
	 / \
	9  20
	  /  \
	 15   7

返回它的最大深度3 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	var maxDepth func(root *TreeNode) int

	/**
	递归解法，目前的最大深数，等于 左右分支最大深度 +1
	*/
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		if leftDepth > rightDepth {
			return leftDepth + 1
		}
		return rightDepth + 1
	}

	fmt.Println(maxDepth(nil))
}
