package main

import (
	"fmt"
	"math"
)

/**
110. 平衡二叉树
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/
func main() {

	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	var deep func(root *TreeNode) int
	deep = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDeep := deep(root.Left)
		if leftDeep == -1 {
			return -1
		}
		rightDeep := deep(root.Right)
		if rightDeep == -1 {
			return -1
		}
		absDeep := math.Abs(float64(leftDeep - rightDeep))
		if absDeep > 1 {
			return -1
		}
		return max(leftDeep, rightDeep) + 1

	}
	isBalanced := func(root *TreeNode) bool {
		return deep(root) != -1
	}

	fmt.Println(isBalanced(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}}}}))

}
