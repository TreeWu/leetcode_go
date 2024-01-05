package main

import (
	"container/list"
	"fmt"
)

/*
*
103. 二叉树的锯齿形层序遍历
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

	  3
	 / \
	9  20
	  /  \
	 15   7

返回锯齿形层序遍历如下：

[

	[3],
	[20,9],
	[15,7]

]
*/
func main() {
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	var zigzagLevelOrder func(root *TreeNode) [][]int
	zigzagLevelOrder = func(root *TreeNode) [][]int {
		if root == nil {
			return [][]int{}
		}
		layer := 0
		res := [][]int{{}}
		l := list.New()
		rightFlag := true
		l.PushFront(root)
		curLen := 1
		for curLen != 0 {
			if rightFlag {
				front := l.Front()
				l.Remove(front)

				node := front.Value.(*TreeNode)
				res[layer] = append(res[layer], node.Val)
				if node.Left != nil {
					l.PushBack(node.Left)
				}
				if node.Right != nil {
					l.PushBack(node.Right)
				}

			} else {
				front := l.Back()
				l.Remove(front)
				node := front.Value.(*TreeNode)
				res[layer] = append(res[layer], node.Val)
				if node.Right != nil {
					l.PushFront(node.Right)
				}
				if node.Left != nil {
					l.PushFront(node.Left)
				}
			}

			if curLen == 1 {
				curLen = l.Len()
				rightFlag = !rightFlag

				if curLen != 0 {
					res = append(res, []int{})
					layer++
				}
			} else {
				curLen--
			}

		}
		return res
	}

	fmt.Println(zigzagLevelOrder(&TreeNode{Val: 3, Right: &TreeNode{Val: 9}, Left: &TreeNode{Val: 20, Right: &TreeNode{Val: 15}, Left: &TreeNode{Val: 7}}}))
}
