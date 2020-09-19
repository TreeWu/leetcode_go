package main

import (
	"container/list"
	"fmt"
)

/**
计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-left-leaves
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	var sumOfLeftLeaves func(root *TreeNode) int
	sumOfLeftLeaves = func(root *TreeNode) int {
		ans := 0
		if root == nil {
			return 0
		}
		leftStock := list.New()
		cur := root
		for cur != nil || leftStock.Len() != 0 {
			for cur != nil {
				leftStock.PushFront(cur)
				cur = cur.Left
			}
			front := leftStock.Back()
			leftStock.Remove(front)
			c := front.Value.(*TreeNode)
			if c.Right == nil {
				ans += c.Val
			}else {
				leftStock.PushFront(c.Right)
			}
			cur = leftStock.Back().Value.(*TreeNode)
		}

		return ans
	}

	fmt.Println(sumOfLeftLeaves(&TreeNode{Val:3, Left: &TreeNode{Val: 9 }, Right: &TreeNode{Val: 20,Left: &TreeNode{Val: 17},Right: &TreeNode{Val: 7}}}))
}
