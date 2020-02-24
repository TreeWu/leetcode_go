package main

import "fmt"

func main() {
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	inorderTraversal := func(root *TreeNode) []int {
		var result []int
		if root == nil {
			return result
		}
		var stack []*TreeNode
		stack = append(stack, root)
		for len(stack) != 0 {
			cur := stack[len(stack)-1]
			for cur.Left != nil {
				stack = append(stack, cur.Left)
				cur = cur.Left
			}
			result = append(result, cur.Val)
			stack = stack[:len(stack)-1]
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
		}

		return result
	}
	traversal := inorderTraversal(root)
	fmt.Println(traversal)
}

/**
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
/*func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		result = append(result, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}

	return result
}
*/
