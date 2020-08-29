package main

import "fmt"

/**
111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
*/
/**
注意是根节点到叶子节点的最短路径，意思是如果一个节点只有左分支没有右分支，那么右分支不能作为终点，只能继续找左节点的最小深度
*/
func main() {

	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	minDepth := func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		min := func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}
		var deep func(root *TreeNode) int
		deep = func(root *TreeNode) int {
			if root == nil {
				return 0
			}
			if root.Right == nil {
				return deep(root.Left) + 1
			} else if root.Left == nil {
				return deep(root.Right) + 1
			} else {
				return min(deep(root.Right), deep(root.Left)) + 1
			}
		}

		return deep(root)
	}

	fmt.Println(minDepth(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}))
}
