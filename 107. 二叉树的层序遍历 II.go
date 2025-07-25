package main

// 107. 二叉树的层序遍历 II
// 中等
// 744
// 相关企业
// 给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
// 示例 1：
//
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[15,7],[9,20],[3]]
// 示例 2：
//
// 输入：root = [1]
// 输出：[[1]]
// 示例 3：
//
// 输入：root = []
// 输出：[]
//
// 提示：
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
func levelOrderBottom(root *TreeNode) [][]int {
	//先层序遍历，再反转
	order := levelOrder(root)

	for i := 0; i < len(order)/2; i++ {
		order[i], order[len(order)-1-i] = order[len(order)-1-i], order[i]
	}
	return order
}
