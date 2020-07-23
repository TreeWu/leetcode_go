package main

import "fmt"

/**
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

 

例如，给出

前序遍历 preorder = [3,9,20,1,15,7]
中序遍历 inorder = [1,9,3,15,20,7]
返回如下的二叉树：

	    3
	   / \
	  9  20
	 /   / \
	1   15  7
	 

限制：

0 <= 节点个数 <= 5000

 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	var buildTree func(preorder []int, inorder []int) *TreeNode

	buildTree = func(preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}
		head := &TreeNode{Val: preorder[0]}
		inHeadIndex := 0
		for i, v := range inorder {
			if v == head.Val {
				inHeadIndex = i
				break
			}
		}
		leftIn := inorder[:inHeadIndex]
		rightIn := inorder[inHeadIndex+1:]
		leftPreMap := make(map[int]interface{})
		for i, _ := range leftIn {
			leftPreMap[leftIn[i]] = nil
		}
		var leftPre []int
		var rightPre []int
		for i, _ := range preorder {
			if _, ok := leftPreMap[preorder[i]]; ok {
				leftPre = append(leftPre, preorder[i])
			} else {
				rightPre = append(rightPre, preorder[i])
			}
		}
		head.Left = buildTree(leftPre, leftIn)
		head.Right = buildTree(rightPre, rightIn)
		return head
	}

	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	fmt.Println(tree)
}
