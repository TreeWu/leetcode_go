package main

/**
剑指 Offer 34. 二叉树中和为某一值的路径
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。



示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]


提示：

节点总数 <= 10000
*/
func main() {

	pathSum := func(root *TreeNode, sum int) [][]int {
		var res [][]int
		var dsf func(root *TreeNode, sum int, arr []int)
		dsf = func(root *TreeNode, sum int, arr []int) {
			if root == nil {
				return
			} else if sum == root.Val && root.Left == nil && root.Right == nil {
				res = append(res, append(arr, root.Val))
			} else {
				arr = append(arr, root.Val)
				ldst := make([]int, len(arr))
				copy(ldst, arr)
				dsf(root.Left, sum-root.Val, ldst)
				rdst := make([]int, len(arr))
				copy(rdst, arr)
				dsf(root.Right, sum-root.Val, rdst)
			}
		}
		dsf(root, sum, nil)

		return res
	}

}
