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

	// 效率较低的解法，每次都只找头，然后递归构建
	buildTree = func(preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}
		// 头肯定是前序遍历的第一位
		head := &TreeNode{Val: preorder[0]}
		// 在中序遍历中找到 头 的位置索引
		inHeadIndex := 0
		for i, v := range inorder {
			if v == head.Val {
				inHeadIndex = i
				break
			}
		}
		//根据 头 索引，划分分别属于 left，right 的中序遍历
		leftIn := inorder[:inHeadIndex]
		rightIn := inorder[inHeadIndex+1:]
		leftPreMap := make(map[int]interface{})

		//反过来需要在前序遍历中，抽出分别属于 left ，right 的前序遍历顺序
		for i := range leftIn {
			leftPreMap[leftIn[i]] = nil
		}
		var leftPre []int
		var rightPre []int
		for i := range preorder {
			// 因为元素都不一样，所以直接排除 头元素即可
			if preorder[i] == head.Val {
				continue
			} else if _, ok := leftPreMap[preorder[i]]; ok {
				leftPre = append(leftPre, preorder[i])
			} else {
				rightPre = append(rightPre, preorder[i])
			}
		}
		// 递归分别构建 左树 右树
		head.Left = buildTree(leftPre, leftIn)
		head.Right = buildTree(rightPre, rightIn)
		return head
	}

	tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	fmt.Println(tree)
}
