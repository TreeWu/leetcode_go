package main

/**
145. 二叉树的后序遍历
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
func main() {

	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	postorderTraversal := func(root *TreeNode) []int {

		var ans []int
		var post func(r *TreeNode)
		post = func(r *TreeNode) {
			if r == nil {
				return
			}
			post(r.Left)
			post(r.Right)
			ans = append(ans, r.Val)
		}
		post(root)

		return ans
	}
	postorderTraversal(nil)
}
