package main

/**
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为NULL 的节点将直接作为新二叉树的节点。

示例1:

输入:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
输出:
合并后的树:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7
注意:合并必须从两个树的根节点开始。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-binary-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import (
	"fmt"
)

func main() {

	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	type mergeTrees = func(t1 *TreeNode, t2 *TreeNode) *TreeNode
	var m mergeTrees
	m = func(t1 *TreeNode, t2 *TreeNode) *TreeNode {
		if t1 == nil {
			return t2
		}
		if t2 == nil {
			return t1
		}
		t1.Val += t2.Val
		t1.Left = m(t1.Left, t2.Left)
		t1.Right = m(t1.Right, t2.Right)
		return t1

	}
	fmt.Println(m(nil, nil))

}
