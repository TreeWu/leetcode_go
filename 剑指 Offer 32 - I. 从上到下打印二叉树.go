package main

import (
	"container/list"
	"fmt"
)

/**
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]


提示：

节点总数 <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
该题初看与 中序遍历 差不多，实则差别很大，中序是 一次打印一个树，但这题一次从左往右只打印一层

该题的解题思路是 使用链表，按照打印顺序 建构链表，
从左往右，把元素放入到链表中，然后从链头开始打印即可

也就是   广度优先   遍历

*/
func main() {

	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	var levelOrder func(root *TreeNode) []int

	levelOrder = func(root *TreeNode) []int {

		var result []int

		if root == nil {
			return result
		}
		stock := list.New()

		stock.PushFront(root)

		for stock.Front() != nil {
			front := stock.Front()
			stock.Remove(front)
			node := front.Value.(*TreeNode)
			result = append(result, node.Val)

			if node.Left != nil {
				stock.PushBack(node.Left)
			}

			if node.Right != nil {
				stock.PushBack(node.Right)
			}
		}
		return result
	}

	fmt.Println(levelOrder(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))

}
