package main

import (
	"fmt"
)

func main() {
	/*	is := []int{1,2}
		is = is[:len(is)-1]
		fmt.Println(is)*/
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
		cur := root
		for cur != nil || len(stack) != 0 {
			for cur != nil {
				stack = append(stack, cur)
				cur = cur.Left
			}
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, cur.Val)
			cur = cur.Right
		}
		return result
	}
	res := inorderTraversal(root)
	fmt.Println(res)
}

/**
给定一个二叉树，返回它的中序遍历。

示例:

输入: [1,null,2,3]
	      1
 		 /   \
		3    2
       /\
      4 5
输出: [1,3,2]
进阶:递归算法很简单，你可以通过迭代算法完成吗？
*/

// 利用栈来解决二叉树的中序遍历  左根右
// 解决方式
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stock := []*TreeNode{}
	cur := root
	for cur != nil || len(stock) != 0 {
		for cur != nil {
			stock = append(stock, cur)
			cur = cur.Left
		}
		cur = stock[len(stock)-1]
		stock = stock[:len(stock)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}

	return res
}
