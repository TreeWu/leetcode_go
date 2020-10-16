package main

/**
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

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

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	var pathSum func(root *TreeNode, sum int) [][]int
	pathSum = func(root *TreeNode, sum int) [][]int {
		var ans [][]int

		var dsf func(root *TreeNode, sum int, arr []int)
		dsf = func(root *TreeNode, sum int, arr []int) {
			if root == nil {
				return
			}
			if sum == root.Val && root.Left == nil && root.Right == nil {
				ans = append(ans, append(arr, root.Val))
				return
			}

			arr = append(arr, root.Val)
			ldst := make([]int, len(arr))
			copy(ldst, arr)
			dsf(root.Left, sum-root.Val, ldst)
			rdst := make([]int, len(arr))
			copy(rdst, arr)
			dsf(root.Right, sum-root.Val, rdst)

		}
		dsf(root, sum, nil)
		return ans
	}
}
