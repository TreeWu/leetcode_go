package main

/**
538. 把二叉搜索树转换为累加树
给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。



例如：

输入: 原始二叉搜索树:
              5
            /   \
           2     13

输出: 转换为累加树:
             18
            /   \
          20     13
*/
func main() {
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	var convertBST func(root *TreeNode) *TreeNode

	convertBST = func(root *TreeNode) *TreeNode {
		sum := 0
		var dsf func(root *TreeNode)
		dsf = func(root *TreeNode) {
			if root != nil {
				dsf(root.Right)
				root.Val = sum + root.Val
				sum = root.Val
				dsf(root.Left)
			}
		}
		dsf(root)
		return root
	}

}
