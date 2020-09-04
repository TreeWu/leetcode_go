package main

import (
	"fmt"
	"strconv"
)

/**
257. 二叉树的所有路径
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
*/
func main() {
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	var binaryTreePaths func(root *TreeNode) []string

	binaryTreePaths = func(root *TreeNode) []string {
		var result []string

		var dsf func(head *TreeNode, preStr string)
		dsf = func(head *TreeNode, preStr string) {
			if head == nil {
				return
			}
			if preStr != "" {
				preStr += "->"
			}
			preStr = preStr + strconv.Itoa(head.Val)
			if head.Left == nil && head.Right == nil {
				result = append(result, preStr)
				return
			}
			dsf(head.Left, preStr)
			dsf(head.Right, preStr)
		}
		return result
	}

	fmt.Println(binaryTreePaths(&TreeNode{Val: 11, Right: &TreeNode{Val: 12}}))

}
