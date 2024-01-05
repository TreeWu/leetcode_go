package main

import "fmt"

/*
*
501. 二叉搜索树中的众数
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

	1
	 \
	  2
	 /
	2

返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
*/
func main() {
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	findMode := func(root *TreeNode) []int {
		var res []int

		var arr []int
		var mid func(r *TreeNode)

		mid = func(r *TreeNode) {
			if r != nil {
				mid(r.Left)
				arr = append(arr, r.Val)
				mid(r.Right)
			}
		}
		mid(root)
		max := 0
		curMax := 0
		preNum := 0

		for i := range arr {

			if arr[i] == preNum {
				curMax++
			} else {
				preNum = arr[i]
				curMax = 1
			}

			if curMax > max {
				res = []int{preNum}
				max = curMax
			} else if curMax == max {
				res = append(res, preNum)
			}

		}

		return res
	}

	fmt.Println(findMode(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
}
