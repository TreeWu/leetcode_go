package main

import "fmt"

/**
106. 从中序与后序遍历序列构造二叉树
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/
func main() {

	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	var buildTree func(inorder []int, postorder []int) *TreeNode

	buildTree = func(inorder []int, postorder []int) *TreeNode {
		if len(inorder) == 0 || len(postorder) == 0 {
			return nil
		}
		curVal := postorder[len(postorder)-1]
		cur := &TreeNode{}
		cur.Val = curVal

		var leftIn, rightIn, leftPost, rightPost []int
		for i := range inorder {
			if curVal == inorder[i] {
				leftIn = inorder[:i]
				rightIn = inorder[i+1:]
				break
			}
		}

		if rightIn != nil && len(rightIn) != 0 {
			for i := range postorder {
				if rightIn[0] == postorder[i] {
					rightPost = postorder[i : len(postorder)-1]
					break
				}
			}
		}

		leftPost = postorder[:len(rightPost)+1]

		cur.Left = buildTree(leftIn, leftPost)
		cur.Right = buildTree(rightIn, rightPost)

		return cur
	}

	//	tree := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	tree := buildTree([]int{2, 1}, []int{2, 1})
	fmt.Println(tree)

}
