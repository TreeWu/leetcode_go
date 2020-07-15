package main

func main() {
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	type rangeSumBST func(root *TreeNode, L int, R int) int

	var r rangeSumBST
	r = func(root *TreeNode, L int, R int) int {
		if root == nil {
			return 0
		}
		if root.Val < L || root.Val > R {
			return r(root.Right, L, R) + r(root.Left, L, R)
		} else {
			return root.Val + r(root.Right, L, R) + r(root.Left, L, R)
		}
	}
}
