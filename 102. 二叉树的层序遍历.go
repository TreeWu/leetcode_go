package main

//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
//
//
//示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
//示例 2：
//
//输入：root = [1]
//输出：[[1]]
//示例 3：
//
//输入：root = []
//输出：[]
//
//
//提示：
//
//树中节点数目在范围 [0, 2000] 内
//-1000 <= Node.val <= 1000

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	stock := []*TreeNode{root}

	for len(stock) != 0 {

		var tmp []int
		cursize := len(stock)
		for i := 0; i < cursize; i++ {
			cur := stock[0]
			stock = stock[1:]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				stock = append(stock, cur.Left)
			}
			if cur.Right != nil {
				stock = append(stock, cur.Right)
			}
		}
		ret = append(ret, tmp)
	}
	return ret
}
