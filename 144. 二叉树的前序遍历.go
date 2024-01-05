package main

func preorderTraversal(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}
	var stock []*TreeNode
	stock = append(stock, root)
	for len(stock) != 0 {
		cur := stock[len(stock)-1]
		stock = stock[:len(stock)-1]
		ret = append(ret, cur.Val)
		if cur.Right != nil {
			stock = append(stock, cur.Right)
		}
		if cur.Left != nil {
			stock = append(stock, cur.Left)
		}
	}

	return ret
}
